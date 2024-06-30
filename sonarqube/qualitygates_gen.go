package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/qualitygates"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Qualitygates service

func (s *Qualitygates) AddGroup(ctx context.Context, r qualitygates.AddGroupRequest) error {
	u := fmt.Sprintf("%s/qualitygates/add_group", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) AddUser(ctx context.Context, r qualitygates.AddUserRequest) error {
	u := fmt.Sprintf("%s/qualitygates/add_user", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Copy(ctx context.Context, r qualitygates.CopyRequest) error {
	u := fmt.Sprintf("%s/qualitygates/copy", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Create(ctx context.Context, r qualitygates.CreateRequest) (*qualitygates.CreateResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/create", API)
	v := new(qualitygates.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) CreateCondition(ctx context.Context, r qualitygates.CreateConditionRequest) (*qualitygates.CreateConditionResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/create_condition", API)
	v := new(qualitygates.CreateConditionResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) DeleteCondition(ctx context.Context, r qualitygates.DeleteConditionRequest) error {
	u := fmt.Sprintf("%s/qualitygates/delete_condition", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Deselect(ctx context.Context, r qualitygates.DeselectRequest) error {
	u := fmt.Sprintf("%s/qualitygates/deselect", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Destroy(ctx context.Context, r qualitygates.DestroyRequest) error {
	u := fmt.Sprintf("%s/qualitygates/destroy", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) GetByProject(ctx context.Context, r qualitygates.GetByProjectRequest) (*qualitygates.GetByProjectResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/get_by_project", API)
	v := new(qualitygates.GetByProjectResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) List(ctx context.Context, r qualitygates.ListRequest) (*qualitygates.ListResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/list", API)
	v := new(qualitygates.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) ProjectStatus(ctx context.Context, r qualitygates.ProjectStatusRequest) (*qualitygates.ProjectStatusResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/project_status", API)
	v := new(qualitygates.ProjectStatusResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) RemoveGroup(ctx context.Context, r qualitygates.RemoveGroupRequest) error {
	u := fmt.Sprintf("%s/qualitygates/remove_group", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) RemoveUser(ctx context.Context, r qualitygates.RemoveUserRequest) error {
	u := fmt.Sprintf("%s/qualitygates/remove_user", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Rename(ctx context.Context, r qualitygates.RenameRequest) error {
	u := fmt.Sprintf("%s/qualitygates/rename", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Search(ctx context.Context, r qualitygates.SearchRequest) (*qualitygates.SearchResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/search", API)
	v := new(qualitygates.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) SearchGroups(ctx context.Context, r qualitygates.SearchGroupsRequest, p paging.Params) (*qualitygates.SearchGroupsResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/search_groups", API)
	v := new(qualitygates.SearchGroupsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) SearchGroupsAll(ctx context.Context, r qualitygates.SearchGroupsRequest) (*qualitygates.SearchGroupsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualitygates.SearchGroupsResponseAll{}
	for {
		res, err := s.SearchGroups(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to qualitygates.SearchGroups: %+v", err)
		}
		response.Groups = append(response.Groups, res.Groups...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Qualitygates) SearchUsers(ctx context.Context, r qualitygates.SearchUsersRequest, p paging.Params) (*qualitygates.SearchUsersResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/search_users", API)
	v := new(qualitygates.SearchUsersResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) SearchUsersAll(ctx context.Context, r qualitygates.SearchUsersRequest) (*qualitygates.SearchUsersResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualitygates.SearchUsersResponseAll{}
	for {
		res, err := s.SearchUsers(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to qualitygates.SearchUsers: %+v", err)
		}
		response.Users = append(response.Users, res.Users...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Qualitygates) Select(ctx context.Context, r qualitygates.SelectRequest) error {
	u := fmt.Sprintf("%s/qualitygates/select", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) SetAsDefault(ctx context.Context, r qualitygates.SetAsDefaultRequest) error {
	u := fmt.Sprintf("%s/qualitygates/set_as_default", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Show(ctx context.Context, r qualitygates.ShowRequest) (*qualitygates.ShowResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/show", API)
	v := new(qualitygates.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) UpdateCondition(ctx context.Context, r qualitygates.UpdateConditionRequest) error {
	u := fmt.Sprintf("%s/qualitygates/update_condition", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
