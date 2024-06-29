package rules

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// CreateRequest Create a custom rule.<br>Requires the 'Administer Quality Profiles' permission
type CreateRequest struct {
	CleanCodeAttribute  string `form:"cleanCodeAttribute,omitempty"`  // Clean code attribute
	CustomKey           string `form:"customKey,omitempty"`           // Key of the custom rule
	Impacts             string `form:"impacts,omitempty"`             // Impacts as semi-colon list of &lt;software_quality&gt;=&lt;severity&gt;
	MarkdownDescription string `form:"markdownDescription,omitempty"` // Rule description in <a href='/formatting/help'>markdown format</a>
	Name                string `form:"name,omitempty"`                // Rule name
	Params              string `form:"params,omitempty"`              // Parameters as semi-colon list of &lt;key&gt;=&lt;value&gt;
	PreventReactivation string `form:"preventReactivation,omitempty"` // If set to true and if the rule has been deactivated (status 'REMOVED'), a status 409 will be returned
	Severity            string `form:"severity,omitempty"`            // Rule severity
	Status              string `form:"status,omitempty"`              // Rule status
	TemplateKey         string `form:"templateKey,omitempty"`         // Key of the template rule in order to create a custom rule
	Type                string `form:"type,omitempty"`                // Rule type
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Rule struct {
		CleanCodeAttribute         string `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string `json:"cleanCodeAttributeCategory,omitempty"`
		CreatedAt                  string `json:"createdAt,omitempty"`
		HtmlDesc                   string `json:"htmlDesc,omitempty"`
		Impacts                    []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IsExternal bool   `json:"isExternal,omitempty"`
		IsTemplate bool   `json:"isTemplate,omitempty"`
		Key        string `json:"key,omitempty"`
		Lang       string `json:"lang,omitempty"`
		LangName   string `json:"langName,omitempty"`
		MdDesc     string `json:"mdDesc,omitempty"`
		Name       string `json:"name,omitempty"`
		Params     []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			HtmlDesc     string `json:"htmlDesc,omitempty"`
			Key          string `json:"key,omitempty"`
			Type         string `json:"type,omitempty"`
		} `json:"params,omitempty"`
		Repo        string   `json:"repo,omitempty"`
		Scope       string   `json:"scope,omitempty"`
		Severity    string   `json:"severity,omitempty"`
		Status      string   `json:"status,omitempty"`
		SysTags     []string `json:"sysTags,omitempty"`
		TemplateKey string   `json:"templateKey,omitempty"`
		Type        string   `json:"type,omitempty"`
	} `json:"rule,omitempty"`
}

// DeleteRequest Delete custom rule.<br/>Requires the 'Administer Quality Profiles' permission
type DeleteRequest struct {
	Key string `form:"key,omitempty"` // Rule key
}

// RepositoriesRequest List available rule repositories
type RepositoriesRequest struct {
	Language string `form:"language,omitempty"` // A language key; if provided, only repositories for the given language will be returned
	Q        string `form:"q,omitempty"`        // A pattern to match repository keys/names against
}

// RepositoriesResponse is the response for RepositoriesRequest
type RepositoriesResponse struct {
	Repositories []struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"repositories,omitempty"`
}

// SearchRequest Search for a collection of relevant rules matching a specified query.<br/>
type SearchRequest struct {
	Activation                   string `form:"activation,omitempty"`                   // Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.
	ActiveSeverities             string `form:"active_severities,omitempty"`            // Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.
	Asc                          string `form:"asc,omitempty"`                          // Ascending sort
	AvailableSince               string `form:"available_since,omitempty"`              // Filters rules added since date. Format is yyyy-MM-dd
	CleanCodeAttributeCategories string `form:"cleanCodeAttributeCategories,omitempty"` // Comma-separated list of Clean Code Attribute Categories
	Cwe                          string `form:"cwe,omitempty"`                          // Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE.
	F                            string `form:"f,omitempty"`                            // Comma-separated list of additional fields to be returned in the response. All the fields are returned by default, except actives.
	Facets                       string `form:"facets,omitempty"`                       // Comma-separated list of the facets to be computed. No facet is computed by default.
	ImpactSeverities             string `form:"impactSeverities,omitempty"`             // Comma-separated list of Software Quality Severities
	ImpactSoftwareQualities      string `form:"impactSoftwareQualities,omitempty"`      // Comma-separated list of Software Qualities
	IncludeExternal              string `form:"include_external,omitempty"`             // Include external engine rules in the results
	Inheritance                  string `form:"inheritance,omitempty"`                  // Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.
	IsTemplate                   string `form:"is_template,omitempty"`                  // Filter template rules
	Languages                    string `form:"languages,omitempty"`                    // Comma-separated list of languages
	OwaspTop10                   string `form:"owaspTop10,omitempty"`                   // Comma-separated list of OWASP Top 10 2017 lowercase categories.
	OwaspTop102021               string `form:"owaspTop10-2021,omitempty"`              // Comma-separated list of OWASP Top 10 2021 lowercase categories.
	PrioritizedRule              string `form:"prioritizedRule,omitempty"`              // Filter on prioritized rules. Ignored if the parameter 'qprofile' is not set.
	Q                            string `form:"q,omitempty"`                            // UTF-8 search query
	Qprofile                     string `form:"qprofile,omitempty"`                     // Quality profile key to filter on. Used only if the parameter 'activation' is set.
	Repositories                 string `form:"repositories,omitempty"`                 // Comma-separated list of repositories
	RuleKey                      string `form:"rule_key,omitempty"`                     // Key of rule to search for
	S                            string `form:"s,omitempty"`                            // Sort field
	SansTop25                    string `form:"sansTop25,omitempty"`                    // Comma-separated list of SANS Top 25 categories.
	Severities                   string `form:"severities,omitempty"`                   // Comma-separated list of default severities. Not the same than severity of rules in Quality profiles.
	SonarsourceSecurity          string `form:"sonarsourceSecurity,omitempty"`          // Comma-separated list of SonarSource security categories. Use 'others' to select rules not associated with any category
	Statuses                     string `form:"statuses,omitempty"`                     // Comma-separated list of status codes
	Tags                         string `form:"tags,omitempty"`                         // Comma-separated list of tags. Returned rules match any of the tags (OR operator)
	TemplateKey                  string `form:"template_key,omitempty"`                 // Key of the template rule to filter on. Used to search for the custom rules based on this template.
	Types                        string `form:"types,omitempty"`                        // Comma-separated list of types. Returned rules match any of the tags (OR operator)
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Actives struct {
		SquidClassCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:ClassCyclomaticComplexity,omitempty"`
		SquidMethodCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:MethodCyclomaticComplexity,omitempty"`
		SquidS1067 []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:S1067,omitempty"`
	} `json:"actives,omitempty"`
	Facets []struct {
		Name   string `json:"name,omitempty"`
		Values []struct {
			Count float64 `json:"count,omitempty"`
			Val   string  `json:"val,omitempty"`
		} `json:"values,omitempty"`
	} `json:"facets,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
	Rules  []struct {
		CleanCodeAttribute         string `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string `json:"cleanCodeAttributeCategory,omitempty"`
		CreatedAt                  string `json:"createdAt,omitempty"`
		DescriptionSections        []struct {
			Content string `json:"content,omitempty"`
			Key     string `json:"key,omitempty"`
			Context struct {
				DisplayName string `json:"displayName,omitempty"`
				Key         string `json:"key,omitempty"`
			} `json:"context,omitempty"`
		} `json:"descriptionSections,omitempty"`
		HtmlDesc string `json:"htmlDesc,omitempty"`
		Impacts  []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		InternalKey string `json:"internalKey,omitempty"`
		IsExternal  bool   `json:"isExternal,omitempty"`
		IsTemplate  bool   `json:"isTemplate,omitempty"`
		Key         string `json:"key,omitempty"`
		Lang        string `json:"lang,omitempty"`
		LangName    string `json:"langName,omitempty"`
		Name        string `json:"name,omitempty"`
		Params      []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			Desc         string `json:"desc,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"params,omitempty"`
		Repo        string   `json:"repo,omitempty"`
		Scope       string   `json:"scope,omitempty"`
		Severity    string   `json:"severity,omitempty"`
		Status      string   `json:"status,omitempty"`
		SysTags     []string `json:"sysTags,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdatedAt   string   `json:"updatedAt,omitempty"`
		HtmlNote    string   `json:"htmlNote,omitempty"`
		MdNote      string   `json:"mdNote,omitempty"`
		NoteLogin   string   `json:"noteLogin,omitempty"`
		TemplateKey string   `json:"templateKey,omitempty"`
	} `json:"rules,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Actives struct {
		SquidClassCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:ClassCyclomaticComplexity,omitempty"`
		SquidMethodCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:MethodCyclomaticComplexity,omitempty"`
		SquidS1067 []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:S1067,omitempty"`
	} `json:"actives,omitempty"`
	Facets []struct {
		Name   string `json:"name,omitempty"`
		Values []struct {
			Count float64 `json:"count,omitempty"`
			Val   string  `json:"val,omitempty"`
		} `json:"values,omitempty"`
	} `json:"facets,omitempty"`
	Rules []struct {
		CleanCodeAttribute         string `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string `json:"cleanCodeAttributeCategory,omitempty"`
		CreatedAt                  string `json:"createdAt,omitempty"`
		DescriptionSections        []struct {
			Content string `json:"content,omitempty"`
			Key     string `json:"key,omitempty"`
			Context struct {
				DisplayName string `json:"displayName,omitempty"`
				Key         string `json:"key,omitempty"`
			} `json:"context,omitempty"`
		} `json:"descriptionSections,omitempty"`
		HtmlDesc string `json:"htmlDesc,omitempty"`
		Impacts  []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		InternalKey string `json:"internalKey,omitempty"`
		IsExternal  bool   `json:"isExternal,omitempty"`
		IsTemplate  bool   `json:"isTemplate,omitempty"`
		Key         string `json:"key,omitempty"`
		Lang        string `json:"lang,omitempty"`
		LangName    string `json:"langName,omitempty"`
		Name        string `json:"name,omitempty"`
		Params      []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			Desc         string `json:"desc,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"params,omitempty"`
		Repo        string   `json:"repo,omitempty"`
		Scope       string   `json:"scope,omitempty"`
		Severity    string   `json:"severity,omitempty"`
		Status      string   `json:"status,omitempty"`
		SysTags     []string `json:"sysTags,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdatedAt   string   `json:"updatedAt,omitempty"`
		HtmlNote    string   `json:"htmlNote,omitempty"`
		MdNote      string   `json:"mdNote,omitempty"`
		NoteLogin   string   `json:"noteLogin,omitempty"`
		TemplateKey string   `json:"templateKey,omitempty"`
	} `json:"rules,omitempty"`
}

// ShowRequest Get detailed information about a rule<br>
type ShowRequest struct {
	Actives string `form:"actives,omitempty"` // Show rule's activations for all profiles ("active rules")
	Key     string `form:"key,omitempty"`     // Rule key
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Actives []struct {
		Inherit string `json:"inherit,omitempty"`
		Params  []struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"params,omitempty"`
		PrioritizedRule bool   `json:"prioritizedRule,omitempty"`
		QProfile        string `json:"qProfile,omitempty"`
		Severity        string `json:"severity,omitempty"`
	} `json:"actives,omitempty"`
	Rule struct {
		CleanCodeAttribute         string `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string `json:"cleanCodeAttributeCategory,omitempty"`
		DefaultRemFnBaseEffort     string `json:"defaultRemFnBaseEffort,omitempty"`
		DefaultRemFnGapMultiplier  string `json:"defaultRemFnGapMultiplier,omitempty"`
		DefaultRemFnType           string `json:"defaultRemFnType,omitempty"`
		DescriptionSections        []struct {
			Content string `json:"content,omitempty"`
			Key     string `json:"key,omitempty"`
			Context struct {
				DisplayName string `json:"displayName,omitempty"`
				Key         string `json:"key,omitempty"`
			} `json:"context,omitempty"`
		} `json:"descriptionSections,omitempty"`
		GapDescription string `json:"gapDescription,omitempty"`
		HtmlDesc       string `json:"htmlDesc,omitempty"`
		Impacts        []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		InternalKey string `json:"internalKey,omitempty"`
		IsExternal  bool   `json:"isExternal,omitempty"`
		Key         string `json:"key,omitempty"`
		Lang        string `json:"lang,omitempty"`
		LangName    string `json:"langName,omitempty"`
		Name        string `json:"name,omitempty"`
		Params      []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			Desc         string `json:"desc,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"params,omitempty"`
		RemFnBaseEffort    string   `json:"remFnBaseEffort,omitempty"`
		RemFnGapMultiplier string   `json:"remFnGapMultiplier,omitempty"`
		RemFnOverloaded    bool     `json:"remFnOverloaded,omitempty"`
		RemFnType          string   `json:"remFnType,omitempty"`
		Repo               string   `json:"repo,omitempty"`
		Scope              string   `json:"scope,omitempty"`
		Severity           string   `json:"severity,omitempty"`
		Status             string   `json:"status,omitempty"`
		SysTags            []string `json:"sysTags,omitempty"`
		Tags               []string `json:"tags,omitempty"`
		Template           bool     `json:"template,omitempty"`
		Type               string   `json:"type,omitempty"`
	} `json:"rule,omitempty"`
}

// TagsRequest List rule tags
type TagsRequest struct {
	Q string `form:"q,omitempty"` // Limit search to tags that contain the supplied string.
}

// TagsResponse is the response for TagsRequest
type TagsResponse struct {
	Tags []string `json:"tags,omitempty"`
}

// UpdateRequest Update an existing rule.<br>Requires the 'Administer Quality Profiles' permission
type UpdateRequest struct {
	Key                        string `form:"key,omitempty"`                           // Key of the rule to update
	MarkdownDescription        string `form:"markdownDescription,omitempty"`           // Rule description (mandatory for custom rule and manual rule) in <a href='/formatting/help'>markdown format</a>
	MarkdownNote               string `form:"markdown_note,omitempty"`                 // Optional note in <a href='/formatting/help'>markdown format</a>. Use empty value to remove current note. Note is not changed if the parameter is not set.
	Name                       string `form:"name,omitempty"`                          // Rule name (mandatory for custom rule)
	Params                     string `form:"params,omitempty"`                        // Parameters as semi-colon list of <key>=<value>, for example 'params=key1=v1;key2=v2' (Only when updating a custom rule)
	RemediationFnBaseEffort    string `form:"remediation_fn_base_effort,omitempty"`    // Base effort of the remediation function of the rule
	RemediationFnType          string `form:"remediation_fn_type,omitempty"`           // Type of the remediation function of the rule
	RemediationFyGapMultiplier string `form:"remediation_fy_gap_multiplier,omitempty"` // Gap multiplier of the remediation function of the rule
	Severity                   string `form:"severity,omitempty"`                      // Rule severity (Only when updating a custom rule)
	Status                     string `form:"status,omitempty"`                        // Rule status (Only when updating a custom rule)
	Tags                       string `form:"tags,omitempty"`                          // Optional comma-separated list of tags to set. Use blank value to remove current tags. Tags are not changed if the parameter is not set.
}

// UpdateResponse is the response for UpdateRequest
type UpdateResponse struct {
	Rule struct {
		CleanCodeAttribute         string `json:"cleanCodeAttribute,omitempty"`
		CleanCodeAttributeCategory string `json:"cleanCodeAttributeCategory,omitempty"`
		CreatedAt                  string `json:"createdAt,omitempty"`
		HtmlDesc                   string `json:"htmlDesc,omitempty"`
		Impacts                    []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		IsExternal bool   `json:"isExternal,omitempty"`
		IsTemplate bool   `json:"isTemplate,omitempty"`
		Key        string `json:"key,omitempty"`
		Lang       string `json:"lang,omitempty"`
		LangName   string `json:"langName,omitempty"`
		MdDesc     string `json:"mdDesc,omitempty"`
		Name       string `json:"name,omitempty"`
		Params     []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			HtmlDesc     string `json:"htmlDesc,omitempty"`
			Key          string `json:"key,omitempty"`
			Type         string `json:"type,omitempty"`
		} `json:"params,omitempty"`
		RemFnOverloaded bool     `json:"remFnOverloaded,omitempty"`
		Repo            string   `json:"repo,omitempty"`
		Scope           string   `json:"scope,omitempty"`
		Severity        string   `json:"severity,omitempty"`
		Status          string   `json:"status,omitempty"`
		SysTags         []string `json:"sysTags,omitempty"`
		Tags            []string `json:"tags,omitempty"`
		TemplateKey     string   `json:"templateKey,omitempty"`
		Type            string   `json:"type,omitempty"`
	} `json:"rule,omitempty"`
}
