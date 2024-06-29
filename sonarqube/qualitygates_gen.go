package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/qualitygates"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Qualitygates service

func (s *Qualitygates) AddGroup(r qualitygates.AddGroupRequest) error {
	u := fmt.Sprintf("%s/qualitygates/add_group", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) AddUser(r qualitygates.AddUserRequest) error {
	u := fmt.Sprintf("%s/qualitygates/add_user", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Copy(r qualitygates.CopyRequest) error {
	u := fmt.Sprintf("%s/qualitygates/copy", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Create(r qualitygates.CreateRequest) (*qualitygates.CreateResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/create", API)
	v := new(qualitygates.CreateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) CreateCondition(r qualitygates.CreateConditionRequest) (*qualitygates.CreateConditionResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/create_condition", API)
	v := new(qualitygates.CreateConditionResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) DeleteCondition(r qualitygates.DeleteConditionRequest) error {
	u := fmt.Sprintf("%s/qualitygates/delete_condition", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Deselect(r qualitygates.DeselectRequest) error {
	u := fmt.Sprintf("%s/qualitygates/deselect", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Destroy(r qualitygates.DestroyRequest) error {
	u := fmt.Sprintf("%s/qualitygates/destroy", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) GetByProject(r qualitygates.GetByProjectRequest) (*qualitygates.GetByProjectResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/get_by_project", API)
	v := new(qualitygates.GetByProjectResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) List(r qualitygates.ListRequest) (*qualitygates.ListResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/list", API)
	v := new(qualitygates.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) ProjectStatus(r qualitygates.ProjectStatusRequest) (*qualitygates.ProjectStatusResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/project_status", API)
	v := new(qualitygates.ProjectStatusResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) RemoveGroup(r qualitygates.RemoveGroupRequest) error {
	u := fmt.Sprintf("%s/qualitygates/remove_group", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) RemoveUser(r qualitygates.RemoveUserRequest) error {
	u := fmt.Sprintf("%s/qualitygates/remove_user", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Rename(r qualitygates.RenameRequest) error {
	u := fmt.Sprintf("%s/qualitygates/rename", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Search(r qualitygates.SearchRequest) (*qualitygates.SearchResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/search", API)
	v := new(qualitygates.SearchResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) SearchGroups(r qualitygates.SearchGroupsRequest, p paging.Params) (*qualitygates.SearchGroupsResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/search_groups", API)
	v := new(qualitygates.SearchGroupsResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) SearchGroupsAll(r qualitygates.SearchGroupsRequest) (*qualitygates.SearchGroupsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualitygates.SearchGroupsResponseAll{}
	for {
		res, err := s.SearchGroups(r, p)
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

func (s *Qualitygates) SearchUsers(r qualitygates.SearchUsersRequest, p paging.Params) (*qualitygates.SearchUsersResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/search_users", API)
	v := new(qualitygates.SearchUsersResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) SearchUsersAll(r qualitygates.SearchUsersRequest) (*qualitygates.SearchUsersResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualitygates.SearchUsersResponseAll{}
	for {
		res, err := s.SearchUsers(r, p)
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

func (s *Qualitygates) Select(r qualitygates.SelectRequest) error {
	u := fmt.Sprintf("%s/qualitygates/select", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) SetAsDefault(r qualitygates.SetAsDefaultRequest) error {
	u := fmt.Sprintf("%s/qualitygates/set_as_default", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Qualitygates) Show(r qualitygates.ShowRequest) (*qualitygates.ShowResponse, error) {
	u := fmt.Sprintf("%s/qualitygates/show", API)
	v := new(qualitygates.ShowResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Qualitygates) UpdateCondition(r qualitygates.UpdateConditionRequest) error {
	u := fmt.Sprintf("%s/qualitygates/update_condition", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
