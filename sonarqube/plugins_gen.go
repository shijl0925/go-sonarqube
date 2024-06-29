package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/plugins"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Plugins service

func (s *Plugins) Available(r plugins.AvailableRequest) (*plugins.AvailableResponse, error) {
	u := fmt.Sprintf("%s/plugins/available", API)
	v := new(plugins.AvailableResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Plugins) CancelAll(r plugins.CancelAllRequest) error {
	u := fmt.Sprintf("%s/plugins/cancel_all", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Install(r plugins.InstallRequest) error {
	u := fmt.Sprintf("%s/plugins/install", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Installed(r plugins.InstalledRequest) (*plugins.InstalledResponse, error) {
	u := fmt.Sprintf("%s/plugins/installed", API)
	v := new(plugins.InstalledResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Plugins) Pending(r plugins.PendingRequest) (*plugins.PendingResponse, error) {
	u := fmt.Sprintf("%s/plugins/pending", API)
	v := new(plugins.PendingResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Plugins) Uninstall(r plugins.UninstallRequest) error {
	u := fmt.Sprintf("%s/plugins/uninstall", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Update(r plugins.UpdateRequest) error {
	u := fmt.Sprintf("%s/plugins/update", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Updates(r plugins.UpdatesRequest) (*plugins.UpdatesResponse, error) {
	u := fmt.Sprintf("%s/plugins/updates", API)
	v := new(plugins.UpdatesResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
