package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/authentication"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Authentication service

// Login - Authenticate a user.
// Since 6.0
func (s *Authentication) Login(ctx context.Context, r authentication.LoginRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/login", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Logout - Logout a user.
// Since 6.3
func (s *Authentication) Logout(ctx context.Context, r authentication.LogoutRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/logout", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Validate - Check credentials.
// Since 3.3
func (s *Authentication) Validate(ctx context.Context, r authentication.ValidateRequest) (*authentication.ValidateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/validate", s.path)
	v := new(authentication.ValidateResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
