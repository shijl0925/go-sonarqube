package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/applications"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Applications service

func (s *Applications) AddProject(ctx context.Context, r applications.AddProjectRequest) error {
	u := fmt.Sprintf("%s/applications/add_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) Create(ctx context.Context, r applications.CreateRequest) (*applications.CreateResponse, error) {
	u := fmt.Sprintf("%s/applications/create", API)
	v := new(applications.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Applications) CreateBranch(ctx context.Context, r applications.CreateBranchRequest) error {
	u := fmt.Sprintf("%s/applications/create_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) Delete(ctx context.Context, r applications.DeleteRequest) error {
	u := fmt.Sprintf("%s/applications/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) DeleteBranch(ctx context.Context, r applications.DeleteBranchRequest) error {
	u := fmt.Sprintf("%s/applications/delete_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) RemoveProject(ctx context.Context, r applications.RemoveProjectRequest) error {
	u := fmt.Sprintf("%s/applications/remove_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) SetTags(ctx context.Context, r applications.SetTagsRequest) error {
	u := fmt.Sprintf("%s/applications/set_tags", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) Show(ctx context.Context, r applications.ShowRequest) (*applications.ShowResponse, error) {
	u := fmt.Sprintf("%s/applications/show", API)
	v := new(applications.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Applications) Update(ctx context.Context, r applications.UpdateRequest) error {
	u := fmt.Sprintf("%s/applications/update", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) UpdateBranch(ctx context.Context, r applications.UpdateBranchRequest) error {
	u := fmt.Sprintf("%s/applications/update_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
