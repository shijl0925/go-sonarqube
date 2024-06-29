package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/editions"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Editions service

func (s *Editions) ActivateGracePeriod(r editions.ActivateGracePeriodRequest) error {
	u := fmt.Sprintf("%s/editions/activate_grace_period", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Editions) SetLicense(r editions.SetLicenseRequest) error {
	u := fmt.Sprintf("%s/editions/set_license", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
