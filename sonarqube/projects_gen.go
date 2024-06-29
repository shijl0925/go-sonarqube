package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/projects"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Projects service

func (s *Projects) BulkDelete(r projects.BulkDeleteRequest) error {
	u := fmt.Sprintf("%s/projects/bulk_delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Projects) Create(r projects.CreateRequest) (*projects.CreateResponse, error) {
	u := fmt.Sprintf("%s/projects/create", API)
	v := new(projects.CreateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Projects) Delete(r projects.DeleteRequest) error {
	u := fmt.Sprintf("%s/projects/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Projects) ExportFindings(r projects.ExportFindingsRequest) (*projects.ExportFindingsResponse, error) {
	u := fmt.Sprintf("%s/projects/export_findings", API)
	v := new(projects.ExportFindingsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Projects) LicenseUsage(r projects.LicenseUsageRequest) (*projects.LicenseUsageResponse, error) {
	u := fmt.Sprintf("%s/projects/license_usage", API)
	v := new(projects.LicenseUsageResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Projects) Search(r projects.SearchRequest, p paging.Params) (*projects.SearchResponse, error) {
	u := fmt.Sprintf("%s/projects/search", API)
	v := new(projects.SearchResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Projects) SearchAll(r projects.SearchRequest) (*projects.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &projects.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to projects.Search: %+v", err)
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

func (s *Projects) UpdateKey(r projects.UpdateKeyRequest) error {
	u := fmt.Sprintf("%s/projects/update_key", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Projects) UpdateVisibility(r projects.UpdateVisibilityRequest) error {
	u := fmt.Sprintf("%s/projects/update_visibility", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
