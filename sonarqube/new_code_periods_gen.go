package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/new_code_periods"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type NewCodePeriods service

func (s *NewCodePeriods) List(ctx context.Context, r new_code_periods.ListRequest) (*new_code_periods.ListResponse, error) {
	u := fmt.Sprintf("%s/new_code_periods/list", API)
	v := new(new_code_periods.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *NewCodePeriods) Set(ctx context.Context, r new_code_periods.SetRequest) error {
	u := fmt.Sprintf("%s/new_code_periods/set", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *NewCodePeriods) Show(ctx context.Context, r new_code_periods.ShowRequest) (*new_code_periods.ShowResponse, error) {
	u := fmt.Sprintf("%s/new_code_periods/show", API)
	v := new(new_code_periods.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *NewCodePeriods) Unset(ctx context.Context, r new_code_periods.UnsetRequest) error {
	u := fmt.Sprintf("%s/new_code_periods/unset", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
