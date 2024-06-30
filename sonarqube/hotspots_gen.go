package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/hotspots"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Hotspots service

// ChangeStatus - Change the status of a Security Hotpot.
// Requires the 'Administer Security Hotspot' permission.
// Since 8.1
// Changelog:
//   10.1: Endpoint visibility change from internal to public
func (s *Hotspots) ChangeStatus(ctx context.Context, r hotspots.ChangeStatusRequest) error {
	u := fmt.Sprintf("%s/hotspots/change_status", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Search - Search for Security Hotpots.
// Requires the 'Browse' permission on the specified project(s).
// For applications, it also requires 'Browse' permission on its child projects.
// When issue indexing is in progress returns 503 service unavailable HTTP code.
// Since 8.1
// Changelog:
//   10.2: Parameter 'projectKey' renamed to 'project'
//   10.0: Parameter 'sansTop25' is deprecated
//   9.8: Endpoint visibility change from internal to public
//   9.8: Add message formatting to issue and locations response
//   9.7: Hotspot flows in the response may contain a description and a type
//   9.7: Hotspot in the response contain the corresponding ruleKey
//   9.6: Added parameters 'pciDss-3.2' and 'pciDss-4.0
func (s *Hotspots) Search(ctx context.Context, r hotspots.SearchRequest, p paging.Params) (*hotspots.SearchResponse, error) {
	u := fmt.Sprintf("%s/hotspots/search", API)
	v := new(hotspots.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Hotspots) SearchAll(ctx context.Context, r hotspots.SearchRequest) (*hotspots.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &hotspots.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to hotspots.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Hotspots = append(response.Hotspots, res.Hotspots...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// Show - Provides the details of a Security Hotspot.
// Since 8.1
// Changelog:
//   10.1: Add the 'codeVariants' response field
//   9.8: Add message formatting to issue and locations response
//   9.7: Hotspot flows in the response may contain a description and a type
//   9.5: The fields rule.riskDescription, rule.fixRecommendations, rule.vulnerabilityDescription of the response are deprecated. /api/rules/show endpoint should be used to fetch rule descriptions.
func (s *Hotspots) Show(ctx context.Context, r hotspots.ShowRequest) (*hotspots.ShowResponse, error) {
	u := fmt.Sprintf("%s/hotspots/show", API)
	v := new(hotspots.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
