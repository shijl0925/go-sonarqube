package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_analyses"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectAnalyses service

// CreateEvent - Create a project analysis event.
// Only event of category 'VERSION' and 'OTHER' can be created.
// Requires one of the following permissions:
//    * 'Administer System'
//    * 'Administer' rights on the specified project
//
// Since 6.3
func (s *ProjectAnalyses) CreateEvent(ctx context.Context, r project_analyses.CreateEventRequest) (*project_analyses.CreateEventResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create_event", s.path)
	v := new(project_analyses.CreateEventResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Delete - Delete a project analysis.
// Requires one of the following permissions:
//    * 'Administer System'
//    * 'Administer' rights on the project of the specified analysis
//
// Since 6.3
func (s *ProjectAnalyses) Delete(ctx context.Context, r project_analyses.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteEvent - Delete a project analysis event.
// Only event of category 'VERSION' and 'OTHER' can be deleted.
// Requires one of the following permissions:
//    * 'Administer System'
//    * 'Administer' rights on the specified project
//
// Since 6.3
func (s *ProjectAnalyses) DeleteEvent(ctx context.Context, r project_analyses.DeleteEventRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete_event", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Search - Search a project analyses and attached events.
// Requires the following permission: 'Browse' on the specified project.
// For applications, it also requires 'Browse' permission on its child projects.
// Since 6.3
// Changelog:
//
//	10.3: Add response field 'qualityProfile' for events related to quality profile changes
//	9.0: Add response field 'detectedCI'
//	7.5: Add QualityGate information on Applications
func (s *ProjectAnalyses) Search(ctx context.Context, r project_analyses.SearchRequest, p paging.Params) (*project_analyses.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(project_analyses.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *ProjectAnalyses) SearchAll(ctx context.Context, r project_analyses.SearchRequest) (*project_analyses.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &project_analyses.SearchResponseAll{}
	for {
		res, _, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to project_analyses.Search: %+v", err)
		}
		response.Analyses = append(response.Analyses, res.Analyses...)
		if res.GetPaging().End() {
			break
		}
		p.P++
	}
	return response, nil
}

// UpdateEvent - Update a project analysis event.
// Only events of category 'VERSION' and 'OTHER' can be updated.
// Requires one of the following permissions:
//    * 'Administer System'
//    * 'Administer' rights on the specified project
//
// Since 6.3
func (s *ProjectAnalyses) UpdateEvent(ctx context.Context, r project_analyses.UpdateEventRequest) (*project_analyses.UpdateEventResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/update_event", s.path)
	v := new(project_analyses.UpdateEventResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
