package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/ce"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Ce service

// Activity - Search for tasks.
// Requires the system administration permission, or project administration permission if component is set.
// Since 5.2
// Changelog:
//   10.4: field "infoMessages" added to response
//   10.1: The use of module keys in parameter 'component' is removed
//   10.1: Warnings field will be now be filled (it was always empty in the past).
//   10.0: Remove deprecated field 'componentId'
//   8.8: field "logs" is dropped
//   7.6: The use of module keys in parameters 'q' is deprecated
//   7.1: field "pullRequest" added
//   6.6: fields "branch" and "branchType" added
//   6.1: field "logs" is deprecated and its value is always false
//   5.5: it's no more possible to specify the page parameter.
func (s *Ce) Activity(ctx context.Context, r ce.ActivityRequest, p paging.Params) (*ce.ActivityResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/activity", s.path)
	v := new(ce.ActivityResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Ce) ActivityAll(ctx context.Context, r ce.ActivityRequest) (*ce.ActivityResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &ce.ActivityResponseAll{}
	for {
		res, _, err := s.Activity(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to ce.Activity: %+v", err)
		}
		response.Tasks = append(response.Tasks, res.Tasks...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// ActivityStatus - Returns CE activity related metrics.
// Requires 'Administer System' permission or 'Administer' rights on the specified project.
// Since 5.5
// Changelog:
//   10.0: Remove deprecated field 'componentId'
//   8.8: Parameter 'componentId' is now deprecated.
//   8.8: Parameter 'componentKey' is now removed. Please use parameter 'component' instead.
//   7.8: New field 'pendingTime' in response, only included when there are pending tasks
//   6.6: New field 'inProgress' in response
func (s *Ce) ActivityStatus(ctx context.Context, r ce.ActivityStatusRequest) (*ce.ActivityStatusResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/activity_status", s.path)
	v := new(ce.ActivityStatusResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Component - Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).
// Requires the following permission: 'Browse' on the specified component.
// Since 5.2
// Changelog:
//   8.8: field "logs" is dropped
//   8.8: Deprecated parameter 'componentId' has been removed.
//   8.8: Parameter 'component' is now required.
//   7.6: The use of module keys in parameter "component" is deprecated
//   6.6: fields "branch" and "branchType" added
//   6.1: field "logs" is deprecated and its value is always false
func (s *Ce) Component(ctx context.Context, r ce.ComponentRequest) (*ce.ComponentResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/component", s.path)
	v := new(ce.ComponentResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Task - Give Compute Engine task details such as type, status, duration and associated component.
// Requires one of the following permissions:
//  * 'Administer' at global or project level
//  * 'Execute Analysis' at global or project level
// Since 6.1, field "logs" is deprecated and its value is always false.
// Since 5.2
// Changelog:
//   10.1: Warnings field will be now always be filled (it is not necessary to mention it explicitly in 'additionalFields'). 'additionalFields' value `warning' is deprecated.
//   10.1: 'Project Administrator' is added to the list of allowed permissions to access this endpoint
//   6.6: fields "branch" and "branchType" added
func (s *Ce) Task(ctx context.Context, r ce.TaskRequest) (*ce.TaskResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/task", s.path)
	v := new(ce.TaskResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
