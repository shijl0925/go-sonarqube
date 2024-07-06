package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_links"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectLinks service

// Create - Create a new project link.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 6.1
// Changelog:
func (s *ProjectLinks) Create(ctx context.Context, r project_links.CreateRequest) (*project_links.CreateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)
	v := new(project_links.CreateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Delete - Delete existing project link.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 6.1
// Changelog:
func (s *ProjectLinks) Delete(ctx context.Context, r project_links.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Search - List links of a project.
// The 'projectId' or 'projectKey' must be provided.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//  * 'Browse' on the specified project
//
// Since 6.1
// Changelog:
func (s *ProjectLinks) Search(ctx context.Context, r project_links.SearchRequest) (*project_links.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(project_links.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
