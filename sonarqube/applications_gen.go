package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/applications"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Applications service

func (s *Applications) AddProject(r applications.AddProjectRequest) error {
	u := fmt.Sprintf("%s/applications/add_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) Create(r applications.CreateRequest) (*applications.CreateResponse, error) {
	u := fmt.Sprintf("%s/applications/create", API)
	v := new(applications.CreateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Applications) CreateBranch(r applications.CreateBranchRequest) error {
	u := fmt.Sprintf("%s/applications/create_branch", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) Delete(r applications.DeleteRequest) error {
	u := fmt.Sprintf("%s/applications/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) DeleteBranch(r applications.DeleteBranchRequest) error {
	u := fmt.Sprintf("%s/applications/delete_branch", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) RemoveProject(r applications.RemoveProjectRequest) error {
	u := fmt.Sprintf("%s/applications/remove_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) SetTags(r applications.SetTagsRequest) error {
	u := fmt.Sprintf("%s/applications/set_tags", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) Show(r applications.ShowRequest) (*applications.ShowResponse, error) {
	u := fmt.Sprintf("%s/applications/show", API)
	v := new(applications.ShowResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Applications) Update(r applications.UpdateRequest) error {
	u := fmt.Sprintf("%s/applications/update", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Applications) UpdateBranch(r applications.UpdateBranchRequest) error {
	u := fmt.Sprintf("%s/applications/update_branch", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
