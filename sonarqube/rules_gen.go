package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/rules"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Rules service

// Create - Create a custom rule.
// Requires the 'Administer Quality Profiles' permission
// Since 4.4
// Changelog:
//
//	10.8: The parameters 'type' and 'severity' are not deprecated anymore.
//	10.4: Add 'impacts' and 'cleanCodeAttribute' parameters to the request
//	10.4: Parameters 'type' and 'severity' are deprecated. Use 'impacts' instead.
//	10.4: Parameter 'preventReactivation' is deprecated. Use api/rules/update endpoint instead.
//	10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//	10.2: Fields 'type' and 'severity' are deprecated in the response. Use 'impacts' instead.
//	10.0: Drop deprecated keys: 'custom_key', 'template_key', 'markdown_description', 'prevent_reactivation'
//	5.5: Creating manual rule is not more possible
func (s *Rules) Create(ctx context.Context, r rules.CreateRequest) (*rules.CreateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)
	v := new(rules.CreateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Delete - Delete custom rule.
// Requires the 'Administer Quality Profiles' permission
// Since 4.4
func (s *Rules) Delete(ctx context.Context, r rules.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Repositories - List available rule repositories
// Since 4.5
func (s *Rules) Repositories(ctx context.Context, r rules.RepositoriesRequest) (*rules.RepositoriesResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/repositories", s.path)
	v := new(rules.RepositoriesResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Search - Search for a collection of relevant rules matching a specified query.
//
// Since 4.4
// Changelog:
//
//	10.8: Possible values 'INFO' and 'BLOCKER' for response field 'impactSeverities' of 'facets' have been added
//	10.8: Possible values 'INFO' and 'BLOCKER' for response field 'severity' of 'impacts' have been added
//	10.8: Parameter 'severities' now supports values: 'INFO','BLOCKER'
//	10.8: The field 'impacts' has been added to the response
//	10.8: The parameters 'severities','types and 'active_severities' are not deprecated anymore.
//	10.8: The values 'severity' and 'types' for the 'facets' parameter are not deprecated anymore.
//	10.8: The fields 'type' and 'severity' in the response are not deprecated anymore.
//	10.8: The value 'severity' for the 'f' parameter is not deprecated anymore.
//	10.6: Parameter 'prioritizedRule has been added
//	10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//	10.2: The fields 'type' and 'severity' are deprecated in the response. Use 'impacts' instead.
//	10.2: The field 'cleanCodeAttribute' has been added to the 'f' parameter.
//	10.2: The value 'severity' for the 'f' parameter has been deprecated.
//	10.2: The values 'cleanCodeAttributeCategories', 'impactSoftwareQualities' and 'impactSeverities' have been added to the 'facets' parameter.
//	10.2: The values 'severity' and 'types' for the 'facets' parameter have been deprecated. Use 'impactSeverities' and 'impactSoftwareQualities' instead.
//	10.2: Parameters 'severities', 'types', and 'active_severities' are now deprecated. Use 'impactSoftwareQualities' and 'impactSeverities' instead.
//	10.0: The deprecated field 'effortToFixDescription' has been removed, use 'gapDescription' instead.
//	10.0: The deprecated field 'debtRemFnCoeff' has been removed, use 'remFnGapMultiplier' instead.
//	10.0: The deprecated field 'defaultDebtRemFnCoeff' has been removed, use 'defaultRemFnGapMultiplier' instead.
//	10.0: The deprecated field 'debtRemFnOffset' has been removed, use 'remFnBaseEffort' instead.
//	10.0: The deprecated field 'defaultDebtRemFnOffset' has been removed, use 'defaultRemFnBaseEffort' instead.
//	10.0: The deprecated field 'debtOverloaded' has been removed, use 'remFnOverloaded' instead.
//	10.0: The field 'defaultDebtRemFnType' has been deprecated, use 'defaultRemFnType' instead
//	10.0: The field 'debtRemFnType' has been deprecated, use 'remFnType' instead
//	10.0: The value 'debtRemFn' for the 'f' parameter has been deprecated, use 'remFn' instead
//	10.0: The value 'defaultDebtRemFn' for the 'f' parameter has been deprecated, use 'defaultRemFn' instead
//	10.0: The value 'sansTop25' for the parameter 'facets' has been deprecated
//	10.0: Parameter 'sansTop25' is deprecated
//	9.8: response fields 'total', 's', 'ps' have been deprecated, please use 'paging' object instead
//	9.8: The field 'paging' has been added to the response
//	9.6: 'descriptionSections' can optionally embed a context field
//	9.6: The field 'educationPrinciples' has been added to the 'f' parameter
//	9.5: The field 'htmlDesc' has been deprecated, use 'descriptionSections' instead
//	9.5: The field 'descriptionSections' has been added to the payload
//	9.5: The field 'descriptionSections' has been added to the 'f' parameter
//	7.5: The field 'updatedAt' has been added to the 'f' parameter
//	7.2: The field 'isExternal' has been added to the response
//	7.2: The field 'includeExternal' has been added to the 'f' parameter
//	7.1: The field 'scope' has been added to the response
//	7.1: The field 'scope' has been added to the 'f' parameter
//	5.5: The field 'effortToFixDescription' has been deprecated, use 'gapDescription' instead
//	5.5: The field 'debtRemFnCoeff' has been deprecated, use 'remFnGapMultiplier' instead
//	5.5: The field 'defaultDebtRemFnCoeff' has been deprecated, use 'defaultRemFnGapMultiplier' instead
//	5.5: The field 'debtRemFnOffset' has been deprecated, use 'remFnBaseEffort' instead
//	5.5: The field 'defaultDebtRemFnOffset' has been deprecated, use 'defaultRemFnBaseEffort' instead
//	5.5: The field 'debtOverloaded' has been deprecated, use 'remFnOverloaded' instead
func (s *Rules) Search(ctx context.Context, r rules.SearchRequest, p paging.Params) (*rules.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(rules.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Rules) SearchAll(ctx context.Context, r rules.SearchRequest) (*rules.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &rules.SearchResponseAll{}
	for {
		res, _, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to rules.Search: %+v", err)
		}
		response.Facets = append(response.Facets, res.Facets...)
		response.Rules = append(response.Rules, res.Rules...)
		if res.GetPaging().End() {
			break
		}
		p.P++
	}
	return response, nil
}

// Show - Get detailed information about a rule
//
// Since 4.2
// Changelog:
//
//	10.8: Possible values 'INFO' and 'BLOCKER' for response field 'severity' of 'impacts' have been added.
//	10.8: The field 'severity' and 'type' in the response  are not deprecated anymore.
//	10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
//	10.2: The field 'severity' and 'type' in the response have been deprecated, use 'impacts' instead.
//	10.0: The deprecated field 'effortToFixDescription' has been removed, use 'gapDescription' instead.
//	10.0: The deprecated field 'debtRemFnCoeff' has been removed, use 'remFnGapMultiplier' instead.
//	10.0: The deprecated field 'defaultDebtRemFnCoeff' has been removed, use 'defaultRemFnGapMultiplier' instead.
//	10.0: The deprecated field 'debtRemFnOffset' has been removed, use 'remFnBaseEffort' instead.
//	10.0: The deprecated field 'defaultDebtRemFnOffset' has been removed, use 'defaultRemFnBaseEffort' instead.
//	10.0: The deprecated field 'debtOverloaded' has been removed, use 'remFnOverloaded' instead.
//	10.0: The field 'defaultDebtRemFnType' has been deprecated, use 'defaultRemFnType' instead
//	10.0: The field 'debtRemFnType' has been deprecated, use 'remFnType' instead
//	9.6: 'descriptionSections' can optionally embed a context field.
//	9.6: 'educationPrinciples' has been added.
//	9.5: The field 'htmlDesc' in the response has been deprecated, it becomes 'descriptionSections'.
//	9.5: The field 'descriptionSections' has been added to the payload.
//	7.1: The field 'scope' has been added.
//	5.5: The field 'effortToFixDescription' in the response has been deprecated, it becomes 'gapDescription'.
//	5.5: The field 'debtRemFnCoeff' in the response has been deprecated, it becomes 'remFnGapMultiplier'.
//	5.5: The field 'defaultDebtRemFnCoeff' in the response has been deprecated, it becomes 'defaultRemFnGapMultiplier'.
//	5.5: The field 'debtRemFnOffset' in the response has been deprecated, it becomes 'remFnBaseEffort'.
//	5.5: The field 'defaultDebtRemFnOffset' in the response has been deprecated, it becomes 'defaultRemFnBaseEffort'.
//	5.5: The field 'debtOverloaded' in the response has been deprecated, it becomes 'remFnOverloaded'.
func (s *Rules) Show(ctx context.Context, r rules.ShowRequest) (*rules.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(rules.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Tags - List rule tags
// Since 4.4
// Changelog:
//
//	9.4: Max page size increased to 500
func (s *Rules) Tags(ctx context.Context, r rules.TagsRequest) (*rules.TagsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/tags", s.path)
	v := new(rules.TagsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Update - Update an existing rule.
// Requires the 'Administer Quality Profiles' permission
// Since 4.4
// Changelog:
//
//	10.8: Parameter impacts was added.
//	10.8: The parameter 'severity' is not deprecated anymore.
//	10.8: The field 'severity' and 'type' in the response are not deprecated anymore.
//	10.4: The parameter 'severity' is deprecated.
//	10.4: Updating a removed rule is now possible.
//	10.2: The field 'severity' and 'type' in the response have been deprecated, use 'impacts' instead.
//	10.2: Add 'impacts', 'cleanCodeAttribute', 'cleanCodeAttributeCategory' fields to the response
func (s *Rules) Update(ctx context.Context, r rules.UpdateRequest) (*rules.UpdateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/update", s.path)
	v := new(rules.UpdateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
