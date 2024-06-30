package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/user_groups"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type UserGroups service

func (s *UserGroups) AddUser(ctx context.Context, r user_groups.AddUserRequest) error {
	u := fmt.Sprintf("%s/user_groups/add_user", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserGroups) Create(ctx context.Context, r user_groups.CreateRequest) (*user_groups.CreateResponse, error) {
	u := fmt.Sprintf("%s/user_groups/create", API)
	v := new(user_groups.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *UserGroups) Delete(ctx context.Context, r user_groups.DeleteRequest) error {
	u := fmt.Sprintf("%s/user_groups/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserGroups) RemoveUser(ctx context.Context, r user_groups.RemoveUserRequest) error {
	u := fmt.Sprintf("%s/user_groups/remove_user", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserGroups) Search(ctx context.Context, r user_groups.SearchRequest, p paging.Params) (*user_groups.SearchResponse, error) {
	u := fmt.Sprintf("%s/user_groups/search", API)
	v := new(user_groups.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *UserGroups) SearchAll(ctx context.Context, r user_groups.SearchRequest) (*user_groups.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &user_groups.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to user_groups.Search: %+v", err)
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

func (s *UserGroups) Update(ctx context.Context, r user_groups.UpdateRequest) error {
	u := fmt.Sprintf("%s/user_groups/update", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserGroups) Users(ctx context.Context, r user_groups.UsersRequest, p paging.Params) (*user_groups.UsersResponse, error) {
	u := fmt.Sprintf("%s/user_groups/users", API)
	v := new(user_groups.UsersResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *UserGroups) UsersAll(ctx context.Context, r user_groups.UsersRequest) (*user_groups.UsersResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &user_groups.UsersResponseAll{}
	for {
		res, err := s.Users(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to user_groups.Users: %+v", err)
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
