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
	u := fmt.Sprintf("%s/add_comment", s.path)
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
	u := fmt.Sprintf("%s/assign", s.path)
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
// Since 5.1
// Changelog:
//   7.4: The maximum size of 'ps' is set to 100
func (s *Issues) Authors(ctx context.Context, r issues.AuthorsRequest) (*issues.AuthorsResponse, error) {
	u := fmt.Sprintf("%s/authors", s.path)
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
	u := fmt.Sprintf("%s/bulk_change", s.path)
	v := new(issues.BulkChangeResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Changelog - Display changelog of an issue.
// Requires the 'Browse' permission on the project of the specified issue.
// Since 4.1
// Changelog:
//   10.4: 'issueStatus' key is added in the differences
//   10.4: 'status', 'resolution', 'severity' and 'type' keys are now deprecated in the differences
//   9.7: 'externalUser' and 'webhookSource' information added to the answer
//   6.3: changes on effort is expressed with the raw value in minutes (instead of the duration previously)
func (s *Issues) Changelog(ctx context.Context, r issues.ChangelogRequest) (*issues.ChangelogResponse, error) {
	u := fmt.Sprintf("%s/changelog", s.path)
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
	u := fmt.Sprintf("%s/delete_comment", s.path)
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
	u := fmt.Sprintf("%s/do_transition", s.path)
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
	u := fmt.Sprintf("%s/edit_comment", s.path)
	v := new(issues.EditCommentResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GitlabSastExport - Return a list of vulnerabilities according to the Gitlab SAST JSON format.
// The JSON produced can be used in GitLab for generating the Vulnerability Report.Requires the 'Browse' or 'Scan' permission on the specified project.
// Since 10.2
// Changelog:
func (s *Issues) GitlabSastExport(ctx context.Context, r issues.GitlabSastExportRequest) (*issues.GitlabSastExportResponse, error) {
	u := fmt.Sprintf("%s/gitlab_sast_export", s.path)
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
	u := fmt.Sprintf("%s/reindex", s.path)

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
// Since 3.6
// Changelog:
//   10.6: Facet 'prioritizedRule' has been added
//   10.6: Param 'prioritizedRule' has been added
//   10.4: Added new param 'fixedInPullRequest'
//   10.4: Value 'wontfix' for 'transition' response field is deprecated, use 'accept' instead
//   10.4: Possible value 'accept' for 'transition' response field has been added
//   10.4: Param 'issueStatuses' has been added
//   10.4: Parameters 'resolutions' and 'statuses' are deprecated in favor of 'issueStatuses'.
//   10.4: Parameters 'severities' and 'types' are deprecated, use 'impactSeverities' and 'impactSoftwareQualities' instead.
//   10.4: Facet 'issueStatuses' has been added
//   10.4: Facets 'resolutions' and 'statuses' are deprecated in favor of 'issueStatuses'
//   10.4: Response fields 'severity' and 'type' are deprecated, use 'impacts' instead.
//   10.4: Response field 'issueStatus' added
//   10.4: Response fields 'status' and 'resolutions' are deprecated, in favor of 'issueStatus'
//   10.4: Possible value 'CONFIRMED' for 'issueStatus' field is deprecated.
//   10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//   10.2: Param 'impactSoftwareQualities' has been added
//   10.2: Param 'impactSeverities' has been added
//   10.2: Param 'cleanCodeAttributeCategories' has been added
//   10.2: Facet 'impactSoftwareQualities' has been added
//   10.2: Facet 'impactSeverities' has been added
//   10.2: Facet 'cleanCodeAttributeCategories' has been added
//   10.2: Parameter 'componentKeys' renamed to 'components'
//   10.1: Add the 'codeVariants' parameter, facet and response field
//   10.0: Parameter 'sansTop25' is deprecated
//   10.0: The value 'sansTop25' for the parameter 'facets' has been deprecated
//   10.0: Deprecated value 'ASSIGNEE' in parameter 's' is dropped
//   10.0: Parameter 'sinceLeakPeriod' is removed, please use 'inNewCodePeriod' instead
//   9.8: Add message formatting to issue and locations response
//   9.8: response fields 'total', 's', 'ps' have been deprecated, please use 'paging' object instead
//   9.7: Issues flows in the response may contain a description and a type
//   9.6: Response field 'fromHotspot' dropped.
//   9.6: Added facets 'pciDss-3.2' and 'pciDss-4.0
//   9.6: Added parameters 'pciDss-3.2' and 'pciDss-4.0
//   9.6: Response field 'ruleDescriptionContextKey' added
//   9.6: New possible value for 'additionalFields' parameter: 'ruleDescriptionContextKey'
//   9.6: Facet 'moduleUuids' is dropped.
//   9.4: Parameter 'sinceLeakPeriod' is deprecated, please use 'inNewCodePeriod' instead
//   9.2: Response field 'quickFixAvailable' added
//   9.1: Deprecated parameters 'authors', 'facetMode' and 'moduleUuids' were dropped
//   8.6: Parameter 'timeZone' added
//   8.5: Facet 'fileUuids' is dropped in favour of the new facet 'files'Note that they are not strictly identical, the latter returns the file paths.
//   8.5: Internal parameter 'fileUuids' has been dropped
//   8.4: parameters 'componentUuids', 'projectKeys' has been dropped.
//   8.2: 'REVIEWED', 'TO_REVIEW' status param values are no longer supported
//   8.2: Security hotspots are no longer returned as type 'SECURITY_HOTSPOT' is not supported anymore, use dedicated api/hotspots
//   8.2: response field 'fromHotspot' has been deprecated and is no more populated
//   8.2: Status 'IN_REVIEW' for Security Hotspots has been deprecated
//   7.8: added new Security Hotspots statuses : TO_REVIEW, IN_REVIEW and REVIEWED
//   7.8: Security hotspots are returned by default
//   7.7: Value 'authors' in parameter 'facets' is deprecated, please use 'author' instead
//   7.6: The use of module keys in parameter 'componentKeys' is deprecated
//   7.4: The facet 'projectUuids' is dropped in favour of the new facet 'projects'. Note that they are not strictly identical, the latter returns the project keys.
//   7.4: Parameter 'facetMode' does not accept anymore deprecated value 'debt'
//   7.3: response field 'fromHotspot' added to issues that are security hotspots
//   7.3: added facets 'sansTop25', 'owaspTop10' and 'cwe'
//   7.2: response field 'externalRuleEngine' added to issues that have been imported from an external rule engine
//   7.2: value 'ASSIGNEE' in parameter 's' is deprecated, it won't have any effect
//   6.5: parameters 'projects', 'projectUuids', 'moduleUuids', 'directories', 'fileUuids' are marked as internal
//   6.3: response field 'email' is renamed 'avatar'
//   5.5: response fields 'reporter' and 'actionPlan' are removed (drop of action plan and manual issue features)
//   5.5: parameters 'reporters', 'actionPlans' and 'planned' are dropped and therefore ignored (drop of action plan and manual issue features)
//   5.5: response field 'debt' is renamed 'effort'
func (s *Issues) Search(ctx context.Context, r issues.SearchRequest, p paging.Params) (*issues.SearchResponse, error) {
	u := fmt.Sprintf("%s/search", s.path)
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
	u := fmt.Sprintf("%s/set_severity", s.path)
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
	u := fmt.Sprintf("%s/set_tags", s.path)
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
	u := fmt.Sprintf("%s/set_type", s.path)
	v := new(issues.SetTypeResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Tags - List tags matching a given query
// Since 5.1
// Changelog:
//   9.4: Max page size increased to 500
//   7.4: Result doesn't include rules tags anymore
func (s *Issues) Tags(ctx context.Context, r issues.TagsRequest) (*issues.TagsResponse, error) {
	u := fmt.Sprintf("%s/tags", s.path)
	v := new(issues.TagsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
