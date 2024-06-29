package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/monitoring"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Monitoring service

func (s *Monitoring) Metrics(r monitoring.MetricsRequest) (*monitoring.MetricsResponse, error) {
	u := fmt.Sprintf("%s/monitoring/metrics", API)
	v := new(monitoring.MetricsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
