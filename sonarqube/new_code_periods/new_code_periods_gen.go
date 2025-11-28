package new_code_periods

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ListRequest Lists the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/configuring-new-code-calculation#setting-specific-new-code-definition-for-project" target="_blank" rel="noopener noreferrer">new code definition</a> for all branches in a project.<br>Requires the permission to browse the project
type ListRequest struct {
	Project string `url:"project"` // Project key
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	NewCodePeriods []struct {
		BranchKey      string `json:"branchKey,omitempty"`
		EffectiveValue string `json:"effectiveValue,omitempty"`
		ProjectKey     string `json:"projectKey,omitempty"`
		Type           string `json:"type,omitempty"`
		Value          string `json:"value,omitempty"`
		Inherited      bool   `json:"inherited,omitempty"`
	} `json:"newCodePeriods,omitempty"`
}

// SetRequest Updates the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/configuring-new-code-calculation#setting-specific-new-code-definition-for-project" target="_blank" rel="noopener noreferrer">new code definition</a> on different levels:<br><ul><li>Not providing a project key and a branch key will update the default value at global level. Existing projects or branches having a specific new code definition will not be impacted</li><li>Project key must be provided to update the value for a project</li><li>Both project and branch keys must be provided to update the value for a branch</li></ul>Requires one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights on the specified project to change the project setting</li></ul>
type SetRequest struct {
	Branch  string `form:"branch,omitempty"`  // Branch key
	Project string `form:"project,omitempty"` // Project key
	Type    string `form:"type"`              // Type<br/>New code definitions of the following types are allowed:<ul><li>SPECIFIC_ANALYSIS - can be set at branch level only</li><li>PREVIOUS_VERSION - can be set at any level (global, project, branch)</li><li>NUMBER_OF_DAYS - can be set at any level (global, project, branch)</li><li>REFERENCE_BRANCH - can only be set for projects and branches</li></ul>
	Value   string `form:"value,omitempty"`   // Value<br/>For each type, a different value is expected:<ul><li>the uuid of an analysis, when type is SPECIFIC_ANALYSIS</li><li>no value, when type is PREVIOUS_VERSION</li><li>a number between 1 and 90, when type is NUMBER_OF_DAYS</li><li>a string, when type is REFERENCE_BRANCH</li></ul>
}

// ShowRequest Shows the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/configuring-new-code-calculation#setting-specific-new-code-definition-for-project" target="_blank" rel="noopener noreferrer">new code definition</a>.<br> If the component requested doesn't exist or if no new code definition is set for it, a value is inherited from the project or from the global setting.Requires one of the following permissions if a component is specified: <ul><li>'Administer' rights on the specified component</li><li>'Execute analysis' rights on the specified component</li></ul>
type ShowRequest struct {
	Branch  string `url:"branch,omitempty"`  // Branch key
	Project string `url:"project,omitempty"` // Project key
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	BranchKey  string `json:"branchKey,omitempty"`
	Inherited  bool   `json:"inherited,omitempty"`
	ProjectKey string `json:"projectKey,omitempty"`
	Type       string `json:"type,omitempty"`
}

// UnsetRequest Unsets the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/configuring-new-code-calculation#setting-specific-new-code-definition-for-project" target="_blank" rel="noopener noreferrer">new code definition</a> for a branch, project or global. It requires the inherited New Code Definition to be compatible with the Clean as You Code methodology, and one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights for a specified component</li></ul>
type UnsetRequest struct {
	Branch  string `form:"branch,omitempty"`  // Branch key
	Project string `form:"project,omitempty"` // Project key
}
