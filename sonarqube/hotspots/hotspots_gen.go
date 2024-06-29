package hotspots

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ChangeStatusRequest Change the status of a Security Hotpot.<br/>Requires the 'Administer Security Hotspot' permission.
type ChangeStatusRequest struct {
	Comment    string `form:"comment,omitempty"`    // Comment text.
	Hotspot    string `form:"hotspot,omitempty"`    // Key of the Security Hotspot
	Resolution string `form:"resolution,omitempty"` // Resolution of the Security Hotspot when new status is REVIEWED, otherwise must not be set.
	Status     string `form:"status,omitempty"`     // New status of the Security Hotspot.
}

// SearchRequest Search for Security Hotpots. <br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects. <br>When issue indexing is in progress returns 503 service unavailable HTTP code.
type SearchRequest struct {
	Branch              string `form:"branch,omitempty"`              // Branch key. Not available in the community edition.
	Cwe                 string `form:"cwe,omitempty"`                 // Comma-separated list of CWE numbers
	Files               string `form:"files,omitempty"`               // Comma-separated list of files. Returns only hotspots found in those files
	Hotspots            string `form:"hotspots,omitempty"`            // Comma-separated list of Security Hotspot keys. This parameter is required unless project is provided.
	InNewCodePeriod     string `form:"inNewCodePeriod,omitempty"`     // If 'inNewCodePeriod' is provided, only Security Hotspots created in the new code period are returned.
	OnlyMine            string `form:"onlyMine,omitempty"`            // If 'projectKey' is provided, returns only Security Hotspots assigned to the current user
	OwaspAsvs40         string `form:"owaspAsvs-4.0,omitempty"`       // Comma-separated list of OWASP ASVS v4.0 categories or rules.
	OwaspAsvsLevel      string `form:"owaspAsvsLevel,omitempty"`      // Filters hotspots with lower or equal OWASP ASVS level to the parameter value. Should be used in combination with the 'owaspAsvs-4.0' parameter.
	OwaspTop10          string `form:"owaspTop10,omitempty"`          // Comma-separated list of OWASP 2017 Top 10 lowercase categories.
	OwaspTop102021      string `form:"owaspTop10-2021,omitempty"`     // Comma-separated list of OWASP 2021 Top 10 lowercase categories.
	PciDss32            string `form:"pciDss-3.2,omitempty"`          // Comma-separated list of PCI DSS v3.2 categories.
	PciDss40            string `form:"pciDss-4.0,omitempty"`          // Comma-separated list of PCI DSS v4.0 categories.
	Project             string `form:"project,omitempty"`             // Key of the project or application. This parameter is required unless hotspots is provided.
	PullRequest         string `form:"pullRequest,omitempty"`         // Pull request id. Not available in the community edition.
	Resolution          string `form:"resolution,omitempty"`          // If 'project' is provided and if status is 'REVIEWED', only Security Hotspots with the specified resolution are returned.
	SansTop25           string `form:"sansTop25,omitempty"`           // Comma-separated list of SANS Top 25 categories.
	SonarsourceSecurity string `form:"sonarsourceSecurity,omitempty"` // Comma-separated list of SonarSource security categories. Use 'others' to select issues not associated with any category
	Status              string `form:"status,omitempty"`              // If 'project' is provided, only Security Hotspots with the specified status are returned.
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Components []struct {
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Hotspots []struct {
		Assignee                 string   `json:"assignee,omitempty"`
		Author                   string   `json:"author,omitempty"`
		Component                string   `json:"component,omitempty"`
		CreationDate             string   `json:"creationDate,omitempty"`
		Flows                    []string `json:"flows,omitempty"`
		Key                      string   `json:"key,omitempty"`
		Line                     float64  `json:"line,omitempty"`
		Message                  string   `json:"message,omitempty"`
		MessageFormattings       []string `json:"messageFormattings,omitempty"`
		Project                  string   `json:"project,omitempty"`
		RuleKey                  string   `json:"ruleKey,omitempty"`
		SecurityCategory         string   `json:"securityCategory,omitempty"`
		Status                   string   `json:"status,omitempty"`
		UpdateDate               string   `json:"updateDate,omitempty"`
		VulnerabilityProbability string   `json:"vulnerabilityProbability,omitempty"`
	} `json:"hotspots,omitempty"`
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
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Hotspots []struct {
		Assignee                 string   `json:"assignee,omitempty"`
		Author                   string   `json:"author,omitempty"`
		Component                string   `json:"component,omitempty"`
		CreationDate             string   `json:"creationDate,omitempty"`
		Flows                    []string `json:"flows,omitempty"`
		Key                      string   `json:"key,omitempty"`
		Line                     float64  `json:"line,omitempty"`
		Message                  string   `json:"message,omitempty"`
		MessageFormattings       []string `json:"messageFormattings,omitempty"`
		Project                  string   `json:"project,omitempty"`
		RuleKey                  string   `json:"ruleKey,omitempty"`
		SecurityCategory         string   `json:"securityCategory,omitempty"`
		Status                   string   `json:"status,omitempty"`
		UpdateDate               string   `json:"updateDate,omitempty"`
		VulnerabilityProbability string   `json:"vulnerabilityProbability,omitempty"`
	} `json:"hotspots,omitempty"`
}

// ShowRequest Provides the details of a Security Hotspot.
type ShowRequest struct {
	Hotspot string `form:"hotspot,omitempty"` // Key of the Security Hotspot
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Assignee        string `json:"assignee,omitempty"`
	Author          string `json:"author,omitempty"`
	CanChangeStatus bool   `json:"canChangeStatus,omitempty"`
	Changelog       []struct {
		Avatar       string `json:"avatar,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		Diffs        []struct {
			Key      string `json:"key,omitempty"`
			NewValue string `json:"newValue,omitempty"`
			OldValue string `json:"oldValue,omitempty"`
		} `json:"diffs,omitempty"`
		IsUserActive bool   `json:"isUserActive,omitempty"`
		User         string `json:"user,omitempty"`
		UserName     string `json:"userName,omitempty"`
	} `json:"changelog,omitempty"`
	CodeVariants []string `json:"codeVariants,omitempty"`
	Comment      []struct {
		CreatedAt string `json:"createdAt,omitempty"`
		HtmlText  string `json:"htmlText,omitempty"`
		Key       string `json:"key,omitempty"`
		Login     string `json:"login,omitempty"`
		Markdown  string `json:"markdown,omitempty"`
	} `json:"comment,omitempty"`
	Component struct {
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"component,omitempty"`
	CreationDate       string  `json:"creationDate,omitempty"`
	Hash               string  `json:"hash,omitempty"`
	Key                string  `json:"key,omitempty"`
	Line               float64 `json:"line,omitempty"`
	Message            string  `json:"message,omitempty"`
	MessageFormattings []struct {
		End   float64 `json:"end,omitempty"`
		Start float64 `json:"start,omitempty"`
		Type  string  `json:"type,omitempty"`
	} `json:"messageFormattings,omitempty"`
	Project struct {
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"project,omitempty"`
	Rule struct {
		Key                      string `json:"key,omitempty"`
		Name                     string `json:"name,omitempty"`
		SecurityCategory         string `json:"securityCategory,omitempty"`
		VulnerabilityProbability string `json:"vulnerabilityProbability,omitempty"`
	} `json:"rule,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdateDate string `json:"updateDate,omitempty"`
	Users      []struct {
		Active bool   `json:"active,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}
