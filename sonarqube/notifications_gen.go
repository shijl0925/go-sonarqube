package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/notifications"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Notifications service

// Add - Add a notification for the authenticated user.
// Requires one of the following permissions:
//   * Authentication if no login is provided. If a project is provided, requires the 'Browse' permission on the specified project.
//   * System administration if a login is provided. If a project is provided, requires the 'Browse' permission on the specified project.
//
// Since 6.3
// Changelog:
func (s *Notifications) Add(ctx context.Context, r notifications.AddRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// List - List notifications of the authenticated user.
// Requires one of the following permissions:
//    * Authentication if no login is provided
//    * System administration if a login is provided
//
// Since 6.3
// Changelog:
func (s *Notifications) List(ctx context.Context, r notifications.ListRequest) (*notifications.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(notifications.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Remove - Remove a notification for the authenticated user.
// Requires one of the following permissions:
//    * Authentication if no login is provided
//    * System administration if a login is provided
//
// Since 6.3
// Changelog:
func (s *Notifications) Remove(ctx context.Context, r notifications.RemoveRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
