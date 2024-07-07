package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/views"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Views service

// AddApplication - Add an existing application to a portfolio.
// Authentication is required for this API endpoint.
// Since 9.3
func (s *Views) AddApplication(ctx context.Context, r views.AddApplicationRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_application", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// AddApplicationBranch - Add a branch of an application selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the application.
// Since 9.3
func (s *Views) AddApplicationBranch(ctx context.Context, r views.AddApplicationBranchRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_application_branch", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// AddPortfolio - Add an existing portfolio to the structure of another portfolio.
// Authentication is required for this API endpoint.
// Since 9.3
func (s *Views) AddPortfolio(ctx context.Context, r views.AddPortfolioRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_portfolio", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// AddProject - Add a project to a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for adding project.
// Since 1.0
// Changelog:
//
//	8.3: Project to which user has `Browse` permission can be used in 'project'
//	7.3: This web service should not be used to add project to an application, api/applications/add_project should be used instead
func (s *Views) AddProject(ctx context.Context, r views.AddProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// AddProjectBranch - Add a branch of a project selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the project.
// Since 9.2
func (s *Views) AddProjectBranch(ctx context.Context, r views.AddProjectBranchRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_project_branch", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Applications - List applications which the user has access to that can be added to a portfolio.
// Authentication is required for this API endpoint
// Since 9.3
func (s *Views) Applications(ctx context.Context, r views.ApplicationsRequest) (*views.ApplicationsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/applications", s.path)
	v := new(views.ApplicationsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Create - Create a new portfolio.
// Requires 'Administer System' permission or 'Create Portfolios' permission,
// Since 1.0
// Changelog:
//
//	9.3: This web service can't create applications. Use 'applications/create' instead
//	9.3: Parameter 'parent' added to create sub-portfolios
//	7.4: Add support of the new permission 'Create Portfolios' permission
//	7.3: This web service should not be used to create an application, api/applications/create should be used instead
//	7.1: The 'visibility' parameter is public
//	2.0: Qualifier field is returned in the response
func (s *Views) Create(ctx context.Context, r views.CreateRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Delete - Delete a portfolio definition.
// Requires 'Administrator' permission on the portfolio.
// Since 1.0
// Changelog:
//
//	9.3: This web service should not be used to delete references to portfolios and applications. Use 'views/remove_portfolio' and 'views/remove_application' instead
//	9.3: This web service can't delete applications. Use 'applications/delete' instead
//	7.3: This web service should not be used to delete an application, api/applications/delete should be used instead
func (s *Views) Delete(ctx context.Context, r views.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// List - List root portfolios.
// Requires authentication. Only portfolios with the admin permission are returned.
// Since 1.0
// Changelog:
//
//	10.0: The applications are removed from the response
//	9.3: Returning applications is now deprecated
//	2.0: Qualifier field is returned in the response
//	2.0: Visibility field is returned in the response
func (s *Views) List(ctx context.Context, r views.ListRequest) (*views.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(views.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Move - Move a portfolio.
// Authentication is required for this API endpoint.
// Since 1.0
func (s *Views) Move(ctx context.Context, r views.MoveRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/move", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// MoveOptions - List possible portfolio destinations.
// Authentication is required for this API endpoint.
// Since 1.0
func (s *Views) MoveOptions(ctx context.Context, r views.MoveOptionsRequest) (*views.MoveOptionsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/move_options", s.path)
	v := new(views.MoveOptionsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Portfolios - List portfolios that can be referenced.
// Authentication is required for this API endpoint.
// Since 9.3
func (s *Views) Portfolios(ctx context.Context, r views.PortfoliosRequest) (*views.PortfoliosResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/portfolios", s.path)
	v := new(views.PortfoliosResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// RemoveApplication - Remove an application from a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 9.3
func (s *Views) RemoveApplication(ctx context.Context, r views.RemoveApplicationRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_application", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveApplicationBranch - Remove a branch of an application selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the application.
// Since 9.3
func (s *Views) RemoveApplicationBranch(ctx context.Context, r views.RemoveApplicationBranchRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_application_branch", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemovePortfolio - Remove a reference to a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 9.3
func (s *Views) RemovePortfolio(ctx context.Context, r views.RemovePortfolioRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_portfolio", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveProject - Remove a project from a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 1.0
// Changelog:
//
//	9.3: The `Browse` permission on a project is no longer required to remove it from a portfolio definition
//	8.3: Project to which user has `Browse` permission can be used in 'project'
//	7.3: This web service should not be used on application, api/applications/remove_project should be used instead
func (s *Views) RemoveProject(ctx context.Context, r views.RemoveProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveProjectBranch - Remove a branch of a project selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the project.
// Since 9.2
func (s *Views) RemoveProjectBranch(ctx context.Context, r views.RemoveProjectBranchRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_project_branch", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetManualMode - Set the projects selection mode of a portfolio on manual selection.
// In order to add project, please use api/view/add_project.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
func (s *Views) SetManualMode(ctx context.Context, r views.SetManualModeRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_manual_mode", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetNoneMode - Set the projects selection mode of a portfolio to none.
// After setting this mode portfolio will not have any projects assigned.
// Requires 'Administrator' permission on the portfolio.
// Since 9.1
func (s *Views) SetNoneMode(ctx context.Context, r views.SetNoneModeRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_none_mode", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetRegexpMode - Set the projects selection mode of a portfolio on regular expression.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
// Changelog:
//
//	9.2: Param 'branch' added
func (s *Views) SetRegexpMode(ctx context.Context, r views.SetRegexpModeRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_regexp_mode", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetRemainingProjectsMode - Set the projects selection mode of a portfolio on unassociated projects in hierarchy.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
// Changelog:
//
//	9.2: Param 'branch' added
func (s *Views) SetRemainingProjectsMode(ctx context.Context, r views.SetRemainingProjectsModeRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_remaining_projects_mode", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetTagsMode - Set the projects selection mode of a portfolio on project tags.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
// Changelog:
//
//	9.2: Param 'branch' added
func (s *Views) SetTagsMode(ctx context.Context, r views.SetTagsModeRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_tags_mode", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Show - Show the details of a portfolio, including its hierarchy and project selection mode.
// Authentication is required for this API endpoint.
// Since 1.0
// Changelog:
//
//	10.1: The deprecated response field 'projects' has been removed. Use 'selectedProjects' instead
//	10.0: This web service no longer shows applications. Use 'api/applications/show' instead
//	9.3: Use of the web service for application is deprecated. api/applications/show should be used
//	9.2: The response field 'projects' for each portfolio is deprecated in favor of the new field 'selectedProjects'
//	7.3: Deleted field of projects is deprecated in the response, replaced by enabled
//	2.0: Qualifier field is returned in the response
func (s *Views) Show(ctx context.Context, r views.ShowRequest) (*views.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(views.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Update - Update a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 1.0
// Changelog:
//
//	9.3: This web service can't update applications. Use 'applications/update' instead
//	7.3: This web service should not be used on application, api/applications/update should be used instead
func (s *Views) Update(ctx context.Context, r views.UpdateRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
