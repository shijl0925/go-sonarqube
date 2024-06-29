package qualityprofiles

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ActivateRuleRequest Activate a rule on a Quality Profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
type ActivateRuleRequest struct {
	Key             string `form:"key,omitempty"`             // Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>
	Params          string `form:"params,omitempty"`          // Parameters as semi-colon list of <code>key=value</code>. Ignored if parameter reset is true.
	PrioritizedRule string `form:"prioritizedRule,omitempty"` // Mark activated rule as prioritized, so all corresponding Issues will have to be fixed.
	Reset           string `form:"reset,omitempty"`           // Reset severity and parameters of activated rule. Set the values defined on parent profile or from rule default values.
	Rule            string `form:"rule,omitempty"`            // Rule key
	Severity        string `form:"severity,omitempty"`        // Severity. Ignored if parameter reset is true.
}

// ActivateRulesRequest Bulk-activate rules on one quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
type ActivateRulesRequest struct {
	Activation                   string `form:"activation,omitempty"`                   // Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.
	ActiveSeverities             string `form:"active_severities,omitempty"`            // Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.
	Asc                          string `form:"asc,omitempty"`                          // Ascending sort
	AvailableSince               string `form:"available_since,omitempty"`              // Filters rules added since date. Format is yyyy-MM-dd
	CleanCodeAttributeCategories string `form:"cleanCodeAttributeCategories,omitempty"` // Comma-separated list of Clean Code Attribute Categories
	Cwe                          string `form:"cwe,omitempty"`                          // Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE.
	ImpactSeverities             string `form:"impactSeverities,omitempty"`             // Comma-separated list of Software Quality Severities
	ImpactSoftwareQualities      string `form:"impactSoftwareQualities,omitempty"`      // Comma-separated list of Software Qualities
	Inheritance                  string `form:"inheritance,omitempty"`                  // Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.
	IsTemplate                   string `form:"is_template,omitempty"`                  // Filter template rules
	Languages                    string `form:"languages,omitempty"`                    // Comma-separated list of languages
	OwaspTop10                   string `form:"owaspTop10,omitempty"`                   // Comma-separated list of OWASP Top 10 2017 lowercase categories.
	OwaspTop102021               string `form:"owaspTop10-2021,omitempty"`              // Comma-separated list of OWASP Top 10 2021 lowercase categories.
	PrioritizedRule              string `form:"prioritizedRule,omitempty"`              // Mark activated rules as prioritized, so all corresponding Issues will have to be fixed.
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
	TargetKey                    string `form:"targetKey,omitempty"`                    // Quality Profile key on which the rule activation is done. To retrieve a quality profile key please see <code>api/qualityprofiles/search</code>
	TargetSeverity               string `form:"targetSeverity,omitempty"`               // Severity to set on the activated rules
	TemplateKey                  string `form:"template_key,omitempty"`                 // Key of the template rule to filter on. Used to search for the custom rules based on this template.
	Types                        string `form:"types,omitempty"`                        // Comma-separated list of types. Returned rules match any of the tags (OR operator)
}

// AddProjectRequest Associate a project with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Administer right on the specified project</li></ul>
type AddProjectRequest struct {
	Language       string `form:"language,omitempty"`       // Quality profile language.
	Project        string `form:"project,omitempty"`        // Project key
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name.
}

// BackupRequest Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore.
type BackupRequest struct {
	Language       string `form:"language,omitempty"`       // Quality profile language.
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name.
}

// BackupResponse is the response for BackupRequest
type BackupResponse string

// ChangeParentRequest Change a quality profile's parent.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
type ChangeParentRequest struct {
	Language             string `form:"language,omitempty"`             // Quality profile language.
	ParentQualityProfile string `form:"parentQualityProfile,omitempty"` // New parent profile name. <br> If no profile is provided, the inheritance link with current parent profile (if any) is broken, which deactivates all rules which come from the parent and are not overridden.
	QualityProfile       string `form:"qualityProfile,omitempty"`       // Quality profile name.
}

// ChangelogRequest Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first).
type ChangelogRequest struct {
	Language       string `form:"language,omitempty"`       // Quality profile language.
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name.
	Since          string `form:"since,omitempty"`          // Start date for the changelog (inclusive). <br>Either a date (server timezone) or datetime can be provided.
	To             string `form:"to,omitempty"`             // End date for the changelog (exclusive, strictly before). <br>Either a date (server timezone) or datetime can be provided.
}

// ChangelogResponse is the response for ChangelogRequest
type ChangelogResponse struct {
	Events []struct {
		Action                     string `json:"action,omitempty"`
		AuthorLogin                string `json:"authorLogin,omitempty"`
		AuthorName                 string `json:"authorName,omitempty"`
		CleanCodeAttributeCategory string `json:"cleanCodeAttributeCategory,omitempty"`
		Date                       string `json:"date,omitempty"`
		Impacts                    []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		Params struct {
			Severity string `json:"severity,omitempty"`
		} `json:"params,omitempty"`
		RuleKey          string `json:"ruleKey,omitempty"`
		RuleName         string `json:"ruleName,omitempty"`
		SonarQubeVersion string `json:"sonarQubeVersion,omitempty"`
	} `json:"events,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from ChangelogResponse
func (r *ChangelogResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// ChangelogResponseAll is the collection for ChangelogRequest
type ChangelogResponseAll struct {
	Events []struct {
		Action                     string `json:"action,omitempty"`
		AuthorLogin                string `json:"authorLogin,omitempty"`
		AuthorName                 string `json:"authorName,omitempty"`
		CleanCodeAttributeCategory string `json:"cleanCodeAttributeCategory,omitempty"`
		Date                       string `json:"date,omitempty"`
		Impacts                    []struct {
			Severity        string `json:"severity,omitempty"`
			SoftwareQuality string `json:"softwareQuality,omitempty"`
		} `json:"impacts,omitempty"`
		Params struct {
			Severity string `json:"severity,omitempty"`
		} `json:"params,omitempty"`
		RuleKey          string `json:"ruleKey,omitempty"`
		RuleName         string `json:"ruleName,omitempty"`
		SonarQubeVersion string `json:"sonarQubeVersion,omitempty"`
	} `json:"events,omitempty"`
}

// CopyRequest Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
type CopyRequest struct {
	FromKey string `form:"fromKey,omitempty"` // Quality profile key
	ToName  string `form:"toName,omitempty"`  // Name for the new quality profile.
}

// CopyResponse is the response for CopyRequest
type CopyResponse struct {
	IsDefault   bool   `json:"isDefault,omitempty"`
	IsInherited bool   `json:"isInherited,omitempty"`
	Key         string `json:"key,omitempty"`
	Language    string `json:"language,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentKey   string `json:"parentKey,omitempty"`
}

// CreateRequest Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission.
type CreateRequest struct {
	Language string `form:"language,omitempty"` // Quality profile language
	Name     string `form:"name,omitempty"`     // Quality profile name
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Profile struct {
		IsDefault    bool   `json:"isDefault,omitempty"`
		IsInherited  bool   `json:"isInherited,omitempty"`
		Key          string `json:"key,omitempty"`
		Language     string `json:"language,omitempty"`
		LanguageName string `json:"languageName,omitempty"`
		Name         string `json:"name,omitempty"`
	} `json:"profile,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// DeactivateRuleRequest Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
type DeactivateRuleRequest struct {
	Key  string `form:"key,omitempty"`  // Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>
	Rule string `form:"rule,omitempty"` // Rule key
}

// DeactivateRulesRequest Bulk deactivate rules on Quality profiles.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
type DeactivateRulesRequest struct {
	Activation                   string `form:"activation,omitempty"`                   // Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.
	ActiveSeverities             string `form:"active_severities,omitempty"`            // Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.
	Asc                          string `form:"asc,omitempty"`                          // Ascending sort
	AvailableSince               string `form:"available_since,omitempty"`              // Filters rules added since date. Format is yyyy-MM-dd
	CleanCodeAttributeCategories string `form:"cleanCodeAttributeCategories,omitempty"` // Comma-separated list of Clean Code Attribute Categories
	Cwe                          string `form:"cwe,omitempty"`                          // Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE.
	ImpactSeverities             string `form:"impactSeverities,omitempty"`             // Comma-separated list of Software Quality Severities
	ImpactSoftwareQualities      string `form:"impactSoftwareQualities,omitempty"`      // Comma-separated list of Software Qualities
	Inheritance                  string `form:"inheritance,omitempty"`                  // Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.
	IsTemplate                   string `form:"is_template,omitempty"`                  // Filter template rules
	Languages                    string `form:"languages,omitempty"`                    // Comma-separated list of languages
	OwaspTop10                   string `form:"owaspTop10,omitempty"`                   // Comma-separated list of OWASP Top 10 2017 lowercase categories.
	OwaspTop102021               string `form:"owaspTop10-2021,omitempty"`              // Comma-separated list of OWASP Top 10 2021 lowercase categories.
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
	TargetKey                    string `form:"targetKey,omitempty"`                    // Quality Profile key on which the rule deactivation is done. To retrieve a profile key please see <code>api/qualityprofiles/search</code>
	TemplateKey                  string `form:"template_key,omitempty"`                 // Key of the template rule to filter on. Used to search for the custom rules based on this template.
	Types                        string `form:"types,omitempty"`                        // Comma-separated list of types. Returned rules match any of the tags (OR operator)
}

// DeleteRequest Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
type DeleteRequest struct {
	Language       string `form:"language,omitempty"`       // Quality profile language.
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name.
}

// ExportRequest Export a quality profile.
type ExportRequest struct {
	ExporterKey    string `form:"exporterKey,omitempty"`    // Output format. If left empty, the same format as api/qualityprofiles/backup is used. Possible values are described by api/qualityprofiles/exporters.
	Language       string `form:"language,omitempty"`       // Quality profile language
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name to export. If left empty, the default profile for the language is exported.
}

// ExportResponse is the response for ExportRequest
type ExportResponse string

// ExportersRequest Lists available profile export formats.
type ExportersRequest struct{}

// ExportersResponse is the response for ExportersRequest
type ExportersResponse struct {
	Exporters []struct {
		Key       string   `json:"key,omitempty"`
		Languages []string `json:"languages,omitempty"`
		Name      string   `json:"name,omitempty"`
	} `json:"exporters,omitempty"`
}

// ImportersRequest List supported importers.
type ImportersRequest struct{}

// ImportersResponse is the response for ImportersRequest
type ImportersResponse struct {
	Importers []struct {
		Key       string   `json:"key,omitempty"`
		Languages []string `json:"languages,omitempty"`
		Name      string   `json:"name,omitempty"`
	} `json:"importers,omitempty"`
}

// InheritanceRequest Show a quality profile's ancestors and children.
type InheritanceRequest struct {
	Language       string `form:"language,omitempty"`       // Quality profile language.
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name.
}

// InheritanceResponse is the response for InheritanceRequest
type InheritanceResponse struct {
	Ancestors []struct {
		ActiveRuleCount float64 `json:"activeRuleCount,omitempty"`
		IsBuiltIn       bool    `json:"isBuiltIn,omitempty"`
		Key             string  `json:"key,omitempty"`
		Name            string  `json:"name,omitempty"`
		Parent          string  `json:"parent,omitempty"`
	} `json:"ancestors,omitempty"`
	Children []struct {
		ActiveRuleCount     float64 `json:"activeRuleCount,omitempty"`
		IsBuiltIn           bool    `json:"isBuiltIn,omitempty"`
		Key                 string  `json:"key,omitempty"`
		Name                string  `json:"name,omitempty"`
		OverridingRuleCount float64 `json:"overridingRuleCount,omitempty"`
	} `json:"children,omitempty"`
	Profile struct {
		ActiveRuleCount     float64 `json:"activeRuleCount,omitempty"`
		IsBuiltIn           bool    `json:"isBuiltIn,omitempty"`
		Key                 string  `json:"key,omitempty"`
		Name                string  `json:"name,omitempty"`
		OverridingRuleCount float64 `json:"overridingRuleCount,omitempty"`
		Parent              string  `json:"parent,omitempty"`
	} `json:"profile,omitempty"`
}

// ProjectsRequest List projects with their association status regarding a quality profile. <br/>Only projects explicitly bound to the profile are returned, those associated with the profile because it is the default one are not. <br/>See api/qualityprofiles/search in order to get the Quality Profile Key.
type ProjectsRequest struct {
	Key      string `form:"key,omitempty"`      // Quality profile key
	Q        string `form:"q,omitempty"`        // Limit search to projects that contain the supplied string.
	Selected string `form:"selected,omitempty"` // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// ProjectsResponse is the response for ProjectsRequest
type ProjectsResponse struct {
	Paging  paging.Paging `json:"paging,omitempty"`
	Results []struct {
		Id       string `json:"id,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"results,omitempty"`
}

// GetPaging extracts the paging from ProjectsResponse
func (r *ProjectsResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// ProjectsResponseAll is the collection for ProjectsRequest
type ProjectsResponseAll struct {
	Results []struct {
		Id       string `json:"id,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"results,omitempty"`
}

// RemoveProjectRequest Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
type RemoveProjectRequest struct {
	Language       string `form:"language,omitempty"`       // Quality profile language.
	Project        string `form:"project,omitempty"`        // Project key
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name.
}

// RenameRequest Rename a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
type RenameRequest struct {
	Key  string `form:"key,omitempty"`  // Quality profile key
	Name string `form:"name,omitempty"` // New quality profile name
}

// RestoreRequest Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
type RestoreRequest struct {
	Backup string `form:"backup,omitempty"` // A profile backup file in XML format, as generated by api/qualityprofiles/backup or the former api/profiles/backup.
}

// SearchRequest Search quality profiles
type SearchRequest struct {
	Defaults       string `form:"defaults,omitempty"`       // If set to true, return only the quality profiles marked as default for each language
	Language       string `form:"language,omitempty"`       // Language key. If provided, only profiles for the given language are returned.
	Project        string `form:"project,omitempty"`        // Project key
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Actions struct {
		Create bool `json:"create,omitempty"`
	} `json:"actions,omitempty"`
	Profiles []struct {
		Actions struct {
			AssociateProjects bool `json:"associateProjects,omitempty"`
			Copy              bool `json:"copy,omitempty"`
			Delete            bool `json:"delete,omitempty"`
			Edit              bool `json:"edit,omitempty"`
			SetAsDefault      bool `json:"setAsDefault,omitempty"`
		} `json:"actions,omitempty"`
		ActiveDeprecatedRuleCount float64 `json:"activeDeprecatedRuleCount,omitempty"`
		ActiveRuleCount           float64 `json:"activeRuleCount,omitempty"`
		IsBuiltIn                 bool    `json:"isBuiltIn,omitempty"`
		IsDefault                 bool    `json:"isDefault,omitempty"`
		IsInherited               bool    `json:"isInherited,omitempty"`
		Key                       string  `json:"key,omitempty"`
		Language                  string  `json:"language,omitempty"`
		LanguageName              string  `json:"languageName,omitempty"`
		LastUsed                  string  `json:"lastUsed,omitempty"`
		Name                      string  `json:"name,omitempty"`
		RuleUpdatedAt             string  `json:"ruleUpdatedAt,omitempty"`
		ParentKey                 string  `json:"parentKey,omitempty"`
		ParentName                string  `json:"parentName,omitempty"`
		ProjectCount              float64 `json:"projectCount,omitempty"`
		UserUpdatedAt             string  `json:"userUpdatedAt,omitempty"`
	} `json:"profiles,omitempty"`
}

// SetDefaultRequest Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
type SetDefaultRequest struct {
	Language       string `form:"language,omitempty"`       // Quality profile language.
	QualityProfile string `form:"qualityProfile,omitempty"` // Quality profile name.
}
