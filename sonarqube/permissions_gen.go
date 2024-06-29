package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/permissions"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Permissions service

func (s *Permissions) AddGroup(r permissions.AddGroupRequest) error {
	u := fmt.Sprintf("%s/permissions/add_group", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddGroupToTemplate(r permissions.AddGroupToTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/add_group_to_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddProjectCreatorToTemplate(r permissions.AddProjectCreatorToTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/add_project_creator_to_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddUser(r permissions.AddUserRequest) error {
	u := fmt.Sprintf("%s/permissions/add_user", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddUserToTemplate(r permissions.AddUserToTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/add_user_to_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) ApplyTemplate(r permissions.ApplyTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/apply_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) BulkApplyTemplate(r permissions.BulkApplyTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/bulk_apply_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) CreateTemplate(r permissions.CreateTemplateRequest) (*permissions.CreateTemplateResponse, error) {
	u := fmt.Sprintf("%s/permissions/create_template", API)
	v := new(permissions.CreateTemplateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Permissions) DeleteTemplate(r permissions.DeleteTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/delete_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveGroup(r permissions.RemoveGroupRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_group", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveGroupFromTemplate(r permissions.RemoveGroupFromTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_group_from_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveProjectCreatorFromTemplate(r permissions.RemoveProjectCreatorFromTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_project_creator_from_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveUser(r permissions.RemoveUserRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_user", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveUserFromTemplate(r permissions.RemoveUserFromTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_user_from_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) SearchTemplates(r permissions.SearchTemplatesRequest) (*permissions.SearchTemplatesResponse, error) {
	u := fmt.Sprintf("%s/permissions/search_templates", API)
	v := new(permissions.SearchTemplatesResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Permissions) SetDefaultTemplate(r permissions.SetDefaultTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/set_default_template", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) UpdateTemplate(r permissions.UpdateTemplateRequest) (*permissions.UpdateTemplateResponse, error) {
	u := fmt.Sprintf("%s/permissions/update_template", API)
	v := new(permissions.UpdateTemplateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
