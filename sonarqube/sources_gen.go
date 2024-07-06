package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/sources"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Sources service

// Raw - Get source code as raw text. Require 'See Source Code' permission on file
// Since 5.0
func (s *Sources) Raw(ctx context.Context, r sources.RawRequest) (*sources.RawResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/raw", s.path)
	v := new(sources.RawResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Scm - Get SCM information of source files. Require See Source Code permission on file's project
// Each element of the result array is composed of:
//  * Line number
//  * Author of the commit
//  * Datetime of the commit (before 5.2 it was only the Date)
//  * Revision of the commit (added in 5.2)
//
// Since 4.4
func (s *Sources) Scm(ctx context.Context, r sources.ScmRequest) (*sources.ScmResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/scm", s.path)
	v := new(sources.ScmResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Show - Get source code. Requires See Source Code permission on file's project
// Each element of the result array is composed of:
//  * Line number
//  * Content of the line
//
// Since 4.4
func (s *Sources) Show(ctx context.Context, r sources.ShowRequest) (*sources.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(sources.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
