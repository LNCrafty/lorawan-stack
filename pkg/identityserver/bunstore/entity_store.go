// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type entityStore struct {
	*baseStore
	*applicationStore
	*clientStore
	*endDeviceStore
	*gatewayStore
	*organizationStore
	*userStore
}

func newEntityStore(baseStore *baseStore) *entityStore {
	return &entityStore{
		baseStore:         baseStore,
		applicationStore:  &applicationStore{baseStore},
		clientStore:       &clientStore{baseStore},
		endDeviceStore:    &endDeviceStore{baseStore},
		gatewayStore:      &gatewayStore{baseStore},
		organizationStore: &organizationStore{baseStore},
		userStore:         &userStore{baseStore},
	}
}

func (s *entityStore) query(ctx context.Context, entityType string, entityIDs ...string) *bun.SelectQuery {
	query := s.DB.NewSelect()

	switch entityType {
	case "application":
		query = query.
			Model(&Application{}).
			Column("id").
			Apply(s.applicationStore.selectWithID(ctx, entityIDs...))
	case "client":
		query = query.
			Model(&Client{}).
			Column("id").
			Apply(s.clientStore.selectWithID(ctx, entityIDs...))
	case "gateway":
		query = query.
			Model(&Gateway{}).
			Column("id").
			Apply(s.gatewayStore.selectWithID(ctx, entityIDs...))
	case "organization", "user":
		query = query.
			Model(&Account{}).
			Apply(selectWithContext(ctx)).
			Column("account_id").
			Where("?TableAlias.account_type = ?", entityType)
		switch len(entityIDs) {
		case 0:
		case 1:
			query = query.Where("?TableAlias.uid = (?)", entityIDs[0])
		default:
			query = query.Where("?TableAlias.uid IN (?)", bun.In(entityIDs))
		}
	}

	query = query.Apply(selectWithSoftDeletedFromContext(ctx))

	return query
}

type identifiers interface {
	GetEntityIdentifiers() *ttnpb.EntityIdentifiers
}

var entityTypeReplacer = strings.NewReplacer(" ", "_")

func getEntityType(ids ttnpb.IDStringer) string {
	return entityTypeReplacer.Replace(ids.EntityType())
}

func (s *entityStore) getEntity(ctx context.Context, ids ttnpb.IDStringer) (entityType, entityUUID string, err error) {
	entityType = getEntityType(ids)

	if entityType == "end_device" {
		entityIDs, ok := ids.(identifiers)
		if !ok {
			entityIDs = getEntityIdentifiers(entityType, ids.IDString())
		}
		devIDs := entityIDs.GetEntityIdentifiers().GetDeviceIds()
		model, err := s.getEndDeviceModelBy(
			ctx,
			s.endDeviceStore.selectWithID(ctx, devIDs.GetApplicationIds().GetApplicationId(), devIDs.GetDeviceId()),
			store.FieldMask{"ids"},
		)
		if err != nil {
			return "", "", err
		}
		return entityType, model.ID, nil
	}

	var uuid string
	err = s.query(ctx, entityType, ids.IDString()).Scan(ctx, &uuid)
	if err != nil {
		err = wrapDriverError(err)
		if errors.IsNotFound(err) {
			return "", "", store.ErrEntityNotFound.WithAttributes(
				"entity_type", ids.EntityType(),
				"entity_id", ids.IDString(),
			)
		}
		return "", "", err
	}
	return entityType, uuid, nil
}

func (s *entityStore) getEntityUUIDs(ctx context.Context, entityType string, entityIDs ...string) ([]string, error) {
	var uuids []string
	err := s.query(ctx, entityType, entityIDs...).Scan(ctx, &uuids)
	if err != nil {
		return nil, wrapDriverError(err)
	}
	return uuids, nil
}

func (s *entityStore) getEntityID(ctx context.Context, entityType, entityUUID string) (string, error) {
	query := s.DB.NewSelect()

	switch entityType {
	case "application":
		query = query.
			Model(&Application{}).
			Column("application_id").
			Where("id = ?", entityUUID)
	case "client":
		query = query.
			Model(&Client{}).
			Column("client_id").
			Where("id = ?", entityUUID)
	case "gateway":
		query = query.
			Model(&Gateway{}).
			Column("gateway_id").
			Where("id = ?", entityUUID)
	case "organization", "user":
		query = query.
			Model(&Account{}).
			Column("uid").
			Where("account_id = ?", entityUUID).
			Where("?TableAlias.account_type = ?", entityType)
	}

	query = query.
		Apply(selectWithContext(ctx)).
		Apply(selectWithSoftDeletedFromContext(ctx))

	var friendlyID string
	err := query.Scan(ctx, &friendlyID)
	if err != nil {
		return "", wrapDriverError(err)
	}

	return friendlyID, nil
}

func getEntityIdentifiers(entityType string, friendlyID string) *ttnpb.EntityIdentifiers {
	switch entityType {
	default:
		panic(fmt.Errorf("invalid entity type: %s", entityType))
	case "application":
		return (&ttnpb.ApplicationIdentifiers{ApplicationId: friendlyID}).GetEntityIdentifiers()
	case "client":
		return (&ttnpb.ClientIdentifiers{ClientId: friendlyID}).GetEntityIdentifiers()
	case "end_device", "end device":
		parts := strings.SplitN(friendlyID, ".", 2)
		return (&ttnpb.EndDeviceIdentifiers{
			ApplicationIds: &ttnpb.ApplicationIdentifiers{
				ApplicationId: parts[0],
			},
			DeviceId: parts[1],
		}).GetEntityIdentifiers()
	case "gateway":
		return (&ttnpb.GatewayIdentifiers{GatewayId: friendlyID}).GetEntityIdentifiers()
	case "organization":
		return (&ttnpb.OrganizationIdentifiers{OrganizationId: friendlyID}).GetEntityIdentifiers()
	case "user":
		return (&ttnpb.UserIdentifiers{UserId: friendlyID}).GetEntityIdentifiers()
	}
}
