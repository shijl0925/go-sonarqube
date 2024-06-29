package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/rules"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Rules service

func (s *Rules) Create(r rules.CreateRequest) (*rules.CreateResponse, error) {
	u := fmt.Sprintf("%s/rules/create", API)
	v := new(rules.CreateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Rules) Delete(r rules.DeleteRequest) error {
	u := fmt.Sprintf("%s/rules/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Rules) Repositories(r rules.RepositoriesRequest) (*rules.RepositoriesResponse, error) {
	u := fmt.Sprintf("%s/rules/repositories", API)
	v := new(rules.RepositoriesResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Rules) Search(r rules.SearchRequest, p paging.Params) (*rules.SearchResponse, error) {
	u := fmt.Sprintf("%s/rules/search", API)
	v := new(rules.SearchResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Rules) SearchAll(r rules.SearchRequest) (*rules.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &rules.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to rules.Search: %+v", err)
		}
		response.Facets = append(response.Facets, res.Facets...)
		response.Rules = append(response.Rules, res.Rules...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Rules) Show(r rules.ShowRequest) (*rules.ShowResponse, error) {
	u := fmt.Sprintf("%s/rules/show", API)
	v := new(rules.ShowResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Rules) Tags(r rules.TagsRequest) (*rules.TagsResponse, error) {
	u := fmt.Sprintf("%s/rules/tags", API)
	v := new(rules.TagsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Rules) Update(r rules.UpdateRequest) (*rules.UpdateResponse, error) {
	u := fmt.Sprintf("%s/rules/update", API)
	v := new(rules.UpdateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
