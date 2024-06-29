package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_badges"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectBadges service

func (s *ProjectBadges) Measure(r project_badges.MeasureRequest) (*project_badges.MeasureResponse, error) {
	u := fmt.Sprintf("%s/project_badges/measure", API)
	v := new(project_badges.MeasureResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectBadges) QualityGate(r project_badges.QualityGateRequest) (*project_badges.QualityGateResponse, error) {
	u := fmt.Sprintf("%s/project_badges/quality_gate", API)
	v := new(project_badges.QualityGateResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectBadges) RenewToken(r project_badges.RenewTokenRequest) error {
	u := fmt.Sprintf("%s/project_badges/renew_token", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectBadges) Token(r project_badges.TokenRequest) (*project_badges.TokenResponse, error) {
	u := fmt.Sprintf("%s/project_badges/token", API)
	v := new(project_badges.TokenResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
