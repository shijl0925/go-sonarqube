package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/applications"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Applications service

// AddProject - Add a project to an application.
// Requires 'Administrator' permission on the application
// Since 7.3
// Changelog:
func (s *Applications) AddProject(ctx context.Context, r applications.AddProjectRequest) error {
	u := fmt.Sprintf("%s/applications/add_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Create - Create a new application.
// Requires 'Administer System' permission or 'Create Applications' permission
// Since 7.3
// Changelog:
//   7.4: This web service is using the 'Create Applications' permission
func (s *Applications) Create(ctx context.Context, r applications.CreateRequest) (*applications.CreateResponse, error) {
	u := fmt.Sprintf("%s/applications/create", API)
	v := new(applications.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// CreateBranch - Create a new branch on a given application.
// Requires 'Administrator' permission on the application and 'Browse' permission on its child projects
// Since 7.3
// Changelog:
func (s *Applications) CreateBranch(ctx context.Context, r applications.CreateBranchRequest) error {
	u := fmt.Sprintf("%s/applications/create_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Delete - Delete an application definition.
// Requires 'Administrator' permission on the application
// Since 7.3
// Changelog:
func (s *Applications) Delete(ctx context.Context, r applications.DeleteRequest) error {
	u := fmt.Sprintf("%s/applications/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// DeleteBranch - Delete a branch on a given application.
// Requires 'Administrator' permission on the application
// Since 7.3
// Changelog:
func (s *Applications) DeleteBranch(ctx context.Context, r applications.DeleteBranchRequest) error {
	u := fmt.Sprintf("%s/applications/delete_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveProject - Remove a project from an application
// Requires 'Administrator' permission on the application
// Since 7.3
// Changelog:
func (s *Applications) RemoveProject(ctx context.Context, r applications.RemoveProjectRequest) error {
	u := fmt.Sprintf("%s/applications/remove_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetTags - Set tags on a application.
// Requires the following permission: 'Administer' rights on the specified application
// Since 8.3
// Changelog:
func (s *Applications) SetTags(ctx context.Context, r applications.SetTagsRequest) error {
	u := fmt.Sprintf("%s/applications/set_tags", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Show - Returns an application and its associated projects.
// Requires the 'Browse' permission on the application and on its child projects.
func (s *Applications) Show(ctx context.Context, r applications.ShowRequest) (*applications.ShowResponse, error) {
	u := fmt.Sprintf("%s/applications/show", API)
	v := new(applications.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Update - Update an application.
// Requires 'Administrator' permission on the application
// Since 7.3
// Changelog:
func (s *Applications) Update(ctx context.Context, r applications.UpdateRequest) error {
	u := fmt.Sprintf("%s/applications/update", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBranch - Update a branch on a given application.
// Requires 'Administrator' permission on the application and 'Browse' permission on its child projects
// Since 7.3
// Changelog:
func (s *Applications) UpdateBranch(ctx context.Context, r applications.UpdateBranchRequest) error {
	u := fmt.Sprintf("%s/applications/update_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
