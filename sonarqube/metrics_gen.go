package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/metrics"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Metrics service

// Search - Search for metrics
// Since 5.2
// Changelog:
//
//	8.4: Field 'id' in the response is deprecated
func (s *Metrics) Search(ctx context.Context, r metrics.SearchRequest, p paging.Params) (*metrics.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(metrics.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Metrics) SearchAll(ctx context.Context, r metrics.SearchRequest) (*metrics.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &metrics.SearchResponseAll{}
	for {
		res, _, err := s.Search(ctx, r, p)
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
// Since 5.2
func (s *Metrics) Types(ctx context.Context, r metrics.TypesRequest) (*metrics.TypesResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/types", s.path)
	v := new(metrics.TypesResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
