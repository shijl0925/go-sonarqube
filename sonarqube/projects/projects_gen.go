package projects

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// BulkDeleteRequest Delete one or several projects.<br />Only the 1'000 first items in project filters are taken into account.<br />Requires 'Administer System' permission.<br />At least one parameter is required among analyzedBefore, projects and q
type BulkDeleteRequest struct {
	AnalyzedBefore    string `json:"analyzedBefore,omitempty"`    // Since 6.6;Filter the projects for which last analysis of any branch is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.
	OnProvisionedOnly string `json:"onProvisionedOnly,omitempty"` // Since 6.6;Filter the projects that are provisioned
	Projects          string `json:"projects,omitempty"`          // Comma-separated list of project keys
	Q                 string `json:"q,omitempty"`                 // Limit to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>
	Qualifiers        string `json:"qualifiers,omitempty"`        // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers
}

// CreateRequest Create a project.<br/>If your project is hosted on a DevOps Platform, please use the import endpoint under api/alm_integrations, so it creates and properly configures the project.Requires 'Create Projects' permission.<br/>
type CreateRequest struct {
	MainBranch             string `json:"mainBranch,omitempty"`             // Since 9.8;Key of the main branch of the project. If not provided, the default main branch key will be used.
	Name                   string `json:"name"`                             // Name of the project. If name is longer than 500, it is abbreviated.
	NewCodeDefinitionType  string `json:"newCodeDefinitionType,omitempty"`  // Since 10.1;Project New Code Definition Type<br/>New code definitions of the following types are allowed:<ul><li>PREVIOUS_VERSION</li><li>NUMBER_OF_DAYS</li><li>REFERENCE_BRANCH - will default to the main branch.</li></ul>
	NewCodeDefinitionValue string `json:"newCodeDefinitionValue,omitempty"` // Since 10.1;Project New Code Definition Value<br/>For each new code definition type, a different value is expected:<ul><li>no value, when the new code definition type is PREVIOUS_VERSION and REFERENCE_BRANCH</li><li>a number between 1 and 90, when the new code definition type is NUMBER_OF_DAYS</li></ul>
	Project                string `json:"project"`                          // Key of the project
	Visibility             string `json:"visibility,omitempty"`             // Since 6.4;Whether the created project should be visible to everyone, or only specific user/groups.<br/>If no visibility is specified, the default project visibility will be used.
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Project struct {
		Key       string `json:"key,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"project,omitempty"`
}

// DeleteRequest Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project.
type DeleteRequest struct {
	Project string `json:"project"` // Project key
}

// ExportFindingsRequest Export all findings (issues and hotspots) of a specific project branch.<br/>Requires 'Administer System' permission. Keep in mind that this endpoint will return all findings, issues and hotspots (no filter), which can take time and use a lot of resources on the SonarQube server side and put pressure on the database until completion. This endpoint can be used to feed third party systems. Either the branch key or the pull request key should be specified, and not both at the same time.
type ExportFindingsRequest struct {
	Branch      string `url:"branch,omitempty"`      // Branch key. When not specified, if no Pull Request key is defined either, it will default to the main branch
	Project     string `url:"project"`               // Project key
	PullRequest string `url:"pullRequest,omitempty"` // Pull Request key. When not specified, the branch data will be returned instead
}

// ExportFindingsResponse is the response for ExportFindingsRequest
type ExportFindingsResponse struct {
	ExportFindings []struct {
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		Branch                     string   `json:"branch,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []string `json:"comments,omitempty"`
		CreatedAt                  string   `json:"createdAt,omitempty"`
		Effort                     string   `json:"effort,omitempty"`
		Impacts                    []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus              string `json:"issueStatus,omitempty"`
		Key                      string `json:"key,omitempty"`
		LineNumber               string `json:"lineNumber,omitempty"`
		Message                  string `json:"message,omitempty"`
		Path                     string `json:"path,omitempty"`
		ProjectKey               string `json:"projectKey,omitempty"`
		RuleReference            string `json:"ruleReference,omitempty"`
		Tags                     string `json:"tags,omitempty"`
		UpdatedAt                string `json:"updatedAt,omitempty"`
		PullRequest              string `json:"pullRequest,omitempty"`
		Resolution               string `json:"resolution,omitempty"`
		SecurityCategory         string `json:"securityCategory,omitempty"`
		Status                   string `json:"status,omitempty"`
		Type                     string `json:"type,omitempty"`
		VulnerabilityProbability string `json:"vulnerabilityProbability,omitempty"`
	} `json:"export_findings,omitempty"`
}

// LicenseUsageRequest Help admins to understand how much each project affects the total number of lines of code. Returns the list of projects together with information about their usage, sorted by lines of code descending.<br/>Requires Administer System permission.
type LicenseUsageRequest struct{}

// LicenseUsageResponse is the response for LicenseUsageRequest
type LicenseUsageResponse struct {
	Projects []struct {
		Branch                 string  `json:"branch,omitempty"`
		LicenseUsagePercentage float64 `json:"licenseUsagePercentage,omitempty"`
		LinesOfCode            float64 `json:"linesOfCode,omitempty"`
		ProjectKey             string  `json:"projectKey,omitempty"`
		ProjectName            string  `json:"projectName,omitempty"`
		PullRequest            string  `json:"pullRequest,omitempty"`
	} `json:"projects,omitempty"`
}

/*
SearchRequest Search for projects or views to administrate them.
<ul>
  <li>The response field 'lastAnalysisDate' takes into account the analysis of all branches and pull requests, not only the main branch.</li>
  <li>The response field 'revision' takes into account the analysis of the main branch only.</li>
</ul>
Requires 'Administer System' permission
*/
type SearchRequest struct {
	AnalyzedBefore    string `url:"analyzedBefore,omitempty"`    // Since 6.6;Filter the projects for which the last analysis of all branches are older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.
	OnProvisionedOnly string `url:"onProvisionedOnly,omitempty"` // Since 6.6;Filter the projects that are provisioned
	Projects          string `url:"projects,omitempty"`          // Since 6.6;Comma-separated list of project keys
	Q                 string `url:"q,omitempty"`                 // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>
	Qualifiers        string `url:"qualifiers,omitempty"`        // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Components []struct {
		Key              string `json:"key,omitempty"`
		LastAnalysisDate string `json:"lastAnalysisDate,omitempty"`
		Managed          bool   `json:"managed,omitempty"`
		Name             string `json:"name,omitempty"`
		Qualifier        string `json:"qualifier,omitempty"`
		Revision         string `json:"revision,omitempty"`
		Visibility       string `json:"visibility,omitempty"`
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
		Key              string `json:"key,omitempty"`
		LastAnalysisDate string `json:"lastAnalysisDate,omitempty"`
		Managed          bool   `json:"managed,omitempty"`
		Name             string `json:"name,omitempty"`
		Qualifier        string `json:"qualifier,omitempty"`
		Revision         string `json:"revision,omitempty"`
		Visibility       string `json:"visibility,omitempty"`
	} `json:"components,omitempty"`
}

// UpdateKeyRequest Update a project all its sub-components keys.<br>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
type UpdateKeyRequest struct {
	From string `json:"from"` // Project key
	To   string `json:"to"`   // New project key
}

// UpdateVisibilityRequest Updates visibility of a project, application or a portfolio.<br>Requires 'Project administer' permission on the specified entity
type UpdateVisibilityRequest struct {
	Project    string `json:"project"`    // Project, application or portfolio key
	Visibility string `json:"visibility"` // New visibility
}
