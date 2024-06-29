package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_tags"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectTags service

func (s *ProjectTags) Search(r project_tags.SearchRequest, p paging.Params) (*project_tags.SearchResponse, error) {
	u := fmt.Sprintf("%s/project_tags/search", API)
	v := new(project_tags.SearchResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectTags) SearchAll(r project_tags.SearchRequest) (*project_tags.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &project_tags.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to project_tags.Search: %+v", err)
		}
		response.Tags = append(response.Tags, res.Tags...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *ProjectTags) Set(r project_tags.SetRequest) error {
	u := fmt.Sprintf("%s/project_tags/set", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
