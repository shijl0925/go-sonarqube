package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/users"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Users service

func (s *Users) Anonymize(ctx context.Context, r users.AnonymizeRequest) error {
	u := fmt.Sprintf("%s/users/anonymize", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Users) ChangePassword(ctx context.Context, r users.ChangePasswordRequest) error {
	u := fmt.Sprintf("%s/users/change_password", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Users) Create(ctx context.Context, r users.CreateRequest) (*users.CreateResponse, error) {
	u := fmt.Sprintf("%s/users/create", API)
	v := new(users.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Users) Deactivate(ctx context.Context, r users.DeactivateRequest) (*users.DeactivateResponse, error) {
	u := fmt.Sprintf("%s/users/deactivate", API)
	v := new(users.DeactivateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Users) Groups(ctx context.Context, r users.GroupsRequest, p paging.Params) (*users.GroupsResponse, error) {
	u := fmt.Sprintf("%s/users/groups", API)
	v := new(users.GroupsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Users) GroupsAll(ctx context.Context, r users.GroupsRequest) (*users.GroupsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &users.GroupsResponseAll{}
	for {
		res, err := s.Groups(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to users.Groups: %+v", err)
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

func (s *Users) Search(ctx context.Context, r users.SearchRequest, p paging.Params) (*users.SearchResponse, error) {
	u := fmt.Sprintf("%s/users/search", API)
	v := new(users.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Users) SearchAll(ctx context.Context, r users.SearchRequest) (*users.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &users.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to users.Search: %+v", err)
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

func (s *Users) Update(ctx context.Context, r users.UpdateRequest) (*users.UpdateResponse, error) {
	u := fmt.Sprintf("%s/users/update", API)
	v := new(users.UpdateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Users) UpdateIdentityProvider(ctx context.Context, r users.UpdateIdentityProviderRequest) error {
	u := fmt.Sprintf("%s/users/update_identity_provider", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Users) UpdateLogin(ctx context.Context, r users.UpdateLoginRequest) error {
	u := fmt.Sprintf("%s/users/update_login", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
