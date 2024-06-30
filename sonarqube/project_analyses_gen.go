package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_analyses"
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
// Changelog:
func (s *ProjectAnalyses) CreateEvent(ctx context.Context, r project_analyses.CreateEventRequest) (*project_analyses.CreateEventResponse, error) {
	u := fmt.Sprintf("%s/project_analyses/create_event", API)
	v := new(project_analyses.CreateEventResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Delete - Delete a project analysis.
// Requires one of the following permissions:
//    * 'Administer System'
//    * 'Administer' rights on the project of the specified analysis
//
// Since 6.3
// Changelog:
func (s *ProjectAnalyses) Delete(ctx context.Context, r project_analyses.DeleteRequest) error {
	u := fmt.Sprintf("%s/project_analyses/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEvent - Delete a project analysis event.
// Only event of category 'VERSION' and 'OTHER' can be deleted.
// Requires one of the following permissions:
//    * 'Administer System'
//    * 'Administer' rights on the specified project
//
// Since 6.3
// Changelog:
func (s *ProjectAnalyses) DeleteEvent(ctx context.Context, r project_analyses.DeleteEventRequest) error {
	u := fmt.Sprintf("%s/project_analyses/delete_event", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Search - Search a project analyses and attached events.
// Requires the following permission: 'Browse' on the specified project.
// For applications, it also requires 'Browse' permission on its child projects.
// Since 6.3
// Changelog:
//   10.3: Add response field 'qualityProfile' for events related to quality profile changes
//   9.0: Add response field 'detectedCI'
//   7.5: Add QualityGate information on Applications
func (s *ProjectAnalyses) Search(ctx context.Context, r project_analyses.SearchRequest, p paging.Params) (*project_analyses.SearchResponse, error) {
	u := fmt.Sprintf("%s/project_analyses/search", API)
	v := new(project_analyses.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectAnalyses) SearchAll(ctx context.Context, r project_analyses.SearchRequest) (*project_analyses.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &project_analyses.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to project_analyses.Search: %+v", err)
		}
		response.Analyses = append(response.Analyses, res.Analyses...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
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
// Changelog:
func (s *ProjectAnalyses) UpdateEvent(ctx context.Context, r project_analyses.UpdateEventRequest) (*project_analyses.UpdateEventResponse, error) {
	u := fmt.Sprintf("%s/project_analyses/update_event", API)
	v := new(project_analyses.UpdateEventResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
