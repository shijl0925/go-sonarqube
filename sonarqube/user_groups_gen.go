package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/user_groups"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type UserGroups service

// AddUser - Add a user to a group.
// 'name' must be provided.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//
//	10.4: Deprecated. Use POST /api/v2/authorizations/group-memberships instead
//	10.0: Parameter 'id' is removed. Use 'name' instead.
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'name' instead.
func (s *UserGroups) AddUser(ctx context.Context, r user_groups.AddUserRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_user", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Create - Create a group.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//
//	10.4: Deprecated. Use POST /api/v2/authorizations/groups instead
//	8.4: Field 'id' format in the response changes from integer to string.
func (s *UserGroups) Create(ctx context.Context, r user_groups.CreateRequest) (*user_groups.CreateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)
	v := new(user_groups.CreateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Delete - Delete a group. The default groups cannot be deleted.
// 'name' must be provided.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//
//	10.4: Deprecated. Use DELETE /api/v2/authorizations/groups instead
//	10.0: Parameter 'id' is removed. Use 'name' instead.
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'name' instead.
func (s *UserGroups) Delete(ctx context.Context, r user_groups.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveUser - Remove a user from a group.
// 'name' must be provided.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//
//	10.4: Deprecated. Use DELETE /api/v2/authorizations/group-memberships instead
//	10.0: Parameter 'id' is removed. Use 'name' instead.
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'name' instead.
func (s *UserGroups) RemoveUser(ctx context.Context, r user_groups.RemoveUserRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_user", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Search - Search for user groups.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//
//	10.4: Deprecated. Use GET /api/v2/authorizations/groups instead
//	10.0: Field 'id' in the response has been removed
//	10.0: New parameter 'managed' to optionally search by managed status
//	10.0: Response includes 'managed' field.
//	8.4: Field 'id' in the response is deprecated. Format changes from integer to string.
//	6.4: Paging response fields moved to a Paging object
//	6.4: 'default' response field has been added
func (s *UserGroups) Search(ctx context.Context, r user_groups.SearchRequest, p paging.Params) (*user_groups.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(user_groups.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *UserGroups) SearchAll(ctx context.Context, r user_groups.SearchRequest) (*user_groups.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &user_groups.SearchResponseAll{}
	for {
		res, _, err := s.Search(ctx, r, p)
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

// Update - Update a group.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//
//	10.4: Deprecated. Use PATCH /api/v2/authorizations/groups instead
//	10.0: Parameter 'id' is removed in favor of 'currentName'
//	8.5: Parameter 'id' deprecated in favor of 'currentName'
//	8.4: Parameter 'id' format changes from integer to string
//	6.4: The default group is no longer editable
func (s *UserGroups) Update(ctx context.Context, r user_groups.UpdateRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Users - Search for users with membership information with respect to a group.
// Requires the following permission: 'Administer System'.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//
//	10.4: Deprecated. Use GET /api/v2/authorizations/group-memberships instead
//	10.0: Field 'managed' added to the payload.
//	10.0: Parameter 'id' is removed. Use 'name' instead.
//	9.8: response fields 'total', 's', 'ps' have been deprecated, please use 'paging' object instead.
//	9.8: The field 'paging' has been added to the response.
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'name' instead.
func (s *UserGroups) Users(ctx context.Context, r user_groups.UsersRequest, p paging.Params) (*user_groups.UsersResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/users", s.path)
	v := new(user_groups.UsersResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *UserGroups) UsersAll(ctx context.Context, r user_groups.UsersRequest) (*user_groups.UsersResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &user_groups.UsersResponseAll{}
	for {
		res, _, err := s.Users(ctx, r, p)
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
