package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/sources"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Sources service

func (s *Sources) Raw(r sources.RawRequest) (*sources.RawResponse, error) {
	u := fmt.Sprintf("%s/sources/raw", API)
	v := new(sources.RawResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Sources) Scm(r sources.ScmRequest) (*sources.ScmResponse, error) {
	u := fmt.Sprintf("%s/sources/scm", API)
	v := new(sources.ScmResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Sources) Show(r sources.ShowRequest) (*sources.ShowResponse, error) {
	u := fmt.Sprintf("%s/sources/show", API)
	v := new(sources.ShowResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
