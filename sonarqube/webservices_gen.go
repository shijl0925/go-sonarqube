package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/webservices"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Webservices service

func (s *Webservices) List(r webservices.ListRequest) (*webservices.ListResponse, error) {
	u := fmt.Sprintf("%s/webservices/list", API)
	v := new(webservices.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Webservices) ResponseExample(r webservices.ResponseExampleRequest) (*webservices.ResponseExampleResponse, error) {
	u := fmt.Sprintf("%s/webservices/response_example", API)
	v := new(webservices.ResponseExampleResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
