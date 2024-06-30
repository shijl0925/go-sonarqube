package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/analysis_cache"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AnalysisCache service

func (s *AnalysisCache) Get(ctx context.Context, r analysis_cache.GetRequest) error {
	u := fmt.Sprintf("%s/analysis_cache/get", API)

	_, err := s.client.Call(ctx, "GET", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
