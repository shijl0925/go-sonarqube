package applications

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddProjectRequest Add a project to an application.<br/>Requires 'Administrator' permission on the application
type AddProjectRequest struct {
	Application string `json:"application"` // Key of the application
	Project     string `json:"project"`     // Key of the project
}

// CreateRequest Create a new application.<br/>Requires 'Administer System' permission or 'Create Applications' permission
type CreateRequest struct {
	Description string `json:"description,omitempty"` // Application description
	Key         string `json:"key,omitempty"`         // Application key. A suitable key will be generated if not provided
	Name        string `json:"name"`                  // Application name
	Visibility  string `json:"visibility,omitempty"`  // Whether the created application should be visible to everyone, or only specific user/groups.<br/>If no visibility is specified, the default visibility will be used.
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Application struct {
		Description string   `json:"description,omitempty"`
		Key         string   `json:"key,omitempty"`
		Name        string   `json:"name,omitempty"`
		Projects    []string `json:"projects,omitempty"`
		Visibility  string   `json:"visibility,omitempty"`
	} `json:"application,omitempty"`
}

// CreateBranchRequest Create a new branch on a given application.<br/>Requires 'Administrator' permission on the application and 'Browse' permission on its child projects
type CreateBranchRequest struct {
	Application   string `json:"application"`   // Application key
	Branch        string `json:"branch"`        // Branch name
	Project       string `json:"project"`       // Project keys. To set several values, the parameter must be called once for each value.
	ProjectBranch string `json:"projectBranch"` // Project branches. To set main branch, provide an empty value. To set several values, the parameter must be called once for each value.
}

// DeleteRequest Delete an application definition.<br/>Requires 'Administrator' permission on the application
type DeleteRequest struct {
	Application string `json:"application"` // Application key
}

// DeleteBranchRequest Delete a branch on a given application.<br/>Requires 'Administrator' permission on the application
type DeleteBranchRequest struct {
	Application string `json:"application"` // Application key
	Branch      string `json:"branch"`      // Branch name
}

// RemoveProjectRequest Remove a project from an application<br/>Requires 'Administrator' permission on the application
type RemoveProjectRequest struct {
	Application string `json:"application"` // Key of the application
	Project     string `json:"project"`     // Key of the project
}

// SetTagsRequest Set tags on a application.<br>Requires the following permission: 'Administer' rights on the specified application
type SetTagsRequest struct {
	Application string `json:"application"` // Application key
	Tags        string `json:"tags"`        // Comma-separated list of tags
}

// ShowRequest Returns an application and its associated projects. <br> Requires the 'Browse' permission on the application and on its child projects.
type ShowRequest struct {
	Application string `url:"application"`      // Application key
	Branch      string `url:"branch,omitempty"` // Branch name
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Application struct {
		Branch   string `json:"branch,omitempty"`
		Branches []struct {
			IsMain bool   `json:"isMain,omitempty"`
			Name   string `json:"name,omitempty"`
		} `json:"branches,omitempty"`
		IsMain   bool   `json:"isMain,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
		Projects []struct {
			Branch   string `json:"branch,omitempty"`
			Enabled  bool   `json:"enabled,omitempty"`
			IsMain   bool   `json:"isMain,omitempty"`
			Key      string `json:"key,omitempty"`
			Name     string `json:"name,omitempty"`
			Selected bool   `json:"selected,omitempty"`
		} `json:"projects,omitempty"`
		Tags       []string `json:"tags,omitempty"`
		Visibility string   `json:"visibility,omitempty"`
	} `json:"application,omitempty"`
}

// UpdateRequest Update an application.<br/>Requires 'Administrator' permission on the application
type UpdateRequest struct {
	Application string `json:"application"`           // Application key
	Description string `json:"description,omitempty"` // New description for the application
	Name        string `json:"name"`                  // New name for the application
}

// UpdateBranchRequest Update a branch on a given application.<br/>Requires 'Administrator' permission on the application and 'Browse' permission on its child projects
type UpdateBranchRequest struct {
	Application   string `json:"application"`   // Application key
	Branch        string `json:"branch"`        // Branch name
	Name          string `json:"name"`          // New branch name
	Project       string `json:"project"`       // Project keys. To set several values, the parameter must be called once for each value.
	ProjectBranch string `json:"projectBranch"` // Project branches. To set main branch, provide an empty value. To set several values, the parameter must be called once for each value.
}
