package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/sources"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Sources service

// Raw - Get source code as raw text. Require 'See Source Code' permission on file
func (s *Sources) Raw(ctx context.Context, r sources.RawRequest) (*sources.RawResponse, error) {
	u := fmt.Sprintf("%s/sources/raw", API)
	v := new(sources.RawResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Scm - Get SCM information of source files. Require See Source Code permission on file's project
// Each element of the result array is composed of:<ol> * Line number
//  * Author of the commit
//  * Datetime of the commit (before 5.2 it was only the Date)
//  * Revision of the commit (added in 5.2)
// </ol>
func (s *Sources) Scm(ctx context.Context, r sources.ScmRequest) (*sources.ScmResponse, error) {
	u := fmt.Sprintf("%s/sources/scm", API)
	v := new(sources.ScmResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Show - Get source code. Requires See Source Code permission on file's project
// Each element of the result array is composed of:<ol> * Line number
//  * Content of the line
// </ol>
func (s *Sources) Show(ctx context.Context, r sources.ShowRequest) (*sources.ShowResponse, error) {
	u := fmt.Sprintf("%s/sources/show", API)
	v := new(sources.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
