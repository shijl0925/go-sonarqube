package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_badges"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectBadges service

// AiCodeAssurance - Generate a badge for project's AI assurance as an SVG.
// Requires 'Browse' permission on the specified project.
// Since 10.7
func (s *ProjectBadges) AiCodeAssurance(ctx context.Context, r project_badges.AiCodeAssuranceRequest) (*project_badges.AiCodeAssuranceResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/ai_code_assurance", s.path)
	v := new(project_badges.AiCodeAssuranceResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Measure - Generate badge for project's measure as an SVG.
// Requires 'Browse' permission on the specified project.
// Since 7.1
// Changelog:
//
//	10.8: The following metric keys are not deprecated anymore: bugs, code_smells, security_hotspots, vulnerabilities
//	10.4: The following metric keys are now deprecated: bugs, code_smells, security_hotspots, vulnerabilities
func (s *ProjectBadges) Measure(ctx context.Context, r project_badges.MeasureRequest) (*project_badges.MeasureResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/measure", s.path)
	v := new(project_badges.MeasureResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// QualityGate - Generate badge for project's quality gate as an SVG.
// Requires 'Browse' permission on the specified project.
// Since 7.1
func (s *ProjectBadges) QualityGate(ctx context.Context, r project_badges.QualityGateRequest) (*project_badges.QualityGateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/quality_gate", s.path)
	v := new(project_badges.QualityGateResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// RenewToken - Creates new token replacing any existing token for project or application badge access for private projects and applications.
// This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.
// Requires 'Administer' permission on the specified project or application.
// If the 'sonar.forceAuthentication' setting is enabled, then a token is required for public projects as well.
// Since 9.2
// Changelog:
//
//	10.1: Application key can be used for project parameter.
func (s *ProjectBadges) RenewToken(ctx context.Context, r project_badges.RenewTokenRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/renew_token", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Token - Retrieve a token to use for project or application badge access for private projects or applications.
// This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.
// Requires 'Browse' permission on the specified project or application.
// If the 'sonar.forceAuthentication' setting is enabled, then a token is required for public projects as well.
// Since 9.2
// Changelog:
//
//	10.1: Application key can be used for project parameter.
func (s *ProjectBadges) Token(ctx context.Context, r project_badges.TokenRequest) (*project_badges.TokenResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/token", s.path)
	v := new(project_badges.TokenResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
