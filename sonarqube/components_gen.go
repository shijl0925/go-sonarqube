package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/components"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Components service

func (s *Components) Search(r components.SearchRequest, p paging.Params) (*components.SearchResponse, error) {
	u := fmt.Sprintf("%s/components/search", API)
	v := new(components.SearchResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Components) SearchAll(r components.SearchRequest) (*components.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &components.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to components.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Components) Show(r components.ShowRequest) (*components.ShowResponse, error) {
	u := fmt.Sprintf("%s/components/show", API)
	v := new(components.ShowResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Components) Tree(r components.TreeRequest, p paging.Params) (*components.TreeResponse, error) {
	u := fmt.Sprintf("%s/components/tree", API)
	v := new(components.TreeResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Components) TreeAll(r components.TreeRequest) (*components.TreeResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &components.TreeResponseAll{}
	for {
		res, err := s.Tree(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to components.Tree: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
