package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/languages"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Languages service

func (s *Languages) List(r languages.ListRequest) (*languages.ListResponse, error) {
	u := fmt.Sprintf("%s/languages/list", API)
	v := new(languages.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
