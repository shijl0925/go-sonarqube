package favorites

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddRequest Add a component (project, portfolio, etc.) as favorite for the authenticated user.<br>Only 100 components by qualifier can be added as favorite.<br>Requires authentication and the following permission: 'Browse' on the component.
type AddRequest struct {
	Component string `form:"component,omitempty"` // Component key. Only components with qualifiers TRK, VW, SVW, APP are supported
}

// RemoveRequest Remove a component (project, portfolio, application etc.) as favorite for the authenticated user.<br>Requires authentication.
type RemoveRequest struct {
	Component string `form:"component,omitempty"` // Component key
}

// SearchRequest Search for the authenticated user favorites.<br>Requires authentication.
type SearchRequest struct{}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Favorites []struct {
		Key       string `json:"key,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"favorites,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Favorites []struct {
		Key       string `json:"key,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"favorites,omitempty"`
}
