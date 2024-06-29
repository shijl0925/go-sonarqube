package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_pull_requests"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectPullRequests service

func (s *ProjectPullRequests) Delete(r project_pull_requests.DeleteRequest) error {
	u := fmt.Sprintf("%s/project_pull_requests/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectPullRequests) List(r project_pull_requests.ListRequest) (*project_pull_requests.ListResponse, error) {
	u := fmt.Sprintf("%s/project_pull_requests/list", API)
	v := new(project_pull_requests.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
