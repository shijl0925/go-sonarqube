package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/permissions"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Permissions service

// AddGroup - Add a permission to a group.
//  This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.
//  The group name must be provided.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//
// Since 5.2
// Changelog:
//   10.0: Parameter 'groupId' is removed. Use 'groupName' instead.
//   8.4: Parameter 'groupId' is deprecated. Format changes from integer to string. Use 'groupName' instead.
func (s *Permissions) AddGroup(ctx context.Context, r permissions.AddGroupRequest) error {
	u := fmt.Sprintf("%s/add_group", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddGroupToTemplate - Add a group to a permission template.
//  The group name must be provided.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
//   10.0: Parameter 'groupId' is removed. Use 'groupName' instead.
//   8.4: Parameter 'groupId' is deprecated. Format changes from integer to string. Use 'groupName' instead.
func (s *Permissions) AddGroupToTemplate(ctx context.Context, r permissions.AddGroupToTemplateRequest) error {
	u := fmt.Sprintf("%s/add_group_to_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddProjectCreatorToTemplate - Add a project creator to a permission template.
// Requires the following permission: 'Administer System'.
// Since 6.0
// Changelog:
func (s *Permissions) AddProjectCreatorToTemplate(ctx context.Context, r permissions.AddProjectCreatorToTemplateRequest) error {
	u := fmt.Sprintf("%s/add_project_creator_to_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddUser - Add permission to a user.
//  This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//
// Since 5.2
// Changelog:
func (s *Permissions) AddUser(ctx context.Context, r permissions.AddUserRequest) error {
	u := fmt.Sprintf("%s/add_user", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddUserToTemplate - Add a user to a permission template.
//  Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) AddUserToTemplate(ctx context.Context, r permissions.AddUserToTemplateRequest) error {
	u := fmt.Sprintf("%s/add_user_to_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// ApplyTemplate - Apply a permission template to one project.
// The project id or project key must be provided.
// The template id or name must be provided.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) ApplyTemplate(ctx context.Context, r permissions.ApplyTemplateRequest) error {
	u := fmt.Sprintf("%s/apply_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// BulkApplyTemplate - Apply a permission template to several components. Managed projects will be ignored.
// The template id or name must be provided.
// Requires the following permission: 'Administer System'.
// Since 5.5
// Changelog:
//   6.7.2: Parameter projects accepts maximum 1000 values
func (s *Permissions) BulkApplyTemplate(ctx context.Context, r permissions.BulkApplyTemplateRequest) error {
	u := fmt.Sprintf("%s/bulk_apply_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// CreateTemplate - Create a permission template.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) CreateTemplate(ctx context.Context, r permissions.CreateTemplateRequest) (*permissions.CreateTemplateResponse, error) {
	u := fmt.Sprintf("%s/create_template", s.path)
	v := new(permissions.CreateTemplateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// DeleteTemplate - Delete a permission template.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) DeleteTemplate(ctx context.Context, r permissions.DeleteTemplateRequest) error {
	u := fmt.Sprintf("%s/delete_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroup - Remove a permission from a group.
//  This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.
//  The group name must be provided.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//
// Since 5.2
// Changelog:
//   10.0: Parameter 'groupId' is removed. Use 'groupName' instead.
//   8.4: Parameter 'groupId' is deprecated. Format changes from integer to string. Use 'groupName' instead.
func (s *Permissions) RemoveGroup(ctx context.Context, r permissions.RemoveGroupRequest) error {
	u := fmt.Sprintf("%s/remove_group", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupFromTemplate - Remove a group from a permission template.
//  The group name must be provided.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
//   10.0: Parameter 'groupId' is removed. Use 'groupName' instead.
//   8.4: Parameter 'groupId' is deprecated. Format changes from integer to string. Use 'groupName' instead.
func (s *Permissions) RemoveGroupFromTemplate(ctx context.Context, r permissions.RemoveGroupFromTemplateRequest) error {
	u := fmt.Sprintf("%s/remove_group_from_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveProjectCreatorFromTemplate - Remove a project creator from a permission template.
// Requires the following permission: 'Administer System'.
// Since 6.0
// Changelog:
func (s *Permissions) RemoveProjectCreatorFromTemplate(ctx context.Context, r permissions.RemoveProjectCreatorFromTemplateRequest) error {
	u := fmt.Sprintf("%s/remove_project_creator_from_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveUser - Remove permission from a user.
//  This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.
//  Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//
// Since 5.2
// Changelog:
func (s *Permissions) RemoveUser(ctx context.Context, r permissions.RemoveUserRequest) error {
	u := fmt.Sprintf("%s/remove_user", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveUserFromTemplate - Remove a user from a permission template.
//  Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) RemoveUserFromTemplate(ctx context.Context, r permissions.RemoveUserFromTemplateRequest) error {
	u := fmt.Sprintf("%s/remove_user_from_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SearchTemplates - List permission templates.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) SearchTemplates(ctx context.Context, r permissions.SearchTemplatesRequest) (*permissions.SearchTemplatesResponse, error) {
	u := fmt.Sprintf("%s/search_templates", s.path)
	v := new(permissions.SearchTemplatesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// SetDefaultTemplate - Set a permission template as default.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) SetDefaultTemplate(ctx context.Context, r permissions.SetDefaultTemplateRequest) error {
	u := fmt.Sprintf("%s/set_default_template", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTemplate - Update a permission template.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Changelog:
func (s *Permissions) UpdateTemplate(ctx context.Context, r permissions.UpdateTemplateRequest) (*permissions.UpdateTemplateResponse, error) {
	u := fmt.Sprintf("%s/update_template", s.path)
	v := new(permissions.UpdateTemplateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
