package issues

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddCommentRequest Add a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
type AddCommentRequest struct {
	Issue string `form:"issue"` // Issue key
	Text  string `form:"text"`  // Comment text
}

// AddCommentResponse is the response for AddCommentRequest
type AddCommentResponse struct {
	Components []struct {
		Enabled   bool    `json:"enabled,omitempty"`
		Key       string  `json:"key,omitempty"`
		LongName  string  `json:"longName,omitempty"`
		Name      string  `json:"name,omitempty"`
		Path      string  `json:"path,omitempty"`
		Qualifier string  `json:"qualifier,omitempty"`
		Id        float64 `json:"id,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// AssignRequest Assign/Unassign an issue. Requires authentication and Browse permission on project
type AssignRequest struct {
	Assignee string `form:"assignee,omitempty"` // Login of the assignee. When not set, it will unassign the issue. Use '_me' to assign to current user
	Issue    string `form:"issue"`              // Issue key
}

// AssignResponse is the response for AssignRequest
type AssignResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// AuthorsRequest Search SCM accounts which match a given query.<br/>Requires authentication.<br/>When issue indexing is in progress returns 503 service unavailable HTTP code.
type AuthorsRequest struct {
	Project string `url:"project,omitempty"` // Since 7.4;Project key
	Q       string `url:"q,omitempty"`       // Limit search to authors that contain the supplied string.
}

// AuthorsResponse is the response for AuthorsRequest
type AuthorsResponse struct {
	Authors []string `json:"authors,omitempty"`
}

// BulkChangeRequest Bulk change on issues. Up to 500 issues can be updated. <br/>Requires authentication.
type BulkChangeRequest struct {
	AddTags           string `form:"add_tags,omitempty"`          // Add tags
	Assign            string `form:"assign,omitempty"`            // To assign the list of issues to a specific user (login), or un-assign all the issues
	Comment           string `form:"comment,omitempty"`           // Add a comment. The comment will only be added to issues that are affected either by a change of type or a change of severity as a result of the same WS call.
	DoTransition      string `form:"do_transition,omitempty"`     // Transition
	Issues            string `form:"issues"`                      // Comma-separated list of issue keys
	RemoveTags        string `form:"remove_tags,omitempty"`       // Remove tags
	SendNotifications string `form:"sendNotifications,omitempty"` // Since 4.0;
	SetSeverity       string `form:"set_severity,omitempty"`      // To change the severity of the list of issues
	SetType           string `form:"set_type,omitempty"`          // Since 5.5;To change the type of the list of issues
}

// BulkChangeResponse is the response for BulkChangeRequest
type BulkChangeResponse struct {
	Failures float64 `json:"failures,omitempty"`
	Ignored  float64 `json:"ignored,omitempty"`
	Success  float64 `json:"success,omitempty"`
	Total    float64 `json:"total,omitempty"`
}

// ChangelogRequest Display changelog of an issue.<br/>Requires the 'Browse' permission on the project of the specified issue.
type ChangelogRequest struct {
	Issue string `url:"issue"` // Issue key
}

// ChangelogResponse is the response for ChangelogRequest
type ChangelogResponse struct {
	Changelog []struct {
		Avatar       string `json:"avatar,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		Diffs        []struct {
			Key      string `json:"key,omitempty"`
			NewValue string `json:"newValue,omitempty"`
			OldValue string `json:"oldValue,omitempty"`
		} `json:"diffs,omitempty"`
		ExternalUser  string `json:"externalUser,omitempty"`
		IsUserActive  bool   `json:"isUserActive,omitempty"`
		User          string `json:"user,omitempty"`
		UserName      string `json:"userName,omitempty"`
		WebhookSource string `json:"webhookSource,omitempty"`
	} `json:"changelog,omitempty"`
}

// DeleteCommentRequest Delete a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
type DeleteCommentRequest struct {
	Comment string `form:"comment"` // Since 6.3;Comment key
}

// DeleteCommentResponse is the response for DeleteCommentRequest
type DeleteCommentResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

/*
DoTransitionRequest Do workflow transition on an issue. Requires authentication and Browse permission on project.<br/>
The transitions 'accept', 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.<br/>
The transitions involving security hotspots require the permission 'Administer Security Hotspot'.
*/
type DoTransitionRequest struct {
	Issue      string `form:"issue"`      // Issue key
	Transition string `form:"transition"` // Transition
}

// DoTransitionResponse is the response for DoTransitionRequest
type DoTransitionResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// EditCommentRequest Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
type EditCommentRequest struct {
	Comment string `form:"comment"` // Since 6.3;Comment key
	Text    string `form:"text"`    // Comment text
}

// EditCommentResponse is the response for EditCommentRequest
type EditCommentResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// GitlabSastExportRequest Return a list of vulnerabilities according to the Gitlab SAST JSON format.<br> The JSON produced can be used in GitLab for generating the Vulnerability Report.Requires the 'Browse' or 'Scan' permission on the specified project.
type GitlabSastExportRequest struct {
	Branch      string `url:"branch,omitempty"`      // Branch key.If this parameter is set, pullRequest must not be set.
	ProjectKey  string `url:"projectKey"`            // The project key for which the vulnerabilities are being fetched
	PullRequest string `url:"pullRequest,omitempty"` // Pull request id.If this parameter is set, branch must not be set.
}

// GitlabSastExportResponse is the response for GitlabSastExportRequest
type GitlabSastExportResponse struct {
	Remediations []string `json:"remediations,omitempty"`
	Scan         struct {
		Analyzer struct {
			Id     string `json:"id,omitempty"`
			Name   string `json:"name,omitempty"`
			Vendor struct {
				Name string `json:"name,omitempty"`
			} `json:"vendor,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"analyzer,omitempty"`
		EndTime  string   `json:"end_time,omitempty"`
		Messages []string `json:"messages,omitempty"`
		Scanner  struct {
			Id     string `json:"id,omitempty"`
			Name   string `json:"name,omitempty"`
			Vendor struct {
				Name string `json:"name,omitempty"`
			} `json:"vendor,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"scanner,omitempty"`
		StartTime string `json:"start_time,omitempty"`
		Status    string `json:"status,omitempty"`
		Type      string `json:"type,omitempty"`
	} `json:"scan,omitempty"`
	Version         string `json:"version,omitempty"`
	Vulnerabilities []struct {
		Category    string `json:"category,omitempty"`
		Description string `json:"description,omitempty"`
		Id          string `json:"id,omitempty"`
		Identifiers []struct {
			Name  string `json:"name,omitempty"`
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"identifiers,omitempty"`
		Links []struct {
			Name string `json:"name,omitempty"`
			Url  string `json:"url,omitempty"`
		} `json:"links,omitempty"`
		Location struct {
			File      string  `json:"file,omitempty"`
			StartLine float64 `json:"start_line,omitempty"`
		} `json:"location,omitempty"`
		Message string `json:"message,omitempty"`
		Name    string `json:"name,omitempty"`
		Scanner struct {
			Id     string `json:"id,omitempty"`
			Name   string `json:"name,omitempty"`
			Vendor struct {
				Name string `json:"name,omitempty"`
			} `json:"vendor,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"scanner,omitempty"`
		Severity string `json:"severity,omitempty"`
	} `json:"vulnerabilities,omitempty"`
}

// ReindexRequest Reindex issues for a project.<br> Require 'Administer System' permission.
type ReindexRequest struct {
	Project string `form:"project"` // Project key
}

// SearchRequest Search for issues.<br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects.<br/>When issue indexing is in progress returns 503 service unavailable HTTP code.
type SearchRequest struct {
	AdditionalFields             string `url:"additionalFields,omitempty"`             // Since 5.2;Comma-separated list of the optional fields to be returned in response. Action plans are dropped in 5.5, it is not returned in the response.
	Asc                          string `url:"asc,omitempty"`                          // Ascending sort
	Assigned                     string `url:"assigned,omitempty"`                     // To retrieve assigned or unassigned issues
	Assignees                    string `url:"assignees,omitempty"`                    // Comma-separated list of assignee logins. The value '__me__' can be used as a placeholder for user who performs the request
	Author                       string `url:"author,omitempty"`                       // SCM accounts. To set several values, the parameter must be called once for each value.
	Branch                       string `url:"branch,omitempty"`                       // Since 6.6;Branch key. Not available in the community edition.
	Casa                         string `url:"casa,omitempty"`                         // Since 10.7;Comma-separated list of CASA categories.
	CleanCodeAttributeCategories string `url:"cleanCodeAttributeCategories,omitempty"` // Since 10.2;Comma-separated list of Clean Code Attribute Categories
	CodeVariants                 string `url:"codeVariants,omitempty"`                 // Since 10.1;Comma-separated list of code variants.
	ComplianceStandards          string `url:"complianceStandards,omitempty"`          // Since 2025.6;Comma-separated list of compliance standards
	Components                   string `url:"components,omitempty"`                   // Comma-separated list of component keys. Retrieve issues associated to a specific list of components (and all its descendants). A component can be a portfolio, project, module, directory or file.
	CreatedAfter                 string `url:"createdAfter,omitempty"`                 // To retrieve issues created after the given date (inclusive). <br>Either a date (use 'timeZone' attribute or it will default to server timezone) or datetime can be provided. <br>If this parameter is set, createdInLast must not be set
	CreatedAt                    string `url:"createdAt,omitempty"`                    // Datetime to retrieve issues created during a specific analysis
	CreatedBefore                string `url:"createdBefore,omitempty"`                // To retrieve issues created before the given date (exclusive). <br>Either a date (use 'timeZone' attribute or it will default to server timezone) or datetime can be provided.
	CreatedInLast                string `url:"createdInLast,omitempty"`                // To retrieve issues created during a time span before the current time (exclusive). Accepted units are 'y' for year, 'm' for month, 'w' for week and 'd' for day. If this parameter is set, createdAfter must not be set
	Cwe                          string `url:"cwe,omitempty"`                          // Comma-separated list of CWE identifiers. Use 'unknown' to select issues not associated to any CWE.
	Facets                       string `url:"facets,omitempty"`                       // Comma-separated list of the facets to be computed. No facet is computed by default.
	FixedInPullRequest           string `url:"fixedInPullRequest,omitempty"`           // Since 10.4;Pull request id to filter issues that would be fixed in the specified project or branch by the pull request. Should not be used together with + 'pullRequest'. At least the 'components' must be be specified when this param is used.  Not available in the community edition.
	FromSonarQubeUpdate          string `url:"fromSonarQubeUpdate,omitempty"`          // To match issues detected because of SonarQube updates
	ImpactSeverities             string `url:"impactSeverities,omitempty"`             // Since 10.2;Comma-separated list of Software Quality Severities
	ImpactSoftwareQualities      string `url:"impactSoftwareQualities,omitempty"`      // Since 10.2;Comma-separated list of Software Qualities
	InNewCodePeriod              string `url:"inNewCodePeriod,omitempty"`              // Since 9.4;To retrieve issues created in the new code period.<br>If this parameter is set to a truthy value, createdAfter must not be set and one component uuid or key must be provided.<br>This parameter is ignored if the requested component is Application or Portfolio.
	IssueStatuses                string `url:"issueStatuses,omitempty"`                // Since 10.4;
	Issues                       string `url:"issues,omitempty"`                       // Comma-separated list of issue keys
	Languages                    string `url:"languages,omitempty"`                    // Comma-separated list of languages. Available since 4.4
	LinkedTicketStatus           string `url:"linkedTicketStatus,omitempty"`           // Has linked JIRA work item
	OnComponentOnly              string `url:"onComponentOnly,omitempty"`              // Return only issues at a component's level, not on its descendants (modules, directories, files, etc). This parameter is only considered when componentKeys is set.
	OwaspAsvs40                  string `url:"owaspAsvs-4.0,omitempty"`                // Since 9.7;Comma-separated list of OWASP ASVS v4.0 categories.
	OwaspAsvsLevel               string `url:"owaspAsvsLevel,omitempty"`               // Since 9.7;Level of OWASP ASVS categories.
	OwaspMobileTop102024         string `url:"owaspMobileTop10-2024,omitempty"`        // Since 2025.3;Comma-separated list of OWASP Mobile Top 10 2024 lowercase categories.
	OwaspTop10                   string `url:"owaspTop10,omitempty"`                   // Since 7.3;Comma-separated list of OWASP Top 10 2017 lowercase categories.
	OwaspTop102021               string `url:"owaspTop10-2021,omitempty"`              // Since 9.4;Comma-separated list of OWASP Top 10 2021 lowercase categories.
	PciDss32                     string `url:"pciDss-3.2,omitempty"`                   // Since 9.6;Comma-separated list of PCI DSS v3.2 categories.
	PciDss40                     string `url:"pciDss-4.0,omitempty"`                   // Since 9.6;Comma-separated list of PCI DSS v4.0 categories.
	PrioritizedRule              string `url:"prioritizedRule,omitempty"`              // To match issues with prioritized rule or not
	PullRequest                  string `url:"pullRequest,omitempty"`                  // Since 7.1;Pull request id. Not available in the community edition.
	Resolutions                  string `url:"resolutions,omitempty"`                  // Comma-separated list of resolutions
	Resolved                     string `url:"resolved,omitempty"`                     // To match resolved or unresolved issues
	Rules                        string `url:"rules,omitempty"`                        // Comma-separated list of coding rule keys. Format is &lt;repository&gt;:&lt;rule&gt;
	S                            string `url:"s,omitempty"`                            // Sort field
	SansTop25                    string `url:"sansTop25,omitempty"`                    // Since 7.3;Deprecated since 10.0;Comma-separated list of SANS Top 25 categories.
	Scopes                       string `url:"scopes,omitempty"`                       // Comma-separated list of scopes. Available since 8.5
	Severities                   string `url:"severities,omitempty"`                   // Comma-separated list of severities
	SonarsourceSecurity          string `url:"sonarsourceSecurity,omitempty"`          // Since 7.8;Comma-separated list of SonarSource security categories. Use 'others' to select issues not associated with any category
	Statuses                     string `url:"statuses,omitempty"`                     // Deprecated since 10.4;Comma-separated list of statuses
	StigASDV5R3                  string `url:"stig-ASD_V5R3,omitempty"`                // Since 10.7;Comma-separated list of STIG V5R3 categories.
	Tags                         string `url:"tags,omitempty"`                         // Comma-separated list of tags.
	TimeZone                     string `url:"timeZone,omitempty"`                     // Since 8.6;To resolve dates passed to 'createdAfter' or 'createdBefore' (does not apply to datetime) and to compute creation date histogram
	Types                        string `url:"types,omitempty"`                        // Since 5.5;Comma-separated list of types.
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Facets []struct {
		Property string `json:"property,omitempty"`
		Values   []struct {
			Count float64 `json:"count,omitempty"`
			Val   string  `json:"val,omitempty"`
		} `json:"values,omitempty"`
	} `json:"facets,omitempty"`
	Issues []struct {
		Actions                    []string `json:"actions,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		CodeVariants               []string `json:"codeVariants,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string `json:"component,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		Effort       string `json:"effort,omitempty"`
		Flows        []struct {
			Locations []struct {
				Msg            string `json:"msg,omitempty"`
				MsgFormattings []struct {
					End   float64 `json:"end,omitempty"`
					Start float64 `json:"start,omitempty"`
					Type  string  `json:"type,omitempty"`
				} `json:"msgFormattings,omitempty"`
				TextRange struct {
					EndLine     float64 `json:"endLine,omitempty"`
					EndOffset   float64 `json:"endOffset,omitempty"`
					StartLine   float64 `json:"startLine,omitempty"`
					StartOffset float64 `json:"startOffset,omitempty"`
				} `json:"textRange,omitempty"`
			} `json:"locations,omitempty"`
		} `json:"flows,omitempty"`
		Hash    string `json:"hash,omitempty"`
		Impacts []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		InternalTags       []string `json:"internalTags,omitempty"`
		IssueStatus        string   `json:"issueStatus,omitempty"`
		Key                string   `json:"key,omitempty"`
		Line               float64  `json:"line,omitempty"`
		LinkedTicketStatus string   `json:"linkedTicketStatus,omitempty"`
		Message            string   `json:"message,omitempty"`
		MessageFormattings []struct {
			End   float64 `json:"end,omitempty"`
			Start float64 `json:"start,omitempty"`
			Type  string  `json:"type,omitempty"`
		} `json:"messageFormattings,omitempty"`
		PrioritizedRule           bool     `json:"prioritizedRule,omitempty"`
		Project                   string   `json:"project,omitempty"`
		QuickFixAvailable         bool     `json:"quickFixAvailable,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issues,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
	Rules  []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Avatar string `json:"avatar,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Facets []struct {
		Property string `json:"property,omitempty"`
		Values   []struct {
			Count float64 `json:"count,omitempty"`
			Val   string  `json:"val,omitempty"`
		} `json:"values,omitempty"`
	} `json:"facets,omitempty"`
	Issues []struct {
		Actions                    []string `json:"actions,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		CodeVariants               []string `json:"codeVariants,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string `json:"component,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		Effort       string `json:"effort,omitempty"`
		Flows        []struct {
			Locations []struct {
				Msg            string `json:"msg,omitempty"`
				MsgFormattings []struct {
					End   float64 `json:"end,omitempty"`
					Start float64 `json:"start,omitempty"`
					Type  string  `json:"type,omitempty"`
				} `json:"msgFormattings,omitempty"`
				TextRange struct {
					EndLine     float64 `json:"endLine,omitempty"`
					EndOffset   float64 `json:"endOffset,omitempty"`
					StartLine   float64 `json:"startLine,omitempty"`
					StartOffset float64 `json:"startOffset,omitempty"`
				} `json:"textRange,omitempty"`
			} `json:"locations,omitempty"`
		} `json:"flows,omitempty"`
		Hash    string `json:"hash,omitempty"`
		Impacts []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		InternalTags       []string `json:"internalTags,omitempty"`
		IssueStatus        string   `json:"issueStatus,omitempty"`
		Key                string   `json:"key,omitempty"`
		Line               float64  `json:"line,omitempty"`
		LinkedTicketStatus string   `json:"linkedTicketStatus,omitempty"`
		Message            string   `json:"message,omitempty"`
		MessageFormattings []struct {
			End   float64 `json:"end,omitempty"`
			Start float64 `json:"start,omitempty"`
			Type  string  `json:"type,omitempty"`
		} `json:"messageFormattings,omitempty"`
		PrioritizedRule           bool     `json:"prioritizedRule,omitempty"`
		Project                   string   `json:"project,omitempty"`
		QuickFixAvailable         bool     `json:"quickFixAvailable,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issues,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Avatar string `json:"avatar,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// SetSeverityRequest Change severity.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
type SetSeverityRequest struct {
	Impact   string `form:"impact,omitempty"`   // Override of impact severity for the rule. Cannot be used as the same time as 'severity'
	Issue    string `form:"issue"`              // Issue key
	Severity string `form:"severity,omitempty"` // New severity
}

// SetSeverityResponse is the response for SetSeverityRequest
type SetSeverityResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Severity                  string   `json:"severity,omitempty"`
		Status                    string   `json:"status,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// SetTagsRequest Set tags on an issue. <br/>Requires authentication and Browse permission on project
type SetTagsRequest struct {
	Issue string `form:"issue"`          // Since 6.3;Issue key
	Tags  string `form:"tags,omitempty"` // Comma-separated list of tags. All tags are removed if parameter is empty or not set.
}

// SetTagsResponse is the response for SetTagsRequest
type SetTagsResponse struct {
	Components []struct {
		Enabled      bool    `json:"enabled,omitempty"`
		Key          string  `json:"key,omitempty"`
		LongName     string  `json:"longName,omitempty"`
		Name         string  `json:"name,omitempty"`
		Path         string  `json:"path,omitempty"`
		ProjectId    float64 `json:"projectId,omitempty"`
		Qualifier    string  `json:"qualifier,omitempty"`
		SubProjectId float64 `json:"subProjectId,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// SetTypeRequest Change type of issue, for instance from 'code smell' to 'bug'.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
// Deprecated: this action has been deprecated since version 10.2
type SetTypeRequest struct {
	Issue string `form:"issue"` // Issue key
	Type  string `form:"type"`  // New type
}

// SetTypeResponse is the response for SetTypeRequest
type SetTypeResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions                    []string `json:"actions,omitempty"`
		Assignee                   string   `json:"assignee,omitempty"`
		Author                     string   `json:"author,omitempty"`
		CleanCodeAttribute         string   `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string   `json:"cleanCodeAttributeCategory,omitempty"`
		Comments                   []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Impacts      []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IssueStatus               string   `json:"issueStatus,omitempty"`
		Key                       string   `json:"key,omitempty"`
		Line                      float64  `json:"line,omitempty"`
		Message                   string   `json:"message,omitempty"`
		Project                   string   `json:"project,omitempty"`
		Rule                      string   `json:"rule,omitempty"`
		RuleDescriptionContextKey string   `json:"ruleDescriptionContextKey,omitempty"`
		Severity                  string   `json:"severity,omitempty"`
		Status                    string   `json:"status,omitempty"`
		Tags                      []string `json:"tags,omitempty"`
		TextRange                 struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// TagsRequest List tags matching a given query
type TagsRequest struct {
	All     string `url:"all,omitempty"`     // Since 9.2;Indicator to search for all tags or only for tags in the main branch of a project
	Branch  string `url:"branch,omitempty"`  // Since 9.2;Branch key
	Project string `url:"project,omitempty"` // Since 7.4;Project key
	Q       string `url:"q,omitempty"`       // Limit search to tags that contain the supplied string.
}

// TagsResponse is the response for TagsRequest
type TagsResponse struct {
	Tags []string `json:"tags,omitempty"`
}
