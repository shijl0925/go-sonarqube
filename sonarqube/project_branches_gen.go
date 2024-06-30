package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_branches"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectBranches service

func (s *ProjectBranches) Delete(ctx context.Context, r project_branches.DeleteRequest) error {
	u := fmt.Sprintf("%s/project_branches/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectBranches) List(ctx context.Context, r project_branches.ListRequest) (*project_branches.ListResponse, error) {
	u := fmt.Sprintf("%s/project_branches/list", API)
	v := new(project_branches.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectBranches) Rename(ctx context.Context, r project_branches.RenameRequest) error {
	u := fmt.Sprintf("%s/project_branches/rename", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectBranches) SetAutomaticDeletionProtection(ctx context.Context, r project_branches.SetAutomaticDeletionProtectionRequest) error {
	u := fmt.Sprintf("%s/project_branches/set_automatic_deletion_protection", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectBranches) SetMain(ctx context.Context, r project_branches.SetMainRequest) error {
	u := fmt.Sprintf("%s/project_branches/set_main", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
