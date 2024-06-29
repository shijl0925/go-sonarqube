package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/measures"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Measures service

func (s *Measures) Component(r measures.ComponentRequest) (*measures.ComponentResponse, error) {
	u := fmt.Sprintf("%s/measures/component", API)
	v := new(measures.ComponentResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Measures) ComponentTree(r measures.ComponentTreeRequest, p paging.Params) (*measures.ComponentTreeResponse, error) {
	u := fmt.Sprintf("%s/measures/component_tree", API)
	v := new(measures.ComponentTreeResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Measures) ComponentTreeAll(r measures.ComponentTreeRequest) (*measures.ComponentTreeResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &measures.ComponentTreeResponseAll{}
	for {
		res, err := s.ComponentTree(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to measures.ComponentTree: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Metrics = append(response.Metrics, res.Metrics...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Measures) SearchHistory(r measures.SearchHistoryRequest, p paging.Params) (*measures.SearchHistoryResponse, error) {
	u := fmt.Sprintf("%s/measures/search_history", API)
	v := new(measures.SearchHistoryResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Measures) SearchHistoryAll(r measures.SearchHistoryRequest) (*measures.SearchHistoryResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &measures.SearchHistoryResponseAll{}
	for {
		res, err := s.SearchHistory(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to measures.SearchHistory: %+v", err)
		}
		response.Measures = append(response.Measures, res.Measures...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
