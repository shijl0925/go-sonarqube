package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/plugins"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Plugins service

func (s *Plugins) Available(ctx context.Context, r plugins.AvailableRequest) (*plugins.AvailableResponse, error) {
	u := fmt.Sprintf("%s/plugins/available", API)
	v := new(plugins.AvailableResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Plugins) CancelAll(ctx context.Context, r plugins.CancelAllRequest) error {
	u := fmt.Sprintf("%s/plugins/cancel_all", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Install(ctx context.Context, r plugins.InstallRequest) error {
	u := fmt.Sprintf("%s/plugins/install", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Installed(ctx context.Context, r plugins.InstalledRequest) (*plugins.InstalledResponse, error) {
	u := fmt.Sprintf("%s/plugins/installed", API)
	v := new(plugins.InstalledResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Plugins) Pending(ctx context.Context, r plugins.PendingRequest) (*plugins.PendingResponse, error) {
	u := fmt.Sprintf("%s/plugins/pending", API)
	v := new(plugins.PendingResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Plugins) Uninstall(ctx context.Context, r plugins.UninstallRequest) error {
	u := fmt.Sprintf("%s/plugins/uninstall", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Update(ctx context.Context, r plugins.UpdateRequest) error {
	u := fmt.Sprintf("%s/plugins/update", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Plugins) Updates(ctx context.Context, r plugins.UpdatesRequest) (*plugins.UpdatesResponse, error) {
	u := fmt.Sprintf("%s/plugins/updates", API)
	v := new(plugins.UpdatesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
