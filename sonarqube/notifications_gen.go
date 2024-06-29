package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/notifications"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Notifications service

func (s *Notifications) Add(r notifications.AddRequest) error {
	u := fmt.Sprintf("%s/notifications/add", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Notifications) List(r notifications.ListRequest) (*notifications.ListResponse, error) {
	u := fmt.Sprintf("%s/notifications/list", API)
	v := new(notifications.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Notifications) Remove(r notifications.RemoveRequest) error {
	u := fmt.Sprintf("%s/notifications/remove", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
