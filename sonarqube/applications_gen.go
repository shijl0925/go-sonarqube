package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/applications"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Applications service

// AddProject - Add a project to an application.
// Requires 'Administrator' permission on the application
// Since 7.3
func (s *Applications) AddProject(ctx context.Context, r applications.AddProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Create - Create a new application.
// Requires 'Administer System' permission or 'Create Applications' permission
// Since 7.3
// Changelog:
//
//	7.4: This web service is using the 'Create Applications' permission
func (s *Applications) Create(ctx context.Context, r applications.CreateRequest) (*applications.CreateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)
	v := new(applications.CreateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// CreateBranch - Create a new branch on a given application.
// Requires 'Administrator' permission on the application and 'Browse' permission on its child projects
// Since 7.3
func (s *Applications) CreateBranch(ctx context.Context, r applications.CreateBranchRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/create_branch", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Delete - Delete an application definition.
// Requires 'Administrator' permission on the application
// Since 7.3
func (s *Applications) Delete(ctx context.Context, r applications.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteBranch - Delete a branch on a given application.
// Requires 'Administrator' permission on the application
// Since 7.3
func (s *Applications) DeleteBranch(ctx context.Context, r applications.DeleteBranchRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete_branch", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveProject - Remove a project from an application
// Requires 'Administrator' permission on the application
// Since 7.3
func (s *Applications) RemoveProject(ctx context.Context, r applications.RemoveProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetTags - Set tags on a application.
// Requires the following permission: 'Administer' rights on the specified application
// Since 8.3
func (s *Applications) SetTags(ctx context.Context, r applications.SetTagsRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_tags", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Show - Returns an application and its associated projects.
// Requires the 'Browse' permission on the application and on its child projects.
// Since 7.3
// Changelog:
//
//	8.3: Change required permission from 'Admin' to 'Browse'
func (s *Applications) Show(ctx context.Context, r applications.ShowRequest) (*applications.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(applications.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Update - Update an application.
// Requires 'Administrator' permission on the application
// Since 7.3
func (s *Applications) Update(ctx context.Context, r applications.UpdateRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateBranch - Update a branch on a given application.
// Requires 'Administrator' permission on the application and 'Browse' permission on its child projects
// Since 7.3
func (s *Applications) UpdateBranch(ctx context.Context, r applications.UpdateBranchRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update_branch", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
