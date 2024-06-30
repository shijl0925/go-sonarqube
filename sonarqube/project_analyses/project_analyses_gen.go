package project_analyses

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// CreateEventRequest Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
type CreateEventRequest struct {
	Analysis string `json:"analysis"`           // Analysis key
	Category string `json:"category,omitempty"` // Category
	Name     string `json:"name"`               // Name
}

// CreateEventResponse is the response for CreateEventRequest
type CreateEventResponse struct {
	Event struct {
		Analysis string `json:"analysis,omitempty"`
		Category string `json:"category,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"event,omitempty"`
}

// DeleteRequest Delete a project analysis.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the project of the specified analysis</li></ul>
type DeleteRequest struct {
	Analysis string `json:"analysis"` // Analysis key
}

// DeleteEventRequest Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
type DeleteEventRequest struct {
	Event string `json:"event"` // Event key
}

// SearchRequest Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project. <br>For applications, it also requires 'Browse' permission on its child projects.
type SearchRequest struct {
	Category string `url:"category,omitempty"` // Event category. Filter analyses that have at least one event of the category specified.
	From     string `url:"from,omitempty"`     // Since 6.5;Filter analyses created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
	Project  string `url:"project"`            // Project key
	To       string `url:"to,omitempty"`       // Since 6.5;Filter analyses created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Analyses []struct {
		BuildString string `json:"buildString,omitempty"`
		Date        string `json:"date,omitempty"`
		Events      []struct {
			Category       string `json:"category,omitempty"`
			Key            string `json:"key,omitempty"`
			Name           string `json:"name,omitempty"`
			QualityProfile struct {
				Key         string `json:"key,omitempty"`
				LanguageKey string `json:"languageKey,omitempty"`
				Name        string `json:"name,omitempty"`
			} `json:"qualityProfile,omitempty"`
		} `json:"events,omitempty"`
		Key                         string `json:"key,omitempty"`
		ManualNewCodePeriodBaseline bool   `json:"manualNewCodePeriodBaseline,omitempty"`
		ProjectVersion              string `json:"projectVersion,omitempty"`
		Revision                    string `json:"revision,omitempty"`
		DetectedCI                  string `json:"detectedCI,omitempty"`
	} `json:"analyses,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Analyses []struct {
		BuildString string `json:"buildString,omitempty"`
		Date        string `json:"date,omitempty"`
		Events      []struct {
			Category       string `json:"category,omitempty"`
			Key            string `json:"key,omitempty"`
			Name           string `json:"name,omitempty"`
			QualityProfile struct {
				Key         string `json:"key,omitempty"`
				LanguageKey string `json:"languageKey,omitempty"`
				Name        string `json:"name,omitempty"`
			} `json:"qualityProfile,omitempty"`
		} `json:"events,omitempty"`
		Key                         string `json:"key,omitempty"`
		ManualNewCodePeriodBaseline bool   `json:"manualNewCodePeriodBaseline,omitempty"`
		ProjectVersion              string `json:"projectVersion,omitempty"`
		Revision                    string `json:"revision,omitempty"`
		DetectedCI                  string `json:"detectedCI,omitempty"`
	} `json:"analyses,omitempty"`
}

// UpdateEventRequest Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
type UpdateEventRequest struct {
	Event string `json:"event"` // Event key
	Name  string `json:"name"`  // New name
}

// UpdateEventResponse is the response for UpdateEventRequest
type UpdateEventResponse struct {
	Event struct {
		Analysis string `json:"analysis,omitempty"`
		Category string `json:"category,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"event,omitempty"`
}
