package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/settings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Settings service

func (s *Settings) ListDefinitions(ctx context.Context, r settings.ListDefinitionsRequest) (*settings.ListDefinitionsResponse, error) {
	u := fmt.Sprintf("%s/settings/list_definitions", API)
	v := new(settings.ListDefinitionsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Settings) Reset(ctx context.Context, r settings.ResetRequest) error {
	u := fmt.Sprintf("%s/settings/reset", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Settings) Set(ctx context.Context, r settings.SetRequest) error {
	u := fmt.Sprintf("%s/settings/set", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Settings) Values(ctx context.Context, r settings.ValuesRequest) (*settings.ValuesResponse, error) {
	u := fmt.Sprintf("%s/settings/values", API)
	v := new(settings.ValuesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
