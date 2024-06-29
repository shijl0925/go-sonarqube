package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_links"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectLinks service

func (s *ProjectLinks) Create(r project_links.CreateRequest) (*project_links.CreateResponse, error) {
	u := fmt.Sprintf("%s/project_links/create", API)
	v := new(project_links.CreateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectLinks) Delete(r project_links.DeleteRequest) error {
	u := fmt.Sprintf("%s/project_links/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectLinks) Search(r project_links.SearchRequest) (*project_links.SearchResponse, error) {
	u := fmt.Sprintf("%s/project_links/search", API)
	v := new(project_links.SearchResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
