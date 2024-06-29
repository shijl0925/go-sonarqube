package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/issues"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Issues service

func (s *Issues) AddComment(r issues.AddCommentRequest) (*issues.AddCommentResponse, error) {
	u := fmt.Sprintf("%s/issues/add_comment", API)
	v := new(issues.AddCommentResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) Assign(r issues.AssignRequest) (*issues.AssignResponse, error) {
	u := fmt.Sprintf("%s/issues/assign", API)
	v := new(issues.AssignResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) Authors(r issues.AuthorsRequest) (*issues.AuthorsResponse, error) {
	u := fmt.Sprintf("%s/issues/authors", API)
	v := new(issues.AuthorsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) BulkChange(r issues.BulkChangeRequest) (*issues.BulkChangeResponse, error) {
	u := fmt.Sprintf("%s/issues/bulk_change", API)
	v := new(issues.BulkChangeResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) Changelog(r issues.ChangelogRequest) (*issues.ChangelogResponse, error) {
	u := fmt.Sprintf("%s/issues/changelog", API)
	v := new(issues.ChangelogResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) DeleteComment(r issues.DeleteCommentRequest) (*issues.DeleteCommentResponse, error) {
	u := fmt.Sprintf("%s/issues/delete_comment", API)
	v := new(issues.DeleteCommentResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) DoTransition(r issues.DoTransitionRequest) (*issues.DoTransitionResponse, error) {
	u := fmt.Sprintf("%s/issues/do_transition", API)
	v := new(issues.DoTransitionResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) EditComment(r issues.EditCommentRequest) (*issues.EditCommentResponse, error) {
	u := fmt.Sprintf("%s/issues/edit_comment", API)
	v := new(issues.EditCommentResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) GitlabSastExport(r issues.GitlabSastExportRequest) (*issues.GitlabSastExportResponse, error) {
	u := fmt.Sprintf("%s/issues/gitlab_sast_export", API)
	v := new(issues.GitlabSastExportResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) Reindex(r issues.ReindexRequest) error {
	u := fmt.Sprintf("%s/issues/reindex", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Issues) Search(r issues.SearchRequest, p paging.Params) (*issues.SearchResponse, error) {
	u := fmt.Sprintf("%s/issues/search", API)
	v := new(issues.SearchResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) SearchAll(r issues.SearchRequest) (*issues.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &issues.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to issues.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Facets = append(response.Facets, res.Facets...)
		response.Issues = append(response.Issues, res.Issues...)
		response.Rules = append(response.Rules, res.Rules...)
		response.Users = append(response.Users, res.Users...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Issues) SetSeverity(r issues.SetSeverityRequest) (*issues.SetSeverityResponse, error) {
	u := fmt.Sprintf("%s/issues/set_severity", API)
	v := new(issues.SetSeverityResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) SetTags(r issues.SetTagsRequest) (*issues.SetTagsResponse, error) {
	u := fmt.Sprintf("%s/issues/set_tags", API)
	v := new(issues.SetTagsResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) SetType(r issues.SetTypeRequest) (*issues.SetTypeResponse, error) {
	u := fmt.Sprintf("%s/issues/set_type", API)
	v := new(issues.SetTypeResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) Tags(r issues.TagsRequest) (*issues.TagsResponse, error) {
	u := fmt.Sprintf("%s/issues/tags", API)
	v := new(issues.TagsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
