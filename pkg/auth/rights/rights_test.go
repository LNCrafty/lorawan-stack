// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package rights

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

func TestContext(t *testing.T) {
	a := assertions.New(t)

	ctx := test.Context()
	rights, ok := fromContext(ctx)
	a.So(ok, should.BeFalse)
	a.So(rights, should.Resemble, Rights{})

	fooRights := Rights{
		ApplicationRights: map[string]*ttnpb.Rights{
			unique.ID(ctx, &ttnpb.ApplicationIdentifiers{
				ApplicationId: "foo-app",
			}): ttnpb.RightsFrom(ttnpb.Right_RIGHT_APPLICATION_INFO),
		},
		ClientRights: map[string]*ttnpb.Rights{
			unique.ID(ctx, &ttnpb.ClientIdentifiers{
				ClientId: "foo-cli",
			}): ttnpb.RightsFrom(ttnpb.Right_RIGHT_CLIENT_ALL),
		},
		GatewayRights: map[string]*ttnpb.Rights{
			unique.ID(ctx, &ttnpb.GatewayIdentifiers{
				GatewayId: "foo-gtw",
			}): ttnpb.RightsFrom(ttnpb.Right_RIGHT_GATEWAY_INFO),
		},
		OrganizationRights: map[string]*ttnpb.Rights{
			unique.ID(ctx, &ttnpb.OrganizationIdentifiers{
				OrganizationId: "foo-org",
			}): ttnpb.RightsFrom(ttnpb.Right_RIGHT_ORGANIZATION_INFO),
		},
		UserRights: map[string]*ttnpb.Rights{
			unique.ID(ctx, &ttnpb.UserIdentifiers{
				UserId: "foo-usr",
			}): ttnpb.RightsFrom(ttnpb.Right_RIGHT_USER_INFO),
		},
	}

	ctx = NewContext(ctx, fooRights)

	rights, ok = fromContext(ctx)
	a.So(ok, should.BeTrue)
	a.So(rights, should.Resemble, fooRights)
	a.So(rights.IncludesApplicationRights(
		unique.ID(ctx, &ttnpb.ApplicationIdentifiers{ApplicationId: "foo-app"}),
		ttnpb.Right_RIGHT_APPLICATION_INFO,
	), should.BeTrue)
	a.So(rights.IncludesClientRights(
		unique.ID(ctx, &ttnpb.ClientIdentifiers{ClientId: "foo-cli"}),
		ttnpb.Right_RIGHT_CLIENT_ALL,
	), should.BeTrue)
	a.So(rights.IncludesGatewayRights(
		unique.ID(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo-gtw"}),
		ttnpb.Right_RIGHT_GATEWAY_INFO,
	), should.BeTrue)
	a.So(rights.IncludesOrganizationRights(
		unique.ID(ctx, &ttnpb.OrganizationIdentifiers{OrganizationId: "foo-org"}),
		ttnpb.Right_RIGHT_ORGANIZATION_INFO,
	), should.BeTrue)
	a.So(rights.IncludesUserRights(
		unique.ID(ctx, &ttnpb.UserIdentifiers{UserId: "foo-usr"}),
		ttnpb.Right_RIGHT_USER_INFO,
	), should.BeTrue)
}
