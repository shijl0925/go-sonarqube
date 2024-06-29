package components

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// SearchRequest Search for components
type SearchRequest struct {
	Q          string `form:"q,omitempty"`          // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul><br>The value length of the param must be between 2 and 15 (inclusive) characters. In case longer value is provided it will be truncated.
	Qualifiers string `form:"qualifiers,omitempty"` // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>APP - Applications</li><li>VW - Portfolios</li><li>SVW - Portfolios</li><li>TRK - Projects</li></ul>
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Components []struct {
		Key       string `json:"key,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Components []struct {
		Key       string `json:"key,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
}

// ShowRequest Returns a component (file, directory, project, portfolio…) and its ancestors. The ancestors are ordered from the parent to the root project. Requires the following permission: 'Browse' on the project of the specified component.
type ShowRequest struct {
	Branch      string `form:"branch,omitempty"`      // Branch key. Not available in the community edition.
	Component   string `form:"component,omitempty"`   // Component key
	PullRequest string `form:"pullRequest,omitempty"` // Pull request id. Not available in the community edition.
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Ancestors []struct {
		AnalysisDate string   `json:"analysisDate,omitempty"`
		Key          string   `json:"key,omitempty"`
		Name         string   `json:"name,omitempty"`
		Path         string   `json:"path,omitempty"`
		Qualifier    string   `json:"qualifier,omitempty"`
		Version      string   `json:"version,omitempty"`
		Description  string   `json:"description,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		Visibility   string   `json:"visibility,omitempty"`
	} `json:"ancestors,omitempty"`
	Component struct {
		AnalysisDate   string `json:"analysisDate,omitempty"`
		Key            string `json:"key,omitempty"`
		Language       string `json:"language,omitempty"`
		LeakPeriodDate string `json:"leakPeriodDate,omitempty"`
		Name           string `json:"name,omitempty"`
		Path           string `json:"path,omitempty"`
		Qualifier      string `json:"qualifier,omitempty"`
		Version        string `json:"version,omitempty"`
	} `json:"component,omitempty"`
}

// TreeRequest Navigate through components based on the chosen strategy.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.
type TreeRequest struct {
	Asc         string `form:"asc,omitempty"`         // Ascending sort
	Branch      string `form:"branch,omitempty"`      // Branch key. Not available in the community edition.
	Component   string `form:"component,omitempty"`   // Base component key. The search is based on this component.
	PullRequest string `form:"pullRequest,omitempty"` // Pull request id. Not available in the community edition.
	Q           string `form:"q,omitempty"`           // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>
	Qualifiers  string `form:"qualifiers,omitempty"`  // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>APP - Applications</li><li>VW - Portfolios</li><li>SVW - Portfolios</li><li>UTS - Test Files</li><li>FIL - Files</li><li>DIR - Directories</li><li>TRK - Projects</li></ul>
	S           string `form:"s,omitempty"`           // Comma-separated list of sort fields
	Strategy    string `form:"strategy,omitempty"`    // Strategy to search for base component descendants:<ul><li>children: return the children components of the base component. Grandchildren components are not returned</li><li>all: return all the descendants components of the base component. Grandchildren are returned.</li><li>leaves: return all the descendant components (files, in general) which don't have other children. They are the leaves of the component tree.</li></ul>
}

// TreeResponse is the response for TreeRequest
type TreeResponse struct {
	BaseComponent struct {
		Description string   `json:"description,omitempty"`
		Key         string   `json:"key,omitempty"`
		Qualifier   string   `json:"qualifier,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Visibility  string   `json:"visibility,omitempty"`
	} `json:"baseComponent,omitempty"`
	Components []struct {
		Key       string `json:"key,omitempty"`
		Language  string `json:"language,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from TreeResponse
func (r *TreeResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// TreeResponseAll is the collection for TreeRequest
type TreeResponseAll struct {
	BaseComponent struct {
		Description string   `json:"description,omitempty"`
		Key         string   `json:"key,omitempty"`
		Qualifier   string   `json:"qualifier,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Visibility  string   `json:"visibility,omitempty"`
	} `json:"baseComponent,omitempty"`
	Components []struct {
		Key       string `json:"key,omitempty"`
		Language  string `json:"language,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
}
