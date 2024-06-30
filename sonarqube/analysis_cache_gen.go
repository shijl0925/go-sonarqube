package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/analysis_cache"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AnalysisCache service

// Get - Get the scanner's cached data for a branch. Requires scan permission on the project. Data is returned gzipped if the corresponding 'Accept-Encoding' header is set in the request.
func (s *AnalysisCache) Get(ctx context.Context, r analysis_cache.GetRequest) error {
	u := fmt.Sprintf("%s/analysis_cache/get", API)

	_, err := s.client.Call(ctx, "GET", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
