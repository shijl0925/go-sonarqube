package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/notifications"
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
func (s *Notifications) Add(ctx context.Context, r notifications.AddRequest) error {
	u := fmt.Sprintf("%s/notifications/add", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// List - List notifications of the authenticated user.
// Requires one of the following permissions:
//    * Authentication if no login is provided
//    * System administration if a login is provided
//
// Since 6.3
// Changelog:
func (s *Notifications) List(ctx context.Context, r notifications.ListRequest) (*notifications.ListResponse, error) {
	u := fmt.Sprintf("%s/notifications/list", API)
	v := new(notifications.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Remove - Remove a notification for the authenticated user.
// Requires one of the following permissions:
//    * Authentication if no login is provided
//    * System administration if a login is provided
//
// Since 6.3
// Changelog:
func (s *Notifications) Remove(ctx context.Context, r notifications.RemoveRequest) error {
	u := fmt.Sprintf("%s/notifications/remove", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
