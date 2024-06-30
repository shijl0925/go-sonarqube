package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/favorites"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Favorites service

func (s *Favorites) Add(ctx context.Context, r favorites.AddRequest) error {
	u := fmt.Sprintf("%s/favorites/add", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Favorites) Remove(ctx context.Context, r favorites.RemoveRequest) error {
	u := fmt.Sprintf("%s/favorites/remove", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Favorites) Search(ctx context.Context, r favorites.SearchRequest, p paging.Params) (*favorites.SearchResponse, error) {
	u := fmt.Sprintf("%s/favorites/search", API)
	v := new(favorites.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Favorites) SearchAll(ctx context.Context, r favorites.SearchRequest) (*favorites.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &favorites.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to favorites.Search: %+v", err)
		}
		response.Favorites = append(response.Favorites, res.Favorites...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
