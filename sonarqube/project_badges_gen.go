package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_badges"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectBadges service

// Measure - Generate badge for project's measure as an SVG.
// Requires 'Browse' permission on the specified project.
func (s *ProjectBadges) Measure(ctx context.Context, r project_badges.MeasureRequest) (*project_badges.MeasureResponse, error) {
	u := fmt.Sprintf("%s/project_badges/measure", API)
	v := new(project_badges.MeasureResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// QualityGate - Generate badge for project's quality gate as an SVG.
// Requires 'Browse' permission on the specified project.
func (s *ProjectBadges) QualityGate(ctx context.Context, r project_badges.QualityGateRequest) (*project_badges.QualityGateResponse, error) {
	u := fmt.Sprintf("%s/project_badges/quality_gate", API)
	v := new(project_badges.QualityGateResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// RenewToken - Creates new token replacing any existing token for project or application badge access for private projects and applications.
// This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.
// Requires 'Administer' permission on the specified project or application.
// Since 9.2
// Changelog:
//   10.1: Application key can be used for project parameter.
func (s *ProjectBadges) RenewToken(ctx context.Context, r project_badges.RenewTokenRequest) error {
	u := fmt.Sprintf("%s/project_badges/renew_token", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Token - Retrieve a token to use for project or application badge access for private projects or applications.
// This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.
// Requires 'Browse' permission on the specified project or application.
func (s *ProjectBadges) Token(ctx context.Context, r project_badges.TokenRequest) (*project_badges.TokenResponse, error) {
	u := fmt.Sprintf("%s/project_badges/token", API)
	v := new(project_badges.TokenResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
