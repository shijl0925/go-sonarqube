package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/qualityprofiles"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Qualityprofiles service

func (s *Qualityprofiles) ActivateRule(ctx context.Context, r qualityprofiles.ActivateRuleRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/activate_rule", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) ActivateRules(ctx context.Context, r qualityprofiles.ActivateRulesRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/activate_rules", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) AddProject(ctx context.Context, r qualityprofiles.AddProjectRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/add_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Backup(ctx context.Context, r qualityprofiles.BackupRequest) (*qualityprofiles.BackupResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/backup", API)
	v := new(qualityprofiles.BackupResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) ChangeParent(ctx context.Context, r qualityprofiles.ChangeParentRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/change_parent", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Changelog(ctx context.Context, r qualityprofiles.ChangelogRequest, p paging.Params) (*qualityprofiles.ChangelogResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/changelog", API)
	v := new(qualityprofiles.ChangelogResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) ChangelogAll(ctx context.Context, r qualityprofiles.ChangelogRequest) (*qualityprofiles.ChangelogResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualityprofiles.ChangelogResponseAll{}
	for {
		res, err := s.Changelog(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to qualityprofiles.Changelog: %+v", err)
		}
		response.Events = append(response.Events, res.Events...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Qualityprofiles) Copy(ctx context.Context, r qualityprofiles.CopyRequest) (*qualityprofiles.CopyResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/copy", API)
	v := new(qualityprofiles.CopyResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Create(ctx context.Context, r qualityprofiles.CreateRequest) (*qualityprofiles.CreateResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/create", API)
	v := new(qualityprofiles.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) DeactivateRule(ctx context.Context, r qualityprofiles.DeactivateRuleRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/deactivate_rule", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) DeactivateRules(ctx context.Context, r qualityprofiles.DeactivateRulesRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/deactivate_rules", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Delete(ctx context.Context, r qualityprofiles.DeleteRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Export(ctx context.Context, r qualityprofiles.ExportRequest) (*qualityprofiles.ExportResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/export", API)
	v := new(qualityprofiles.ExportResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Exporters(ctx context.Context, r qualityprofiles.ExportersRequest) (*qualityprofiles.ExportersResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/exporters", API)
	v := new(qualityprofiles.ExportersResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Importers(ctx context.Context, r qualityprofiles.ImportersRequest) (*qualityprofiles.ImportersResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/importers", API)
	v := new(qualityprofiles.ImportersResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Inheritance(ctx context.Context, r qualityprofiles.InheritanceRequest) (*qualityprofiles.InheritanceResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/inheritance", API)
	v := new(qualityprofiles.InheritanceResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Projects(ctx context.Context, r qualityprofiles.ProjectsRequest, p paging.Params) (*qualityprofiles.ProjectsResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/projects", API)
	v := new(qualityprofiles.ProjectsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) ProjectsAll(ctx context.Context, r qualityprofiles.ProjectsRequest) (*qualityprofiles.ProjectsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualityprofiles.ProjectsResponseAll{}
	for {
		res, err := s.Projects(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to qualityprofiles.Projects: %+v", err)
		}
		response.Results = append(response.Results, res.Results...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Qualityprofiles) RemoveProject(ctx context.Context, r qualityprofiles.RemoveProjectRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/remove_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Rename(ctx context.Context, r qualityprofiles.RenameRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/rename", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Restore(ctx context.Context, r qualityprofiles.RestoreRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/restore", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Search(ctx context.Context, r qualityprofiles.SearchRequest) (*qualityprofiles.SearchResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/search", API)
	v := new(qualityprofiles.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) SetDefault(ctx context.Context, r qualityprofiles.SetDefaultRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/set_default", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
