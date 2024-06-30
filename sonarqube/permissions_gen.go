package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/permissions"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Permissions service

func (s *Permissions) AddGroup(ctx context.Context, r permissions.AddGroupRequest) error {
	u := fmt.Sprintf("%s/permissions/add_group", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddGroupToTemplate(ctx context.Context, r permissions.AddGroupToTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/add_group_to_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddProjectCreatorToTemplate(ctx context.Context, r permissions.AddProjectCreatorToTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/add_project_creator_to_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddUser(ctx context.Context, r permissions.AddUserRequest) error {
	u := fmt.Sprintf("%s/permissions/add_user", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) AddUserToTemplate(ctx context.Context, r permissions.AddUserToTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/add_user_to_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) ApplyTemplate(ctx context.Context, r permissions.ApplyTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/apply_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) BulkApplyTemplate(ctx context.Context, r permissions.BulkApplyTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/bulk_apply_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) CreateTemplate(ctx context.Context, r permissions.CreateTemplateRequest) (*permissions.CreateTemplateResponse, error) {
	u := fmt.Sprintf("%s/permissions/create_template", API)
	v := new(permissions.CreateTemplateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Permissions) DeleteTemplate(ctx context.Context, r permissions.DeleteTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/delete_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveGroup(ctx context.Context, r permissions.RemoveGroupRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_group", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveGroupFromTemplate(ctx context.Context, r permissions.RemoveGroupFromTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_group_from_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveProjectCreatorFromTemplate(ctx context.Context, r permissions.RemoveProjectCreatorFromTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_project_creator_from_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveUser(ctx context.Context, r permissions.RemoveUserRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_user", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) RemoveUserFromTemplate(ctx context.Context, r permissions.RemoveUserFromTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/remove_user_from_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) SearchTemplates(ctx context.Context, r permissions.SearchTemplatesRequest) (*permissions.SearchTemplatesResponse, error) {
	u := fmt.Sprintf("%s/permissions/search_templates", API)
	v := new(permissions.SearchTemplatesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Permissions) SetDefaultTemplate(ctx context.Context, r permissions.SetDefaultTemplateRequest) error {
	u := fmt.Sprintf("%s/permissions/set_default_template", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Permissions) UpdateTemplate(ctx context.Context, r permissions.UpdateTemplateRequest) (*permissions.UpdateTemplateResponse, error) {
	u := fmt.Sprintf("%s/permissions/update_template", API)
	v := new(permissions.UpdateTemplateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
