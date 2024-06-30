package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_links"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectLinks service

func (s *ProjectLinks) Create(ctx context.Context, r project_links.CreateRequest) (*project_links.CreateResponse, error) {
	u := fmt.Sprintf("%s/project_links/create", API)
	v := new(project_links.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectLinks) Delete(ctx context.Context, r project_links.DeleteRequest) error {
	u := fmt.Sprintf("%s/project_links/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectLinks) Search(ctx context.Context, r project_links.SearchRequest) (*project_links.SearchResponse, error) {
	u := fmt.Sprintf("%s/project_links/search", API)
	v := new(project_links.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
