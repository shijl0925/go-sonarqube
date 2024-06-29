package editions

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ActivateGracePeriodRequest Enable a license 7-days grace period if the Server ID is invalid. Require 'Administer System' permission.
type ActivateGracePeriodRequest struct{}

// SetLicenseRequest Set the license for enabling features of commercial editions. Require 'Administer System' permission.
type SetLicenseRequest struct {
	License string `form:"license,omitempty"` //
}
