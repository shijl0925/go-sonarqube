package views

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddApplicationRequest Add an existing application to a portfolio.<br>Authentication is required for this API endpoint.
type AddApplicationRequest struct {
	Application string `json:"application"` // Key of the application to be added
	Portfolio   string `json:"portfolio"`   // Key of the portfolio where the application will be added
}

// AddApplicationBranchRequest Add a branch of an application selected in a portfolio.<br/>Requires 'Administrator' permission on the portfolio and 'Browse' permission for the application.
type AddApplicationBranchRequest struct {
	Application string `json:"application"` // Key of the application
	Branch      string `json:"branch"`      // Key of the branch
	Key         string `json:"key"`         // Key of the portfolio
}

// AddPortfolioRequest Add an existing portfolio to the structure of another portfolio.<br> Authentication is required for this API endpoint.
type AddPortfolioRequest struct {
	Portfolio string `json:"portfolio"` // Key of the portfolio where a reference will be added
	Reference string `json:"reference"` // Key of the portfolio to be added
}

// AddProjectRequest Add a project to a portfolio.<br/>Requires 'Administrator' permission on the portfolio and 'Browse' permission for adding project.
type AddProjectRequest struct {
	Key     string `json:"key"`     // Key of the portfolio
	Project string `json:"project"` // Key of the project
}

// AddProjectBranchRequest Add a branch of a project selected in a portfolio.<br/>Requires 'Administrator' permission on the portfolio and 'Browse' permission for the project.
type AddProjectBranchRequest struct {
	Branch  string `json:"branch"`  // Key of the branch
	Key     string `json:"key"`     // Key of the portfolio
	Project string `json:"project"` // Key of the project
}

// ApplicationsRequest List applications which the user has access to that can be added to a portfolio. <br> Authentication is required for this API endpoint
type ApplicationsRequest struct {
	Portfolio string `url:"portfolio"` // Key of the would-be parent portfolio
}

// ApplicationsResponse is the response for ApplicationsRequest
type ApplicationsResponse struct {
	Applications []struct {
		Disabled bool   `json:"disabled,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"applications,omitempty"`
}

// CreateRequest Create a new portfolio.<br/>Requires 'Administer System' permission or 'Create Portfolios' permission,
type CreateRequest struct {
	Description string `json:"description,omitempty"` // Description for the new portfolio, can be left blank
	Key         string `json:"key,omitempty"`         // Key for the new portfolio. A suitable key will be generated if not provided
	Name        string `json:"name"`                  // Name for the new portfolio
	Parent      string `json:"parent,omitempty"`      // Key of the parent portfolio, when creating a sub portfolio
	Visibility  string `json:"visibility,omitempty"`  // Since 2.0;Whether the created portfolio or application should be visible to everyone, or only specific user/groups.<br/>Only applies to root portfolios. If no visibility is specified, the default visibility will be used.
}

// DeleteRequest Delete a portfolio definition. <br/>Requires 'Administrator' permission on the portfolio.
type DeleteRequest struct {
	Key string `json:"key"` // Portfolio key
}

// ListRequest List root portfolios. <br>Requires authentication. Only portfolios with the admin permission are returned.
type ListRequest struct{}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Views []struct {
		Key        string `json:"key,omitempty"`
		Name       string `json:"name,omitempty"`
		Qualifier  string `json:"qualifier,omitempty"`
		Visibility string `json:"visibility,omitempty"`
	} `json:"views,omitempty"`
}

// MoveRequest Move a portfolio. <br> Authentication is required for this API endpoint.
type MoveRequest struct {
	Destination string `json:"destination"` // Key of the destination portfolio
	Key         string `json:"key"`         // Key of the portfolio to move
}

// MoveOptionsRequest List possible portfolio destinations. <br> Authentication is required for this API endpoint.
type MoveOptionsRequest struct {
	Key string `url:"key"` // Key of the portfolio to move
}

// MoveOptionsResponse is the response for MoveOptionsRequest
type MoveOptionsResponse struct {
	Views []struct {
		Key  string `json:"key,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"views,omitempty"`
}

// PortfoliosRequest List portfolios that can be referenced. <br> Authentication is required for this API endpoint.
type PortfoliosRequest struct {
	Portfolio string `url:"portfolio"` // Key of the would-be parent portfolio
}

// PortfoliosResponse is the response for PortfoliosRequest
type PortfoliosResponse struct {
	Portfolios []struct {
		Disabled bool   `json:"disabled,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"portfolios,omitempty"`
}

// RemoveApplicationRequest Remove an application from a portfolio. <br/>Requires 'Administrator' permission on the portfolio.
type RemoveApplicationRequest struct {
	Application string `json:"application"` // Key of the application to be removed
	Portfolio   string `json:"portfolio"`   // Portfolio key
}

// RemoveApplicationBranchRequest Remove a branch of an application selected in a portfolio.<br/>Requires 'Administrator' permission on the portfolio and 'Browse' permission for the application.
type RemoveApplicationBranchRequest struct {
	Application string `json:"application"` // Key of the project
	Branch      string `json:"branch"`      // Key of the branch
	Key         string `json:"key"`         // Key of the portfolio
}

// RemovePortfolioRequest Remove a reference to a portfolio. <br/>Requires 'Administrator' permission on the portfolio.
type RemovePortfolioRequest struct {
	Portfolio string `json:"portfolio"` // Portfolio key
	Reference string `json:"reference"` // Key of the referenced portfolio to be removed
}

// RemoveProjectRequest Remove a project from a portfolio.<br/>Requires 'Administrator' permission on the portfolio.
type RemoveProjectRequest struct {
	Key     string `json:"key"`     // Key of the portfolio
	Project string `json:"project"` // Key of the project
}

// RemoveProjectBranchRequest Remove a branch of a project selected in a portfolio.<br/>Requires 'Administrator' permission on the portfolio and 'Browse' permission for the project.
type RemoveProjectBranchRequest struct {
	Branch  string `json:"branch"`  // Key of the branch
	Key     string `json:"key"`     // Key of the portfolio
	Project string `json:"project"` // Key of the project
}

// SetManualModeRequest Set the projects selection mode of a portfolio on manual selection.<br/>In order to add project, please use api/view/add_project.<br/>Requires 'Administrator' permission on the portfolio.
type SetManualModeRequest struct {
	Portfolio string `json:"portfolio"` // Key of the portfolio or sub-portfolio to update
}

// SetNoneModeRequest Set the projects selection mode of a portfolio to none.<br/>After setting this mode portfolio will not have any projects assigned.<br/>Requires 'Administrator' permission on the portfolio.
type SetNoneModeRequest struct {
	Portfolio string `json:"portfolio"` // Key of the portfolio or sub-portfolio to update
}

// SetRegexpModeRequest Set the projects selection mode of a portfolio on regular expression.<br/>Requires 'Administrator' permission on the portfolio.
type SetRegexpModeRequest struct {
	Branch    string `json:"branch,omitempty"` // Selects a branch in all matched projects, instead of using their main branches
	Portfolio string `json:"portfolio"`        // Key of the portfolio or sub-portfolio to update
	Regexp    string `json:"regexp"`           // A valid regexp with respect to the JDK's ``java.util.regex.Pattern`` class
}

// SetRemainingProjectsModeRequest Set the projects selection mode of a portfolio on unassociated projects in hierarchy.<br/>Requires 'Administrator' permission on the portfolio.
type SetRemainingProjectsModeRequest struct {
	Branch    string `json:"branch,omitempty"` // Selects a branch in all matched projects, instead of using their main branches
	Portfolio string `json:"portfolio"`        // Key of the portfolio or sub-portfolio to update
}

// SetTagsModeRequest Set the projects selection mode of a portfolio on project tags.<br/>Requires 'Administrator' permission on the portfolio.
type SetTagsModeRequest struct {
	Branch    string `json:"branch,omitempty"` // Selects a branch in all matched projects, instead of using their main branches
	Portfolio string `json:"portfolio"`        // Key of the portfolio or sub-portfolio to update
	Tags      string `json:"tags"`             // Comma-separated list of tags. It's not possible to set nothing.
}

// ShowRequest Show the details of a portfolio, including its hierarchy and project selection mode. <br> Authentication is required for this API endpoint.
type ShowRequest struct {
	Key string `url:"key"` // The key of the portfolio
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Desc         string `json:"desc,omitempty"`
	Key          string `json:"key,omitempty"`
	Name         string `json:"name,omitempty"`
	Qualifier    string `json:"qualifier,omitempty"`
	ReferencedBy []struct {
		Key              string   `json:"key,omitempty"`
		Name             string   `json:"name,omitempty"`
		Qualifier        string   `json:"qualifier,omitempty"`
		ReferencedBy     []string `json:"referencedBy,omitempty"`
		SelectedProjects []struct {
			ProjectKey string `json:"projectKey,omitempty"`
		} `json:"selectedProjects,omitempty"`
		SelectionMode string `json:"selectionMode,omitempty"`
		SubViews      []struct {
			Desc             string   `json:"desc,omitempty"`
			Key              string   `json:"key,omitempty"`
			Name             string   `json:"name,omitempty"`
			OriginalKey      string   `json:"originalKey,omitempty"`
			Qualifier        string   `json:"qualifier,omitempty"`
			Visibility       string   `json:"visibility,omitempty"`
			SelectedBranches []string `json:"selectedBranches,omitempty"`
		} `json:"subViews,omitempty"`
	} `json:"referencedBy,omitempty"`
	SelectedProjects []struct {
		ProjectKey string `json:"projectKey,omitempty"`
	} `json:"selectedProjects,omitempty"`
	SelectionMode string `json:"selectionMode,omitempty"`
	SubViews      []struct {
		Key              string   `json:"key,omitempty"`
		Name             string   `json:"name,omitempty"`
		OriginalKey      string   `json:"originalKey,omitempty"`
		Qualifier        string   `json:"qualifier,omitempty"`
		SelectedBranches []string `json:"selectedBranches,omitempty"`
		Visibility       string   `json:"visibility,omitempty"`
	} `json:"subViews,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

// UpdateRequest Update a portfolio.<br/>Requires 'Administrator' permission on the portfolio.
type UpdateRequest struct {
	Description string `json:"description,omitempty"` // New description for the portfolio
	Key         string `json:"key"`                   // Key of the portfolio to update
	Name        string `json:"name"`                  // New name for the portfolio
}
