// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package ttnpb

import (
	"regexp"
	"strings"

	"github.com/gobwas/glob"
)

var (
	// FieldPathSettingsBlacklistedIDs is the field path for the blacklisted IDs field.
	FieldPathSettingsBlacklistedIDs = regexp.MustCompile(`^blacklisted_ids$`)

	// FieldPathSettingsUserRegistrationSkipValidation is the field path for the
	// user registration flow skip validation field.
	FieldPathSettingsUserRegistrationSkipValidation = regexp.MustCompile(`^user_registration.skip_validation$`)

	// FieldPathSettingsUserRegistrationSelfRegistration is the field path for the
	// user registration flow self registration field.
	FieldPathSettingsUserRegistrationSelfRegistration = regexp.MustCompile(`^user_registration.self_registration$`)

	// FieldPathSettingsUserRegistrationAdminApproval is the field path for the
	// user registration flow admin approval field.
	FieldPathSettingsUserRegistrationAdminApproval = regexp.MustCompile(`^user_registration.admin_approval$`)

	// FieldPathSettingsValidationTokenTTL is the field path for the validation token TTL field.
	FieldPathSettingsValidationTokenTTL = regexp.MustCompile(`^validation_token_ttl$`)

	// FieldPathSettingsAllowedEmails is the field path for the allowed emails field.
	FieldPathSettingsAllowedEmails = regexp.MustCompile(`^allowed_emails$`)
)

// IsEmailAllowed checks whether an input email is allowed given the glob list
// of allowed emails in the settings.
func (s *IdentityServerSettings) IsEmailAllowed(email string) bool {
	if s.AllowedEmails == nil || len(s.AllowedEmails) == 0 {
		return true
	}

	found := false
	for i, _ := range s.AllowedEmails {
		found = glob.MustCompile(strings.ToLower(s.AllowedEmails[i])).Match(strings.ToLower(email))
		if found {
			break
		}
	}

	return found
}
