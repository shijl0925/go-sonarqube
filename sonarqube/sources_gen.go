package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/sources"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Sources service

func (s *Sources) Raw(ctx context.Context, r sources.RawRequest) (*sources.RawResponse, error) {
	u := fmt.Sprintf("%s/sources/raw", API)
	v := new(sources.RawResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Sources) Scm(ctx context.Context, r sources.ScmRequest) (*sources.ScmResponse, error) {
	u := fmt.Sprintf("%s/sources/scm", API)
	v := new(sources.ScmResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Sources) Show(ctx context.Context, r sources.ShowRequest) (*sources.ShowResponse, error) {
	u := fmt.Sprintf("%s/sources/show", API)
	v := new(sources.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
