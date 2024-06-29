package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/ce"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Ce service

func (s *Ce) Activity(r ce.ActivityRequest, p paging.Params) (*ce.ActivityResponse, error) {
	u := fmt.Sprintf("%s/ce/activity", API)
	v := new(ce.ActivityResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Ce) ActivityAll(r ce.ActivityRequest) (*ce.ActivityResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &ce.ActivityResponseAll{}
	for {
		res, err := s.Activity(r, p)
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

func (s *Ce) ActivityStatus(r ce.ActivityStatusRequest) (*ce.ActivityStatusResponse, error) {
	u := fmt.Sprintf("%s/ce/activity_status", API)
	v := new(ce.ActivityStatusResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Ce) Component(r ce.ComponentRequest) (*ce.ComponentResponse, error) {
	u := fmt.Sprintf("%s/ce/component", API)
	v := new(ce.ComponentResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Ce) Task(r ce.TaskRequest) (*ce.TaskResponse, error) {
	u := fmt.Sprintf("%s/ce/task", API)
	v := new(ce.TaskResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
