package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/authentication"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Authentication service

// Login - Authenticate a user.
// Since 6.0
// Changelog:
func (s *Authentication) Login(ctx context.Context, r authentication.LoginRequest) error {
	u := fmt.Sprintf("%s/authentication/login", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Logout - Logout a user.
// Since 6.3
// Changelog:
func (s *Authentication) Logout(ctx context.Context, r authentication.LogoutRequest) error {
	u := fmt.Sprintf("%s/authentication/logout", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Validate - Check credentials.
func (s *Authentication) Validate(ctx context.Context, r authentication.ValidateRequest) (*authentication.ValidateResponse, error) {
	u := fmt.Sprintf("%s/authentication/validate", API)
	v := new(authentication.ValidateResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
