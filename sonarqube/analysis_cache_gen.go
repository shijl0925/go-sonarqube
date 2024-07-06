package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/analysis_cache"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AnalysisCache service

// Get - Get the scanner's cached data for a branch. Requires scan permission on the project. Data is returned gzipped if the corresponding 'Accept-Encoding' header is set in the request.
// Since 9.4
// Changelog:
//   9.9: The web service is no longer internal
func (s *AnalysisCache) Get(ctx context.Context, r analysis_cache.GetRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/get", s.path)

	resp, err := s.client.Call(ctx, "GET", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
