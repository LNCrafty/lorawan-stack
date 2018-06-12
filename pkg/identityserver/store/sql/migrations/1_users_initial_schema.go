// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

package migrations

func init() {
	const forwards = `
		CREATE TABLE IF NOT EXISTS accounts (
			id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			account_id   STRING(36) UNIQUE NOT NULL,
			type         INT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS users (
			id                        UUID PRIMARY KEY REFERENCES accounts(id),
			user_id                   STRING(36) UNIQUE NOT NULL REFERENCES accounts(account_id),
			name                      STRING NOT NULL DEFAULT '',
			email                     STRING UNIQUE NOT NULL,
			password                  STRING NOT NULL,
			password_updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			require_password_update   BOOL NOT NULL DEFAULT FALSE,
			validated_at              TIMESTAMP DEFAULT NULL,
			state                     INT NOT NULL DEFAULT 0,
			admin                     BOOL NOT NULL DEFAULT false,
			created_at                TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at                TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		CREATE UNIQUE INDEX IF NOT EXISTS users_email ON users (email);
		CREATE TABLE IF NOT EXISTS validation_tokens (
			id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			validation_token   STRING UNIQUE NOT NULL,
			user_id            UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_in         INTEGER NOT NULL
		);
		CREATE TABLE IF NOT EXISTS users_api_keys (
			user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			key_name   STRING(36) NOT NULL,
			key        STRING NOT NULL UNIQUE,
			PRIMARY KEY(user_id, key_name)
		);
		CREATE TABLE IF NOT EXISTS users_api_keys_rights (
			user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			key_name   STRING(36) NOT NULL,
			"right"    STRING NOT NULL,
			PRIMARY KEY(user_id, key_name, "right")
		);
	`

	const backwards = `
		DROP TABLE IF EXISTS users_api_keys_rights;
		DROP TABLE IF EXISTS users_api_keys;
		DROP TABLE IF EXISTS validation_tokens;
		DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS accounts;
	`

	Registry.Register(1, "1_users_initial_schema", forwards, backwards)
}
