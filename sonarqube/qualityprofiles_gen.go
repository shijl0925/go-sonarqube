package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/qualityprofiles"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Qualityprofiles service

func (s *Qualityprofiles) ActivateRule(r qualityprofiles.ActivateRuleRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/activate_rule", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) ActivateRules(r qualityprofiles.ActivateRulesRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/activate_rules", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) AddProject(r qualityprofiles.AddProjectRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/add_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Backup(r qualityprofiles.BackupRequest) (*qualityprofiles.BackupResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/backup", API)
	v := new(qualityprofiles.BackupResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) ChangeParent(r qualityprofiles.ChangeParentRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/change_parent", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Changelog(r qualityprofiles.ChangelogRequest, p paging.Params) (*qualityprofiles.ChangelogResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/changelog", API)
	v := new(qualityprofiles.ChangelogResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) ChangelogAll(r qualityprofiles.ChangelogRequest) (*qualityprofiles.ChangelogResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualityprofiles.ChangelogResponseAll{}
	for {
		res, err := s.Changelog(r, p)
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

func (s *Qualityprofiles) Copy(r qualityprofiles.CopyRequest) (*qualityprofiles.CopyResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/copy", API)
	v := new(qualityprofiles.CopyResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Create(r qualityprofiles.CreateRequest) (*qualityprofiles.CreateResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/create", API)
	v := new(qualityprofiles.CreateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) DeactivateRule(r qualityprofiles.DeactivateRuleRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/deactivate_rule", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) DeactivateRules(r qualityprofiles.DeactivateRulesRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/deactivate_rules", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Delete(r qualityprofiles.DeleteRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Export(r qualityprofiles.ExportRequest) (*qualityprofiles.ExportResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/export", API)
	v := new(qualityprofiles.ExportResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Exporters(r qualityprofiles.ExportersRequest) (*qualityprofiles.ExportersResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/exporters", API)
	v := new(qualityprofiles.ExportersResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Importers(r qualityprofiles.ImportersRequest) (*qualityprofiles.ImportersResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/importers", API)
	v := new(qualityprofiles.ImportersResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Inheritance(r qualityprofiles.InheritanceRequest) (*qualityprofiles.InheritanceResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/inheritance", API)
	v := new(qualityprofiles.InheritanceResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) Projects(r qualityprofiles.ProjectsRequest, p paging.Params) (*qualityprofiles.ProjectsResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/projects", API)
	v := new(qualityprofiles.ProjectsResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) ProjectsAll(r qualityprofiles.ProjectsRequest) (*qualityprofiles.ProjectsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualityprofiles.ProjectsResponseAll{}
	for {
		res, err := s.Projects(r, p)
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

func (s *Qualityprofiles) RemoveProject(r qualityprofiles.RemoveProjectRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/remove_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Rename(r qualityprofiles.RenameRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/rename", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Restore(r qualityprofiles.RestoreRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/restore", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualityprofiles) Search(r qualityprofiles.SearchRequest) (*qualityprofiles.SearchResponse, error) {
	u := fmt.Sprintf("%s/qualityprofiles/search", API)
	v := new(qualityprofiles.SearchResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualityprofiles) SetDefault(r qualityprofiles.SetDefaultRequest) error {
	u := fmt.Sprintf("%s/qualityprofiles/set_default", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
