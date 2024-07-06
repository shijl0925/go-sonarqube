package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/favorites"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Favorites service

// Add - Add a component (project, portfolio, etc.) as favorite for the authenticated user.
// Only 100 components by qualifier can be added as favorite.
// Requires authentication and the following permission: 'Browse' on the component.
// Since 6.3
// Changelog:
//   10.1: The use of module keys in parameter 'component' is removed
//   8.4: It's no longer possible to set a file as favorite
//   7.7: It's no longer possible to have more than 100 favorites by qualifier
//   7.7: It's no longer possible to set a directory as favorite
//   7.6: The use of module keys in parameter 'component' is deprecated
func (s *Favorites) Add(ctx context.Context, r favorites.AddRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Remove - Remove a component (project, portfolio, application etc.) as favorite for the authenticated user.
// Requires authentication.
// Since 6.3
// Changelog:
//   10.1: The use of module keys in parameter 'component' is removed
//   7.6: The use of module keys in parameter 'component' is deprecated
func (s *Favorites) Remove(ctx context.Context, r favorites.RemoveRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Search - Search for the authenticated user favorites.
// Requires authentication.
// Since 6.3
// Changelog:
func (s *Favorites) Search(ctx context.Context, r favorites.SearchRequest, p paging.Params) (*favorites.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(favorites.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Favorites) SearchAll(ctx context.Context, r favorites.SearchRequest) (*favorites.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &favorites.SearchResponseAll{}
	for {
		res, _, err := s.Search(ctx, r, p)
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
