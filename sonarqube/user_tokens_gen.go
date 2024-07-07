package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/user_tokens"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type UserTokens service

// Generate - Generate a user access token.
// Please keep your tokens secret. They enable to authenticate and analyze projects.
// It requires administration permissions to specify a 'login' and generate a token for another user. Otherwise, a token is generated for the current user.
// Since 5.3
// Changelog:
//
//	9.6: Response field 'expirationDate' added
func (s *UserTokens) Generate(ctx context.Context, r user_tokens.GenerateRequest) (*user_tokens.GenerateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/generate", s.path)
	v := new(user_tokens.GenerateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Revoke - Revoke a user access token.
// It requires administration permissions to specify a 'login' and revoke a token for another user. Otherwise, the token for the current user is revoked.
// Since 5.3
func (s *UserTokens) Revoke(ctx context.Context, r user_tokens.RevokeRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/revoke", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Search - List the access tokens of a user.
// The login must exist and active.
// Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user is using a token many times in less than one hour.
// It requires administration permissions to specify a 'login' and list the tokens of another user. Otherwise, tokens for the current user are listed.
// Authentication is required for this API endpoint
// Since 5.3
// Changelog:
//
//	9.6: New field 'expirationDate' is added to response
//	7.7: New field 'lastConnectionDate' is added to response
func (s *UserTokens) Search(ctx context.Context, r user_tokens.SearchRequest) (*user_tokens.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(user_tokens.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
