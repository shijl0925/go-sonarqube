package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_branches"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectBranches service

// Delete - Delete a non-main branch of a project or application.
// Requires 'Administer' rights on the specified project or application.
// Since 6.6
func (s *ProjectBranches) Delete(ctx context.Context, r project_branches.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// List - List the branches of a project or application.
// Requires 'Browse' or 'Execute analysis' rights on the specified project or application.
// Since 6.6
// Changelog:
//
//	10.6: Field 'branchId' added to the response
//	7.2: Application can be used on this web service
func (s *ProjectBranches) List(ctx context.Context, r project_branches.ListRequest) (*project_branches.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(project_branches.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Rename - Rename the main branch of a project or application.
// Requires 'Administer' permission on the specified project or application.
// Since 6.6
func (s *ProjectBranches) Rename(ctx context.Context, r project_branches.RenameRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/rename", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetAutomaticDeletionProtection - Protect a specific branch from automatic deletion. Protection can't be disabled for the main branch.
// Requires 'Administer' permission on the specified project or application.
// Since 8.1
func (s *ProjectBranches) SetAutomaticDeletionProtection(ctx context.Context, r project_branches.SetAutomaticDeletionProtectionRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_automatic_deletion_protection", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetMain - Allow to set a new main branch.
// . Caution, only applicable on projects.
// Requires 'Administer' rights on the specified project or application.
// Since 10.2
func (s *ProjectBranches) SetMain(ctx context.Context, r project_branches.SetMainRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_main", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
