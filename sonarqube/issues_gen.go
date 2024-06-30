package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/issues"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Issues service

// AddComment - Add a comment.
// Requires authentication and the following permission: 'Browse' on the project of the specified issue.
// Since 3.6
// Changelog:
//   10.4: The response fields 'severity' and 'type' are deprecated. Please use 'impacts' instead.
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   6.5: the database ids of the components are removed from the response
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
//   6.3: the response returns the issue with all its details
func (s *Issues) AddComment(ctx context.Context, r issues.AddCommentRequest) (*issues.AddCommentResponse, error) {
	u := fmt.Sprintf("%s/issues/add_comment", API)
	v := new(issues.AddCommentResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Assign - Assign/Unassign an issue. Requires authentication and Browse permission on project
// Since 3.6
// Changelog:
//   10.4: The response fields 'severity' and 'type' are deprecated. Please use 'impacts' instead.
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   6.5: the database ids of the components are removed from the response
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
func (s *Issues) Assign(ctx context.Context, r issues.AssignRequest) (*issues.AssignResponse, error) {
	u := fmt.Sprintf("%s/issues/assign", API)
	v := new(issues.AssignResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Authors - Search SCM accounts which match a given query.
// Requires authentication.
// When issue indexing is in progress returns 503 service unavailable HTTP code.
func (s *Issues) Authors(ctx context.Context, r issues.AuthorsRequest) (*issues.AuthorsResponse, error) {
	u := fmt.Sprintf("%s/issues/authors", API)
	v := new(issues.AuthorsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// BulkChange - Bulk change on issues. Up to 500 issues can be updated.
// Requires authentication.
// Since 3.7
// Changelog:
//   10.4: Transitions 'wontfix' and 'confirm' are now deprecated. Use transition 'accept' instead. The transition 'unconfirm' is deprecated too.
//   10.4: Transition 'accept' is now supported.
//   10.2: Parameters 'set_severity' and 'set_type' are now deprecated.
//   8.2: Security hotspots are no longer supported and will be ignored.
//   8.2: Transitions 'setinreview', 'resolveasreviewed' and 'openasvulnerability' are no more supported
//   6.3: 'actions' parameter is ignored
func (s *Issues) BulkChange(ctx context.Context, r issues.BulkChangeRequest) (*issues.BulkChangeResponse, error) {
	u := fmt.Sprintf("%s/issues/bulk_change", API)
	v := new(issues.BulkChangeResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Changelog - Display changelog of an issue.
// Requires the 'Browse' permission on the project of the specified issue.
func (s *Issues) Changelog(ctx context.Context, r issues.ChangelogRequest) (*issues.ChangelogResponse, error) {
	u := fmt.Sprintf("%s/issues/changelog", API)
	v := new(issues.ChangelogResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// DeleteComment - Delete a comment.
// Requires authentication and the following permission: 'Browse' on the project of the specified issue.
// Since 3.6
// Changelog:
//   10.4: The response fields 'severity' and 'type' are deprecated. Please use 'impacts' instead.
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
//   6.5: the database ids of the components are removed from the response
//   6.3: the response returns the issue with all its details
//   6.3: the 'key' parameter is renamed 'comment'
func (s *Issues) DeleteComment(ctx context.Context, r issues.DeleteCommentRequest) (*issues.DeleteCommentResponse, error) {
	u := fmt.Sprintf("%s/issues/delete_comment", API)
	v := new(issues.DeleteCommentResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// DoTransition - Do workflow transition on an issue. Requires authentication and Browse permission on project.
//
// The transitions 'accept', 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.
//
// The transitions involving security hotspots require the permission 'Administer Security Hotspot'.
//
// Since 3.6
// Changelog:
//   10.4: The transitions 'wontfix' and 'confirm' are deprecated. Please use 'accept' instead. The transition 'unconfirm' is deprecated too.
//   10.4: Add transition 'accept'.
//   10.4: The response fields 'severity' and 'type' are deprecated. Please use 'impacts' instead.
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   8.1: transitions 'setinreview' and 'openasvulnerability' are no more supported
//   7.8: added 'setinreview', resolveasreviewed, openasvulnerability and resetastoreview transitions for security hotspots
//   7.3: added transitions for security hotspots
//   6.5: the database ids of the components are removed from the response
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
func (s *Issues) DoTransition(ctx context.Context, r issues.DoTransitionRequest) (*issues.DoTransitionResponse, error) {
	u := fmt.Sprintf("%s/issues/do_transition", API)
	v := new(issues.DoTransitionResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// EditComment - Edit a comment.
// Requires authentication and the following permission: 'Browse' on the project of the specified issue.
// Since 3.6
// Changelog:
//   10.4: The response fields 'severity' and 'type' are deprecated. Please use 'impacts' instead.
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   6.5: the database ids of the components are removed from the response
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
//   6.3: the response returns the issue with all its details
//   6.3: the 'key' parameter has been renamed comment
func (s *Issues) EditComment(ctx context.Context, r issues.EditCommentRequest) (*issues.EditCommentResponse, error) {
	u := fmt.Sprintf("%s/issues/edit_comment", API)
	v := new(issues.EditCommentResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GitlabSastExport - Return a list of vulnerabilities according to the Gitlab SAST JSON format.
// The JSON produced can be used in GitLab for generating the Vulnerability Report.Requires the 'Browse' or 'Scan' permission on the specified project.
func (s *Issues) GitlabSastExport(ctx context.Context, r issues.GitlabSastExportRequest) (*issues.GitlabSastExportResponse, error) {
	u := fmt.Sprintf("%s/issues/gitlab_sast_export", API)
	v := new(issues.GitlabSastExportResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Reindex - Reindex issues for a project.
// Require 'Administer System' permission.
// Since 9.8
// Changelog:
func (s *Issues) Reindex(ctx context.Context, r issues.ReindexRequest) error {
	u := fmt.Sprintf("%s/issues/reindex", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Search - Search for issues.
// Requires the 'Browse' permission on the specified project(s).
// For applications, it also requires 'Browse' permission on its child projects.
// When issue indexing is in progress returns 503 service unavailable HTTP code.
func (s *Issues) Search(ctx context.Context, r issues.SearchRequest, p paging.Params) (*issues.SearchResponse, error) {
	u := fmt.Sprintf("%s/issues/search", API)
	v := new(issues.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Issues) SearchAll(ctx context.Context, r issues.SearchRequest) (*issues.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &issues.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to issues.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Facets = append(response.Facets, res.Facets...)
		response.Issues = append(response.Issues, res.Issues...)
		response.Rules = append(response.Rules, res.Rules...)
		response.Users = append(response.Users, res.Users...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// SetSeverity - Change severity.
// Requires the following permissions:
//    * 'Authentication'
//    * 'Browse' rights on project of the specified issue
//    * 'Administer Issues' rights on project of the specified issue
//
// Since 3.6
// Deprecated since 10.2
// Changelog:
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: This endpoint is now deprecated.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   6.5: the database ids of the components are removed from the response
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
func (s *Issues) SetSeverity(ctx context.Context, r issues.SetSeverityRequest) (*issues.SetSeverityResponse, error) {
	u := fmt.Sprintf("%s/issues/set_severity", API)
	v := new(issues.SetSeverityResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// SetTags - Set tags on an issue.
// Requires authentication and Browse permission on project
// Since 5.1
// Changelog:
//   10.4: The response fields 'severity' and 'type' are deprecated. Please use 'impacts' instead.
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   6.5: the database ids of the components are removed from the response
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
//   6.4: response contains issue information instead of list of tags
func (s *Issues) SetTags(ctx context.Context, r issues.SetTagsRequest) (*issues.SetTagsResponse, error) {
	u := fmt.Sprintf("%s/issues/set_tags", API)
	v := new(issues.SetTagsResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// SetType - Change type of issue, for instance from 'code smell' to 'bug'.
// Requires the following permissions:
//    * 'Authentication'
//    * 'Browse' rights on project of the specified issue
//    * 'Administer Issues' rights on project of the specified issue
//
// Since 5.5
// Deprecated since 10.2
// Changelog:
//   10.4: The response fields 'status' and 'resolution' are deprecated. Please use 'issueStatus' instead.
//   10.4: Add 'issueStatus' field to the response.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   10.2: This endpoint is now deprecated.
//   9.6: Response field 'ruleDescriptionContextKey' added
//   8.8: The response field components.uuid is removed
//   6.5: the database ids of the components are removed from the response
//   6.5: the response field components.uuid is deprecated. Use components.key instead.
func (s *Issues) SetType(ctx context.Context, r issues.SetTypeRequest) (*issues.SetTypeResponse, error) {
	u := fmt.Sprintf("%s/issues/set_type", API)
	v := new(issues.SetTypeResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Tags - List tags matching a given query
func (s *Issues) Tags(ctx context.Context, r issues.TagsRequest) (*issues.TagsResponse, error) {
	u := fmt.Sprintf("%s/issues/tags", API)
	v := new(issues.TagsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
