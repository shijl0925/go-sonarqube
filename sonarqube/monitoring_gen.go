package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/monitoring"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Monitoring service

// Metrics - Return monitoring metrics in Prometheus format.
// Support content type 'text/plain' (default) and 'application/openmetrics-text'.
// this endpoint can be access using a Bearer token, that needs to be defined in sonar.properties with the 'sonar.web.systemPasscode' key.
// Since 9.3
func (s *Monitoring) Metrics(ctx context.Context, r monitoring.MetricsRequest) (*monitoring.MetricsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/metrics", s.path)
	v := new(monitoring.MetricsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
