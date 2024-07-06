package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_branches"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectBranches service

// Delete - Delete a non-main branch of a project or application.
// Requires 'Administer' rights on the specified project or application.
// Since 6.6
// Changelog:
func (s *ProjectBranches) Delete(ctx context.Context, r project_branches.DeleteRequest) error {
	u := fmt.Sprintf("%s/delete", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// List - List the branches of a project or application.
// Requires 'Browse' or 'Execute analysis' rights on the specified project or application.
// Since 6.6
// Changelog:
//   10.6: Field 'branchId' added to the response
//   7.2: Application can be used on this web service
func (s *ProjectBranches) List(ctx context.Context, r project_branches.ListRequest) (*project_branches.ListResponse, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(project_branches.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Rename - Rename the main branch of a project or application.
// Requires 'Administer' permission on the specified project or application.
// Since 6.6
// Changelog:
func (s *ProjectBranches) Rename(ctx context.Context, r project_branches.RenameRequest) error {
	u := fmt.Sprintf("%s/rename", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetAutomaticDeletionProtection - Protect a specific branch from automatic deletion. Protection can't be disabled for the main branch.
// Requires 'Administer' permission on the specified project or application.
// Since 8.1
// Changelog:
func (s *ProjectBranches) SetAutomaticDeletionProtection(ctx context.Context, r project_branches.SetAutomaticDeletionProtectionRequest) error {
	u := fmt.Sprintf("%s/set_automatic_deletion_protection", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetMain - Allow to set a new main branch.
// . Caution, only applicable on projects.
// Requires 'Administer' rights on the specified project or application.
// Since 10.2
// Changelog:
func (s *ProjectBranches) SetMain(ctx context.Context, r project_branches.SetMainRequest) error {
	u := fmt.Sprintf("%s/set_main", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
