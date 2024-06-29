package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/favorites"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Favorites service

func (s *Favorites) Add(r favorites.AddRequest) error {
	u := fmt.Sprintf("%s/favorites/add", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Favorites) Remove(r favorites.RemoveRequest) error {
	u := fmt.Sprintf("%s/favorites/remove", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Favorites) Search(r favorites.SearchRequest, p paging.Params) (*favorites.SearchResponse, error) {
	u := fmt.Sprintf("%s/favorites/search", API)
	v := new(favorites.SearchResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Favorites) SearchAll(r favorites.SearchRequest) (*favorites.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &favorites.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
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
