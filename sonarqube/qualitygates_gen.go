package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/qualitygates"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Qualitygates service

// AddGroup - Allow a group of users to edit a Quality Gate.
// Requires one of the following permissions:
//    * 'Administer Quality Gates'
//    * Edit right on the specified quality gate
//
// Since 9.2
func (s *Qualitygates) AddGroup(ctx context.Context, r qualitygates.AddGroupRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_group", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// AddUser - Allow a user to edit a Quality Gate.
// Requires one of the following permissions:
//    * 'Administer Quality Gates'
//    * Edit right on the specified quality gate
//
// Since 9.2
func (s *Qualitygates) AddUser(ctx context.Context, r qualitygates.AddUserRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_user", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Copy - Copy a Quality Gate.
// 'sourceName' must be provided. Requires the 'Administer Quality Gates' permission.
// Since 4.3
// Changelog:
//
//	10.0: Field 'id' in the response is deprecated
//	10.0: Parameter 'id' is removed. Use 'sourceName' instead.
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'sourceName' instead.
//	8.4: Parameter 'sourceName' added
func (s *Qualitygates) Copy(ctx context.Context, r qualitygates.CopyRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/copy", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Create - Create a Quality Gate.
// Requires the 'Administer Quality Gates' permission.
// Since 4.3
// Changelog:
//
//	10.0: Field 'id' in the response is removed.
//	8.4: Field 'id' in the response is deprecated. Format changes from integer to string.
func (s *Qualitygates) Create(ctx context.Context, r qualitygates.CreateRequest) (*qualitygates.CreateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)
	v := new(qualitygates.CreateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// CreateCondition - Add a new condition to a quality gate.
// Parameter 'gateName' must be provided. Requires the 'Administer Quality Gates' permission.
// Since 4.3
// Changelog:
//
//	10.0: Parameter 'gateId' is removed. Use 'gateName' instead.
//	8.4: Parameter 'gateName' added
//	8.4: Parameter 'gateId' is deprecated. Use 'gateName' instead.
//	7.6: Removed optional 'warning' and 'period' parameters
//	7.6: Made 'error' parameter mandatory
//	7.6: Reduced the possible values of 'op' parameter to LT and GT
func (s *Qualitygates) CreateCondition(ctx context.Context, r qualitygates.CreateConditionRequest) (*qualitygates.CreateConditionResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create_condition", s.path)
	v := new(qualitygates.CreateConditionResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// DeleteCondition - Delete a condition from a quality gate.
// Requires the 'Administer Quality Gates' permission.
// Since 4.3
func (s *Qualitygates) DeleteCondition(ctx context.Context, r qualitygates.DeleteConditionRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete_condition", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deselect - Remove the association of a project from a quality gate.
// Requires one of the following permissions:
//  * 'Administer Quality Gates'
//  * 'Administer' rights on the project
//
// Since 4.3
// Changelog:
//
//	8.3: The parameter 'projectId' was removed
//	6.6: The parameter 'gateId' was removed
func (s *Qualitygates) Deselect(ctx context.Context, r qualitygates.DeselectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/deselect", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Destroy - Delete a Quality Gate.
// Parameter 'name' must be specified. Requires the 'Administer Quality Gates' permission.
// Since 4.3
// Changelog:
//
//	10.0: Parameter 'id' is removed. Use 'name' instead.
//	8.4: Parameter 'name' added
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'name' instead.
func (s *Qualitygates) Destroy(ctx context.Context, r qualitygates.DestroyRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/destroy", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetByProject - Get the quality gate of a project.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//  * 'Browse' on the specified project
//
// Since 6.1
// Changelog:
//
//	10.0: Field 'id' in the response has been removed
//	8.4: Field 'id' in the response is deprecated. Format changes from integer to string.
//	6.6: The parameter 'projectId' has been removed
//	6.6: The parameter 'projectKey' has been renamed to 'project'
//	6.6: This webservice is now part of the public API
func (s *Qualitygates) GetByProject(ctx context.Context, r qualitygates.GetByProjectRequest) (*qualitygates.GetByProjectResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/get_by_project", s.path)
	v := new(qualitygates.GetByProjectResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// List - Get a list of quality gates
// Since 4.3
// Changelog:
//
//	10.0: Field 'default' in the response has been removed
//	10.0: Field 'id' in the response has been removed
//	9.9: 'caycStatus' field is added on quality gate
//	8.4: Field 'id' in the response is deprecated. Format changes from integer to string.
//	7.0: 'isDefault' field is added on quality gate
//	7.0: 'default' field on root level is deprecated
//	7.0: 'isBuiltIn' field is added in the response
//	7.0: 'actions' fields are added in the response
func (s *Qualitygates) List(ctx context.Context, r qualitygates.ListRequest) (*qualitygates.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(qualitygates.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// ProjectStatus - Get the quality gate status of a project or a Compute Engine task.
// Either 'analysisId', 'projectId' or 'projectKey' must be provided
// The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.
// Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//  * 'Browse' on the specified project
//  * 'Execute Analysis' on the specified project
//
// Since 5.3
// Changelog:
//
//	10.0: The fields 'periods' and 'periodIndex' in the response are removed
//	9.9: 'caycStatus' field is added to the response
//	9.5: The 'Execute Analysis' permission also allows to access the endpoint
//	8.5: The field 'periods' in the response is deprecated. Use 'period' instead
//	7.7: The parameters 'branch' and 'pullRequest' were added
//	7.6: The field 'warning' in the response is deprecated
//	6.4: The field 'ignoredConditions' is added to the response
func (s *Qualitygates) ProjectStatus(ctx context.Context, r qualitygates.ProjectStatusRequest) (*qualitygates.ProjectStatusResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/project_status", s.path)
	v := new(qualitygates.ProjectStatusResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// RemoveGroup - Remove the ability from a group to edit a Quality Gate.
// Requires one of the following permissions:
//    * 'Administer Quality Gates'
//    * Edit right on the specified quality gate
//
// Since 9.2
func (s *Qualitygates) RemoveGroup(ctx context.Context, r qualitygates.RemoveGroupRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_group", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveUser - Remove the ability from an user to edit a Quality Gate.
// Requires one of the following permissions:
//    * 'Administer Quality Gates'
//    * Edit right on the specified quality gate
//
// Since 9.2
func (s *Qualitygates) RemoveUser(ctx context.Context, r qualitygates.RemoveUserRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_user", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Rename - Rename a Quality Gate.
// 'currentName' must be specified. Requires the 'Administer Quality Gates' permission.
// Since 4.3
// Changelog:
//
//	10.0: Field 'id' in the response is deprecated
//	10.0: Parameter 'id' is removed. Use 'currentName' instead.
//	8.4: Parameter 'currentName' added
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'currentName' instead.
func (s *Qualitygates) Rename(ctx context.Context, r qualitygates.RenameRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/rename", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Search - Search for projects associated (or not) to a quality gate.
// Only authorized projects for the current user will be returned.
// Since 4.3
// Changelog:
//
//	10.0: deprecated 'more' response field has been removed
//	10.0: Parameter 'gateId' is removed. Use 'gateName' instead.
//	8.4: Parameter 'gateName' added
//	8.4: Parameter 'gateId' is deprecated. Format changes from integer to string. Use 'gateName' instead.
//	7.9: New field 'paging' in response
//	7.9: New field 'key' returning the project key in 'results' response
//	7.9: Field 'more' is deprecated in the response
func (s *Qualitygates) Search(ctx context.Context, r qualitygates.SearchRequest) (*qualitygates.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(qualitygates.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// SearchGroups - List the groups that are allowed to edit a Quality Gate.
// Requires one of the following permissions:
//    * 'Administer Quality Gates'
//    * Edit right on the specified quality gate
//
// Since 9.2
func (s *Qualitygates) SearchGroups(ctx context.Context, r qualitygates.SearchGroupsRequest, p paging.Params) (*qualitygates.SearchGroupsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search_groups", s.path)
	v := new(qualitygates.SearchGroupsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Qualitygates) SearchGroupsAll(ctx context.Context, r qualitygates.SearchGroupsRequest) (*qualitygates.SearchGroupsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualitygates.SearchGroupsResponseAll{}
	for {
		res, _, err := s.SearchGroups(ctx, r, p)
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

// SearchUsers - List the users that are allowed to edit a Quality Gate.
// Requires one of the following permissions:
//    * 'Administer Quality Gates'
//    * Edit right on the specified quality gate
//
// Since 9.2
func (s *Qualitygates) SearchUsers(ctx context.Context, r qualitygates.SearchUsersRequest, p paging.Params) (*qualitygates.SearchUsersResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search_users", s.path)
	v := new(qualitygates.SearchUsersResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Qualitygates) SearchUsersAll(ctx context.Context, r qualitygates.SearchUsersRequest) (*qualitygates.SearchUsersResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualitygates.SearchUsersResponseAll{}
	for {
		res, _, err := s.SearchUsers(ctx, r, p)
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

// Select - Associate a project to a quality gate.
// Requires one of the following permissions:
//    * 'Administer Quality Gates'
//    * 'Administer' right on the specified project
//
// Since 4.3
// Changelog:
//
//	10.0: Parameter 'gateId' is removed. Use 'gateName' instead.
//	8.4: Parameter 'gateName' added
//	8.4: Parameter 'gateId' is deprecated. Format changes from integer to string. Use 'gateName' instead.
//	8.3: The parameter 'projectId' was removed
func (s *Qualitygates) Select(ctx context.Context, r qualitygates.SelectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/select", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetAsDefault - Set a quality gate as the default quality gate.
// Parameter 'name' must be specified. Requires the 'Administer Quality Gates' permission.
// Since 4.3
// Changelog:
//
//	10.0: Parameter 'id' is removed. Use 'name' instead.
//	8.4: Parameter 'name' added
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'name' instead.
func (s *Qualitygates) SetAsDefault(ctx context.Context, r qualitygates.SetAsDefaultRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_as_default", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Show - Display the details of a quality gate
// Since 4.3
// Changelog:
//
//	10.3: 'isDefault' field is added to the response
//	10.0: Field 'id' in the response has been removed
//	10.0: Parameter 'id' is removed. Use 'name' instead.
//	9.9: 'caycStatus' field is added to the response
//	8.4: Parameter 'id' is deprecated. Format changes from integer to string. Use 'name' instead.
//	8.4: Field 'id' in the response is deprecated.
//	7.6: 'period' and 'warning' fields of conditions are removed from the response
//	7.0: 'isBuiltIn' field is added to the response
//	7.0: 'actions' field is added in the response
func (s *Qualitygates) Show(ctx context.Context, r qualitygates.ShowRequest) (*qualitygates.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(qualitygates.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// UpdateCondition - Update a condition attached to a quality gate.
// Requires the 'Administer Quality Gates' permission.
// Since 4.3
// Changelog:
//
//	8.4: Parameter 'id' format changes from integer to string.
//	7.6: Removed optional 'warning' and 'period' parameters
//	7.6: Made 'error' parameter mandatory
//	7.6: Reduced the possible values of 'op' parameter to LT and GT
func (s *Qualitygates) UpdateCondition(ctx context.Context, r qualitygates.UpdateConditionRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update_condition", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
