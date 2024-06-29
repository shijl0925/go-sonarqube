package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/views"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Views service

func (s *Views) AddApplication(r views.AddApplicationRequest) error {
	u := fmt.Sprintf("%s/views/add_application", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) AddApplicationBranch(r views.AddApplicationBranchRequest) error {
	u := fmt.Sprintf("%s/views/add_application_branch", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) AddPortfolio(r views.AddPortfolioRequest) error {
	u := fmt.Sprintf("%s/views/add_portfolio", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) AddProject(r views.AddProjectRequest) error {
	u := fmt.Sprintf("%s/views/add_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) AddProjectBranch(r views.AddProjectBranchRequest) error {
	u := fmt.Sprintf("%s/views/add_project_branch", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) Applications(r views.ApplicationsRequest) (*views.ApplicationsResponse, error) {
	u := fmt.Sprintf("%s/views/applications", API)
	v := new(views.ApplicationsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Views) Create(r views.CreateRequest) error {
	u := fmt.Sprintf("%s/views/create", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) Delete(r views.DeleteRequest) error {
	u := fmt.Sprintf("%s/views/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) List(r views.ListRequest) (*views.ListResponse, error) {
	u := fmt.Sprintf("%s/views/list", API)
	v := new(views.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Views) Move(r views.MoveRequest) error {
	u := fmt.Sprintf("%s/views/move", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) MoveOptions(r views.MoveOptionsRequest) (*views.MoveOptionsResponse, error) {
	u := fmt.Sprintf("%s/views/move_options", API)
	v := new(views.MoveOptionsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Views) Portfolios(r views.PortfoliosRequest) (*views.PortfoliosResponse, error) {
	u := fmt.Sprintf("%s/views/portfolios", API)
	v := new(views.PortfoliosResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Views) RemoveApplication(r views.RemoveApplicationRequest) error {
	u := fmt.Sprintf("%s/views/remove_application", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) RemoveApplicationBranch(r views.RemoveApplicationBranchRequest) error {
	u := fmt.Sprintf("%s/views/remove_application_branch", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) RemovePortfolio(r views.RemovePortfolioRequest) error {
	u := fmt.Sprintf("%s/views/remove_portfolio", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) RemoveProject(r views.RemoveProjectRequest) error {
	u := fmt.Sprintf("%s/views/remove_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) RemoveProjectBranch(r views.RemoveProjectBranchRequest) error {
	u := fmt.Sprintf("%s/views/remove_project_branch", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) SetManualMode(r views.SetManualModeRequest) error {
	u := fmt.Sprintf("%s/views/set_manual_mode", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) SetNoneMode(r views.SetNoneModeRequest) error {
	u := fmt.Sprintf("%s/views/set_none_mode", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) SetRegexpMode(r views.SetRegexpModeRequest) error {
	u := fmt.Sprintf("%s/views/set_regexp_mode", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) SetRemainingProjectsMode(r views.SetRemainingProjectsModeRequest) error {
	u := fmt.Sprintf("%s/views/set_remaining_projects_mode", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) SetTagsMode(r views.SetTagsModeRequest) error {
	u := fmt.Sprintf("%s/views/set_tags_mode", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Views) Show(r views.ShowRequest) (*views.ShowResponse, error) {
	u := fmt.Sprintf("%s/views/show", API)
	v := new(views.ShowResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Views) Update(r views.UpdateRequest) error {
	u := fmt.Sprintf("%s/views/update", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
