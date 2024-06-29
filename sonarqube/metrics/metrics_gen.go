package metrics

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// SearchRequest Search for metrics
type SearchRequest struct{}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Metrics []struct {
		Custom      bool    `json:"custom,omitempty"`
		Description string  `json:"description,omitempty"`
		Direction   float64 `json:"direction,omitempty"`
		Domain      string  `json:"domain,omitempty"`
		Hidden      bool    `json:"hidden,omitempty"`
		Id          string  `json:"id,omitempty"`
		Key         string  `json:"key,omitempty"`
		Name        string  `json:"name,omitempty"`
		Qualitative bool    `json:"qualitative,omitempty"`
		Type        string  `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Metrics []struct {
		Custom      bool    `json:"custom,omitempty"`
		Description string  `json:"description,omitempty"`
		Direction   float64 `json:"direction,omitempty"`
		Domain      string  `json:"domain,omitempty"`
		Hidden      bool    `json:"hidden,omitempty"`
		Id          string  `json:"id,omitempty"`
		Key         string  `json:"key,omitempty"`
		Name        string  `json:"name,omitempty"`
		Qualitative bool    `json:"qualitative,omitempty"`
		Type        string  `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
}

// TypesRequest List all available metric types.
type TypesRequest struct{}

// TypesResponse is the response for TypesRequest
type TypesResponse struct {
	Types []string `json:"types,omitempty"`
}
