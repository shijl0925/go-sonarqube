package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_tags"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectTags service

// Search - Search tags
// Since 6.4
// Changelog:
//
//	9.2: Parameter 'page' added
func (s *ProjectTags) Search(ctx context.Context, r project_tags.SearchRequest, p paging.Params) (*project_tags.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(project_tags.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *ProjectTags) SearchAll(ctx context.Context, r project_tags.SearchRequest) (*project_tags.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &project_tags.SearchResponseAll{}
	for {
		res, _, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to project_tags.Search: %+v", err)
		}
		response.Tags = append(response.Tags, res.Tags...)
		if res.GetPaging().End() {
			break
		}
		p.P++
	}
	return response, nil
}

// Set - Set tags on a project.
// Requires the following permission: 'Administer' rights on the specified project
// Since 6.4
func (s *ProjectTags) Set(ctx context.Context, r project_tags.SetRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
