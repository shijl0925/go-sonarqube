package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/editions"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Editions service

// ActivateGracePeriod - Enable a license 7-days grace period if the Server ID is invalid. Require 'Administer System' permission.
// Since 10.3
// Changelog:
func (s *Editions) ActivateGracePeriod(ctx context.Context, r editions.ActivateGracePeriodRequest) error {
	u := fmt.Sprintf("%s/activate_grace_period", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetLicense - Set the license for enabling features of commercial editions. Require 'Administer System' permission.
// Since 7.2
// Changelog:
func (s *Editions) SetLicense(ctx context.Context, r editions.SetLicenseRequest) error {
	u := fmt.Sprintf("%s/set_license", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
