package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/hotspots"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Hotspots service

func (s *Hotspots) ChangeStatus(ctx context.Context, r hotspots.ChangeStatusRequest) error {
	u := fmt.Sprintf("%s/hotspots/change_status", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Hotspots) Search(ctx context.Context, r hotspots.SearchRequest, p paging.Params) (*hotspots.SearchResponse, error) {
	u := fmt.Sprintf("%s/hotspots/search", API)
	v := new(hotspots.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Hotspots) SearchAll(ctx context.Context, r hotspots.SearchRequest) (*hotspots.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &hotspots.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to hotspots.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Hotspots = append(response.Hotspots, res.Hotspots...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Hotspots) Show(ctx context.Context, r hotspots.ShowRequest) (*hotspots.ShowResponse, error) {
	u := fmt.Sprintf("%s/hotspots/show", API)
	v := new(hotspots.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
