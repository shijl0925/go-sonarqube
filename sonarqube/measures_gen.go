package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/measures"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Measures service

// Component - Return component with specified measures.
// Requires the following permission: 'Browse' on the project of specified component.
// Since 5.4
// Changelog:
//   10.5: The metrics 'new_blocker_violations', 'new_critical_violations', 'new_major_violations', 'new_minor_violations', 'new_info_violations', 'blocker_violations', 'critical_violations', 'major_violations', 'minor_violations', 'info_violations' are now deprecated without exact replacement. Use 'maintainability_issues', 'reliability_issues' and 'security_issues' instead.
//   10.5: Added new accepted values for the 'metricKeys' param: 'new_maintainability_issues', 'new_reliability_issues', 'new_security_issues'
//   10.4: The metrics 'bugs', 'new_bugs', 'vulnerabilities', 'new_vulnerabilities', 'code_smells', 'new_code_smells', 'high_impact_accepted_issues' are now deprecated without exact replacement. Use 'maintainability_issues', 'reliability_issues' and 'security_issues' instead.
//   10.4: Added new accepted values for the 'metricKeys' param: 'maintainability_issues', 'reliability_issues', 'security_issues'
//   10.4: The metrics 'open_issues', 'reopened_issues' and 'confirmed_issues' are now deprecated in the response. Consume 'violations' instead.
//   10.4: The use of 'open_issues', 'reopened_issues' and 'confirmed_issues' values in 'metricKeys' param are now deprecated. Use 'violations' instead.
//   10.4: The metric 'wont_fix_issues' is now deprecated in the response. Consume 'accepted_issues' instead.
//   10.4: The use of 'wont_fix_issues' value in 'metricKeys' param is now deprecated. Use 'accepted_issues' instead.
//   10.4: Added new accepted value for the 'metricKeys' param: 'accepted_issues'.
//   10.1: The use of module keys in parameter 'component' is removed
//   10.0: The use of the following metrics in 'metricKeys' parameter is not deprecated anymore: 'releasability_effort', 'security_rating_effort', 'reliability_rating_effort', 'security_review_rating_effort', 'maintainability_rating_effort', 'last_change_on_maintainability_rating', 'last_change_on_releasability_rating', 'last_change_on_reliability_rating', 'last_change_on_security_rating', 'last_change_on_security_review_rating'
//   10.0: the response field periods under measures field is removed.
//   10.0: the option `periods` of 'additionalFields' request field is removed.
//   9.3: When the new code period is set to 'reference branch', the response field 'date' under the 'period' field has been removed
//   9.3: The use of the following metrics in 'metricKeys' parameter is deprecated: 'releasability_effort', 'security_rating_effort', 'reliability_rating_effort', 'security_review_rating_effort', 'maintainability_rating_effort', 'last_change_on_maintainability_rating', 'last_change_on_releasability_rating', 'last_change_on_reliability_rating', 'last_change_on_security_rating', 'last_change_on_security_review_rating'
//   8.8: deprecated response field 'id' has been removed
//   8.8: deprecated response field 'refId' has been removed.
//   8.1: the response field periods under measures field is deprecated. Use period instead.
//   8.1: the response field periods is deprecated. Use period instead.
//   7.6: The use of module keys in parameter 'component' is deprecated
//   6.6: the response field 'id' is deprecated. Use 'key' instead.
//   6.6: the response field 'refId' is deprecated. Use 'refKey' instead.
func (s *Measures) Component(ctx context.Context, r measures.ComponentRequest) (*measures.ComponentResponse, error) {
	u := fmt.Sprintf("%s/component", s.path)
	v := new(measures.ComponentResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ComponentTree - Navigate through components based on the chosen strategy with specified measures.
// Requires the following permission: 'Browse' on the specified project.
// For applications, it also requires 'Browse' permission on its child projects.
// When limiting search with the q parameter, directories are not returned.
// Since 5.4
// Changelog:
//   10.5: Added new accepted values for the 'metricKeys' param: 'new_maintainability_issues', 'new_reliability_issues', 'new_security_issues'
//   10.5: The metrics 'new_blocker_violations', 'new_critical_violations', 'new_major_violations', 'new_minor_violations', 'new_info_violations', 'blocker_violations', 'critical_violations', 'major_violations', 'minor_violations', 'info_violations' are now deprecated without exact replacement. Use 'maintainability_issues', 'reliability_issues' and 'security_issues' instead.
//   10.5: Added new accepted values for the 'metricKeys' param: 'maintainability_issues', 'reliability_issues', 'security_issues'
//   10.4: The metrics 'bugs', 'new_bugs', 'vulnerabilities', 'new_vulnerabilities', 'code_smells', 'new_code_smells', 'high_impact_accepted_issues' are now deprecated without exact replacement. Use 'maintainability_issues', 'reliability_issues' and 'security_issues' instead.
//   10.4: The metrics 'open_issues', 'reopened_issues' and 'confirmed_issues' are now deprecated in the response. Consume 'violations' instead.
//   10.4: The use of 'open_issues', 'reopened_issues' and 'confirmed_issues' values in 'metricKeys' param are now deprecated. Use 'violations' instead.
//   10.4: The metric 'wont_fix_issues' is now deprecated in the response. Consume 'accepted_issues' instead.
//   10.4: The use of 'wont_fix_issues' value in 'metricKeys' and 'metricSort' params is now deprecated. Use 'accepted_issues' instead.
//   10.4: Added new accepted value for the 'metricKeys' and 'metricSort' param: 'accepted_issues'.
//   10.1: The use of 'BRC' as value for parameter 'qualifiers' is removed
//   10.0: The use of the following metrics in 'metricKeys' parameter is not deprecated anymore: 'releasability_effort', 'security_rating_effort', 'reliability_rating_effort', 'security_review_rating_effort', 'maintainability_rating_effort', 'last_change_on_maintainability_rating', 'last_change_on_releasability_rating', 'last_change_on_reliability_rating', 'last_change_on_security_rating', 'last_change_on_security_review_rating'
//   10.0: the response field periods under measures field is removed.
//   10.0: the option `periods` of 'additionalFields' request field is removed.
//   9.3: The use of the following metrics in 'metricKeys' parameter is deprecated: 'releasability_effort', 'security_rating_effort', 'reliability_rating_effort', 'security_review_rating_effort', 'maintainability_rating_effort', 'last_change_on_maintainability_rating', 'last_change_on_releasability_rating', 'last_change_on_reliability_rating', 'last_change_on_security_rating', 'last_change_on_security_review_rating'
//   8.8: parameter 'component' is now required
//   8.8: deprecated parameter 'baseComponentId' has been removed
//   8.8: deprecated parameter 'baseComponentKey' has been removed.
//   8.8: deprecated response field 'id' has been removed
//   8.8: deprecated response field 'refId' has been removed.
//   8.1: the response field periods under measures field is deprecated. Use period instead.
//   8.1: the response field periods is deprecated. Use period instead.
//   7.6: The use of module keys in parameter 'component' is deprecated
//   7.2: field 'bestValue' is added to the response
//   6.6: the response field 'id' is deprecated. Use 'key' instead.
//   6.6: the response field 'refId' is deprecated. Use 'refKey' instead.
//   6.3: Number of metric keys is limited to 15
func (s *Measures) ComponentTree(ctx context.Context, r measures.ComponentTreeRequest, p paging.Params) (*measures.ComponentTreeResponse, error) {
	u := fmt.Sprintf("%s/component_tree", s.path)
	v := new(measures.ComponentTreeResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Measures) ComponentTreeAll(ctx context.Context, r measures.ComponentTreeRequest) (*measures.ComponentTreeResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &measures.ComponentTreeResponseAll{}
	for {
		res, err := s.ComponentTree(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to measures.ComponentTree: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Metrics = append(response.Metrics, res.Metrics...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// SearchHistory - Search measures history of a component.
// Measures are ordered chronologically.
// Pagination applies to the number of measures for each metric.
// Requires the following permission: 'Browse' on the specified component.
// For applications, it also requires 'Browse' permission on its child projects.
// Since 6.3
// Changelog:
//   10.5: The metrics 'new_blocker_violations', 'new_critical_violations', 'new_major_violations', 'new_minor_violations', 'new_info_violations', 'blocker_violations', 'critical_violations', 'major_violations', 'minor_violations', 'info_violations' are now deprecated without exact replacement. Use 'maintainability_issues', 'reliability_issues' and 'security_issues' instead.
//   10.5: Added new accepted values for the 'metricKeys' param: 'new_maintainability_issues', 'new_reliability_issues', 'new_security_issues'
//   10.4: The metrics 'bugs', 'new_bugs', 'vulnerabilities', 'new_vulnerabilities', 'code_smells', 'new_code_smells', 'high_impact_accepted_issues' are now deprecated without exact replacement. Use 'maintainability_issues', 'reliability_issues' and 'security_issues' instead.
//   10.4: The metrics 'open_issues', 'reopened_issues' and 'confirmed_issues' are now deprecated in the response. Consume 'violations' instead.
//   10.4: The use of 'open_issues', 'reopened_issues' and 'confirmed_issues' values in 'metricKeys' param are now deprecated. Use 'violations' instead.
//   10.4: The metric 'wont_fix_issues' is now deprecated in the response. Consume 'accepted_issues' instead.
//   10.4: The use of 'wont_fix_issues' value in 'metricKeys' param is now deprecated. Use 'accepted_issues' instead.
//   10.4: Added new accepted value for the 'metricKeys' param: 'accepted_issues'.
//   10.0: The use of the following metrics in 'metricKeys' parameter is not deprecated anymore: 'releasability_effort', 'security_rating_effort', 'reliability_rating_effort', 'security_review_rating_effort', 'maintainability_rating_effort', 'last_change_on_maintainability_rating', 'last_change_on_releasability_rating', 'last_change_on_reliability_rating', 'last_change_on_security_rating', 'last_change_on_security_review_rating'
//   9.3: The use of the following metrics in 'metrics' parameter is deprecated: 'releasability_effort', 'security_rating_effort', 'reliability_rating_effort', 'security_review_rating_effort', 'maintainability_rating_effort', 'last_change_on_maintainability_rating', 'last_change_on_releasability_rating', 'last_change_on_reliability_rating', 'last_change_on_security_rating', 'last_change_on_security_review_rating'
//   7.6: The use of module keys in parameter 'component' is deprecated
func (s *Measures) SearchHistory(ctx context.Context, r measures.SearchHistoryRequest, p paging.Params) (*measures.SearchHistoryResponse, error) {
	u := fmt.Sprintf("%s/search_history", s.path)
	v := new(measures.SearchHistoryResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Measures) SearchHistoryAll(ctx context.Context, r measures.SearchHistoryRequest) (*measures.SearchHistoryResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &measures.SearchHistoryResponseAll{}
	for {
		res, err := s.SearchHistory(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to measures.SearchHistory: %+v", err)
		}
		response.Measures = append(response.Measures, res.Measures...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
