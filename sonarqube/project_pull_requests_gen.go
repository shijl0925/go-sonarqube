package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_pull_requests"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectPullRequests service

// Delete - Delete a pull request.
// Requires 'Administer' rights on the specified project.
// Since 7.1
func (s *ProjectPullRequests) Delete(ctx context.Context, r project_pull_requests.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// List - List the pull requests of a project.
// One of the following permissions is required:
//  * 'Browse' rights on the specified project
//  * 'Execute Analysis' rights on the specified project
//
// Since 7.1
// Changelog:
//
//	9.8: Response fields: 'bugs', 'vulnerabilities', 'codeSmells' has been dropped.
//	8.4: Response fields: 'bugs', 'vulnerabilities', 'codeSmells' are deprecated.
func (s *ProjectPullRequests) List(ctx context.Context, r project_pull_requests.ListRequest) (*project_pull_requests.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(project_pull_requests.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
