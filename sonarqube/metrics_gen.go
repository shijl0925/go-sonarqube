package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/metrics"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Metrics service

// Search - Search for metrics
func (s *Metrics) Search(ctx context.Context, r metrics.SearchRequest, p paging.Params) (*metrics.SearchResponse, error) {
	u := fmt.Sprintf("%s/metrics/search", API)
	v := new(metrics.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Metrics) SearchAll(ctx context.Context, r metrics.SearchRequest) (*metrics.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &metrics.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to metrics.Search: %+v", err)
		}
		response.Metrics = append(response.Metrics, res.Metrics...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// Types - List all available metric types.
func (s *Metrics) Types(ctx context.Context, r metrics.TypesRequest) (*metrics.TypesResponse, error) {
	u := fmt.Sprintf("%s/metrics/types", API)
	v := new(metrics.TypesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
