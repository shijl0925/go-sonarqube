package measures

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ComponentRequest Return component with specified measures.<br>Requires one of the following permissions:<ul><li>'Browse' on the project of the specified component</li><li>'Execute Analysis' on the project of the specified component</li></ul>
type ComponentRequest struct {
	AdditionalFields string `url:"additionalFields,omitempty"` // Comma-separated list of additional fields that can be returned in the response.
	Branch           string `url:"branch,omitempty"`           // Since 6.6;Branch key. Not available in the community edition.
	Component        string `url:"component"`                  // Component key
	MetricKeys       string `url:"metricKeys"`                 // Comma-separated list of metric keys
	PullRequest      string `url:"pullRequest,omitempty"`      // Since 7.1;Pull request id. Not available in the community edition.
}

// ComponentResponse is the response for ComponentRequest
type ComponentResponse struct {
	Component struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Measures []struct {
			Metric string `json:"metric,omitempty"`
			Value  string `json:"value,omitempty"`
			Period struct {
				BestValue bool   `json:"bestValue,omitempty"`
				Value     string `json:"value,omitempty"`
			} `json:"period,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"component,omitempty"`
	Metrics []struct {
		Description           string `json:"description,omitempty"`
		Domain                string `json:"domain,omitempty"`
		Hidden                bool   `json:"hidden,omitempty"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter,omitempty"`
		Key                   string `json:"key,omitempty"`
		Name                  string `json:"name,omitempty"`
		Qualitative           bool   `json:"qualitative,omitempty"`
		Type                  string `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
	Period struct {
		Date      string `json:"date,omitempty"`
		Mode      string `json:"mode,omitempty"`
		Parameter string `json:"parameter,omitempty"`
	} `json:"period,omitempty"`
}

// ComponentTreeRequest Navigate through components based on the chosen strategy with specified measures.<br>Requires the following permission: 'Browse' on the specified project.<br>For applications, it also requires 'Browse' permission on its child projects. <br>When limiting search with the q parameter, directories are not returned.
type ComponentTreeRequest struct {
	AdditionalFields string `url:"additionalFields,omitempty"` // Comma-separated list of additional fields that can be returned in the response.
	Asc              string `url:"asc,omitempty"`              // Ascending sort
	Branch           string `url:"branch,omitempty"`           // Since 6.6;Branch key. Not available in the community edition.
	Component        string `url:"component"`                  // Component key. The search is based on this component.
	MetricKeys       string `url:"metricKeys"`                 // Comma-separated list of metric keys. Types DISTRIB are not allowed. For type DATA only maintainability_issues, security_issues, new_reliability_issues, new_security_issues, new_maintainability_issues, reliability_issues metrics are supported
	MetricPeriodSort string `url:"metricPeriodSort,omitempty"` // Since 5.5;Sort measures by leak period or not ?. The 's' parameter must contain the 'metricPeriod' value.
	MetricSort       string `url:"metricSort,omitempty"`       // Metric key to sort by. The 's' parameter must contain the 'metric' or 'metricPeriod' value. It must be part of the 'metricKeys' parameter
	MetricSortFilter string `url:"metricSortFilter,omitempty"` // Filter components. Sort must be on a metric. Possible values are: <ul><li>all: return all components</li><li>withMeasuresOnly: filter out components that do not have a measure on the sorted metric</li></ul>
	PullRequest      string `url:"pullRequest,omitempty"`      // Since 7.1;Pull request id. Not available in the community edition.
	Q                string `url:"q,omitempty"`                // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>
	Qualifiers       string `url:"qualifiers,omitempty"`       // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>APP - Applications</li><li>VW - Portfolios</li><li>SVW - Portfolios</li><li>UTS - Test Files</li><li>FIL - Files</li><li>DIR - Directories</li><li>TRK - Projects</li></ul>
	S                string `url:"s,omitempty"`                // Comma-separated list of sort fields
	Strategy         string `url:"strategy,omitempty"`         // Strategy to search for base component descendants:<ul><li>children: return the children components of the base component. Grandchildren components are not returned</li><li>all: return all the descendants components of the base component. Grandchildren are returned.</li><li>leaves: return all the descendant components (files, in general) which don't have other children. They are the leaves of the component tree.</li></ul>
}

// ComponentTreeResponse is the response for ComponentTreeRequest
type ComponentTreeResponse struct {
	BaseComponent struct {
		Key      string `json:"key,omitempty"`
		Measures []struct {
			Metric string `json:"metric,omitempty"`
			Period struct {
				Value string `json:"value,omitempty"`
			} `json:"period,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"baseComponent,omitempty"`
	Components []struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Measures []struct {
			Metric string `json:"metric,omitempty"`
			Period struct {
				Value string `json:"value,omitempty"`
			} `json:"period,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Metrics []struct {
		Description           string `json:"description,omitempty"`
		Domain                string `json:"domain,omitempty"`
		Hidden                bool   `json:"hidden,omitempty"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter,omitempty"`
		Key                   string `json:"key,omitempty"`
		Name                  string `json:"name,omitempty"`
		Qualitative           bool   `json:"qualitative,omitempty"`
		Type                  string `json:"type,omitempty"`
		BestValue             string `json:"bestValue,omitempty"`
	} `json:"metrics,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
	Period struct {
		Date      string `json:"date,omitempty"`
		Mode      string `json:"mode,omitempty"`
		Parameter string `json:"parameter,omitempty"`
	} `json:"period,omitempty"`
}

// GetPaging extracts the paging from ComponentTreeResponse
func (r *ComponentTreeResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// ComponentTreeResponseAll is the collection for ComponentTreeRequest
type ComponentTreeResponseAll struct {
	BaseComponent struct {
		Key      string `json:"key,omitempty"`
		Measures []struct {
			Metric string `json:"metric,omitempty"`
			Period struct {
				Value string `json:"value,omitempty"`
			} `json:"period,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"baseComponent,omitempty"`
	Components []struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Measures []struct {
			Metric string `json:"metric,omitempty"`
			Period struct {
				Value string `json:"value,omitempty"`
			} `json:"period,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Metrics []struct {
		Description           string `json:"description,omitempty"`
		Domain                string `json:"domain,omitempty"`
		Hidden                bool   `json:"hidden,omitempty"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter,omitempty"`
		Key                   string `json:"key,omitempty"`
		Name                  string `json:"name,omitempty"`
		Qualitative           bool   `json:"qualitative,omitempty"`
		Type                  string `json:"type,omitempty"`
		BestValue             string `json:"bestValue,omitempty"`
	} `json:"metrics,omitempty"`
	Period struct {
		Date      string `json:"date,omitempty"`
		Mode      string `json:"mode,omitempty"`
		Parameter string `json:"parameter,omitempty"`
	} `json:"period,omitempty"`
}

// SearchHistoryRequest Search measures history of a component.<br>Measures are ordered chronologically.<br>Pagination applies to the number of measures for each metric.<br>Requires the following permission: 'Browse' on the specified component. <br>For applications, it also requires 'Browse' permission on its child projects.
type SearchHistoryRequest struct {
	Branch      string `url:"branch,omitempty"`      // Since 6.6;Branch key. Not available in the community edition.
	Component   string `url:"component"`             // Component key
	From        string `url:"from,omitempty"`        // Filter measures created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
	Metrics     string `url:"metrics"`               // Comma-separated list of metric keys
	PullRequest string `url:"pullRequest,omitempty"` // Since 7.1;Pull request id. Not available in the community edition.
	To          string `url:"to,omitempty"`          // Filter measures created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
}

// SearchHistoryResponse is the response for SearchHistoryRequest
type SearchHistoryResponse struct {
	Measures []struct {
		History []struct {
			Date  string `json:"date,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"history,omitempty"`
		Metric string `json:"metric,omitempty"`
	} `json:"measures,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchHistoryResponse
func (r *SearchHistoryResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchHistoryResponseAll is the collection for SearchHistoryRequest
type SearchHistoryResponseAll struct {
	Measures []struct {
		History []struct {
			Date  string `json:"date,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"history,omitempty"`
		Metric string `json:"metric,omitempty"`
	} `json:"measures,omitempty"`
}
