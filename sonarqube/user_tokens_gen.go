package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/user_tokens"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type UserTokens service

func (s *UserTokens) Generate(ctx context.Context, r user_tokens.GenerateRequest) (*user_tokens.GenerateResponse, error) {
	u := fmt.Sprintf("%s/user_tokens/generate", API)
	v := new(user_tokens.GenerateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *UserTokens) Revoke(ctx context.Context, r user_tokens.RevokeRequest) error {
	u := fmt.Sprintf("%s/user_tokens/revoke", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserTokens) Search(ctx context.Context, r user_tokens.SearchRequest) (*user_tokens.SearchResponse, error) {
	u := fmt.Sprintf("%s/user_tokens/search", API)
	v := new(user_tokens.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
