package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/webservices"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Webservices service

// List - List web services
// Since 4.2
// Changelog:
func (s *Webservices) List(ctx context.Context, r webservices.ListRequest) (*webservices.ListResponse, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(webservices.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ResponseExample - Display web service response example
// Since 4.4
// Changelog:
func (s *Webservices) ResponseExample(ctx context.Context, r webservices.ResponseExampleRequest) (*webservices.ResponseExampleResponse, error) {
	u := fmt.Sprintf("%s/response_example", s.path)
	v := new(webservices.ResponseExampleResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
