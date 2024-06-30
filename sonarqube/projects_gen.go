package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/projects"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Projects service

// BulkDelete - Delete one or several projects.
// Only the 1'000 first items in project filters are taken into account.
// Requires 'Administer System' permission.
// At least one parameter is required among analyzedBefore, projects and q
// Since 5.2
// Changelog:
//   9.1: The parameter 'analyzedBefore' takes into account the analysis of all branches and pull requests, not only the main branch.
//   7.8: At least one parameter is required among analyzedBefore, projects and q
func (s *Projects) BulkDelete(ctx context.Context, r projects.BulkDeleteRequest) error {
	u := fmt.Sprintf("%s/projects/bulk_delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Create - Create a project.
// If your project is hosted on a DevOps Platform, please use the import endpoint under api/alm_integrations, so it creates and properly configures the project.Requires 'Create Projects' permission.
//
// Since 4.0
// Changelog:
//   9.8: Field 'mainBranch' added to the request
//   7.1: The 'visibility' parameter is public
func (s *Projects) Create(ctx context.Context, r projects.CreateRequest) (*projects.CreateResponse, error) {
	u := fmt.Sprintf("%s/projects/create", API)
	v := new(projects.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Delete - Delete a project.
// Requires 'Administer System' permission or 'Administer' permission on the project.
// Since 5.2
// Changelog:
func (s *Projects) Delete(ctx context.Context, r projects.DeleteRequest) error {
	u := fmt.Sprintf("%s/projects/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// ExportFindings - Export all findings (issues and hotspots) of a specific project branch.
// Requires 'Administer System' permission. Keep in mind that this endpoint will return all findings, issues and hotspots (no filter), which can take time and use a lot of resources on the SonarQube server side and put pressure on the database until completion. This endpoint can be used to feed third party systems. Either the branch key or the pull request key should be specified, and not both at the same time.
func (s *Projects) ExportFindings(ctx context.Context, r projects.ExportFindingsRequest) (*projects.ExportFindingsResponse, error) {
	u := fmt.Sprintf("%s/projects/export_findings", API)
	v := new(projects.ExportFindingsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// LicenseUsage - Help admins to understand how much each project affects the total number of lines of code. Returns the list of projects together with information about their usage, sorted by lines of code descending.
// Requires Administer System permission.
func (s *Projects) LicenseUsage(ctx context.Context, r projects.LicenseUsageRequest) (*projects.LicenseUsageResponse, error) {
	u := fmt.Sprintf("%s/projects/license_usage", API)
	v := new(projects.LicenseUsageResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Search - Search for projects or views to administrate them.
//
//
//    * The response field 'lastAnalysisDate' takes into account the analysis of all branches and pull requests, not only the main branch.
//
//    * The response field 'revision' takes into account the analysis of the main branch only.
//
//
// Requires 'Administer System' permission
func (s *Projects) Search(ctx context.Context, r projects.SearchRequest, p paging.Params) (*projects.SearchResponse, error) {
	u := fmt.Sprintf("%s/projects/search", API)
	v := new(projects.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Projects) SearchAll(ctx context.Context, r projects.SearchRequest) (*projects.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &projects.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to projects.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// UpdateKey - Update a project all its sub-components keys.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified project
//
// Since 6.1
// Changelog:
//   7.1: Ability to update key of a disabled module
func (s *Projects) UpdateKey(ctx context.Context, r projects.UpdateKeyRequest) error {
	u := fmt.Sprintf("%s/projects/update_key", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateVisibility - Updates visibility of a project, application or a portfolio.
// Requires 'Project administer' permission on the specified entity
// Since 6.4
// Changelog:
func (s *Projects) UpdateVisibility(ctx context.Context, r projects.UpdateVisibilityRequest) error {
	u := fmt.Sprintf("%s/projects/update_visibility", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
