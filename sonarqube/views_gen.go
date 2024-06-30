package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/views"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Views service

// AddApplication - Add an existing application to a portfolio.
// Authentication is required for this API endpoint.
// Since 9.3
// Changelog:
func (s *Views) AddApplication(ctx context.Context, r views.AddApplicationRequest) error {
	u := fmt.Sprintf("%s/views/add_application", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddApplicationBranch - Add a branch of an application selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the application.
// Since 9.3
// Changelog:
func (s *Views) AddApplicationBranch(ctx context.Context, r views.AddApplicationBranchRequest) error {
	u := fmt.Sprintf("%s/views/add_application_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddPortfolio - Add an existing portfolio to the structure of another portfolio.
// Authentication is required for this API endpoint.
// Since 9.3
// Changelog:
func (s *Views) AddPortfolio(ctx context.Context, r views.AddPortfolioRequest) error {
	u := fmt.Sprintf("%s/views/add_portfolio", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddProject - Add a project to a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for adding project.
// Since 1.0
// Changelog:
//   8.3: Project to which user has `Browse` permission can be used in 'project'
//   7.3: This web service should not be used to add project to an application, api/applications/add_project should be used instead
func (s *Views) AddProject(ctx context.Context, r views.AddProjectRequest) error {
	u := fmt.Sprintf("%s/views/add_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// AddProjectBranch - Add a branch of a project selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the project.
// Since 9.2
// Changelog:
func (s *Views) AddProjectBranch(ctx context.Context, r views.AddProjectBranchRequest) error {
	u := fmt.Sprintf("%s/views/add_project_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Applications - List applications which the user has access to that can be added to a portfolio.
// Authentication is required for this API endpoint
func (s *Views) Applications(ctx context.Context, r views.ApplicationsRequest) (*views.ApplicationsResponse, error) {
	u := fmt.Sprintf("%s/views/applications", API)
	v := new(views.ApplicationsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Create - Create a new portfolio.
// Requires 'Administer System' permission or 'Create Portfolios' permission,
// Since 1.0
// Changelog:
//   9.3: This web service can't create applications. Use 'applications/create' instead
//   9.3: Parameter 'parent' added to create sub-portfolios
//   7.4: Add support of the new permission 'Create Portfolios' permission
//   7.3: This web service should not be used to create an application, api/applications/create should be used instead
//   7.1: The 'visibility' parameter is public
//   2.0: Qualifier field is returned in the response
func (s *Views) Create(ctx context.Context, r views.CreateRequest) error {
	u := fmt.Sprintf("%s/views/create", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Delete - Delete a portfolio definition.
// Requires 'Administrator' permission on the portfolio.
// Since 1.0
// Changelog:
//   9.3: This web service should not be used to delete references to portfolios and applications. Use 'views/remove_portfolio' and 'views/remove_application' instead
//   9.3: This web service can't delete applications. Use 'applications/delete' instead
//   7.3: This web service should not be used to delete an application, api/applications/delete should be used instead
func (s *Views) Delete(ctx context.Context, r views.DeleteRequest) error {
	u := fmt.Sprintf("%s/views/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// List - List root portfolios.
// Requires authentication. Only portfolios with the admin permission are returned.
func (s *Views) List(ctx context.Context, r views.ListRequest) (*views.ListResponse, error) {
	u := fmt.Sprintf("%s/views/list", API)
	v := new(views.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Move - Move a portfolio.
// Authentication is required for this API endpoint.
// Since 1.0
// Changelog:
func (s *Views) Move(ctx context.Context, r views.MoveRequest) error {
	u := fmt.Sprintf("%s/views/move", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// MoveOptions - List possible portfolio destinations.
// Authentication is required for this API endpoint.
func (s *Views) MoveOptions(ctx context.Context, r views.MoveOptionsRequest) (*views.MoveOptionsResponse, error) {
	u := fmt.Sprintf("%s/views/move_options", API)
	v := new(views.MoveOptionsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Portfolios - List portfolios that can be referenced.
// Authentication is required for this API endpoint.
func (s *Views) Portfolios(ctx context.Context, r views.PortfoliosRequest) (*views.PortfoliosResponse, error) {
	u := fmt.Sprintf("%s/views/portfolios", API)
	v := new(views.PortfoliosResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// RemoveApplication - Remove an application from a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 9.3
// Changelog:
func (s *Views) RemoveApplication(ctx context.Context, r views.RemoveApplicationRequest) error {
	u := fmt.Sprintf("%s/views/remove_application", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveApplicationBranch - Remove a branch of an application selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the application.
// Since 9.3
// Changelog:
func (s *Views) RemoveApplicationBranch(ctx context.Context, r views.RemoveApplicationBranchRequest) error {
	u := fmt.Sprintf("%s/views/remove_application_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemovePortfolio - Remove a reference to a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 9.3
// Changelog:
func (s *Views) RemovePortfolio(ctx context.Context, r views.RemovePortfolioRequest) error {
	u := fmt.Sprintf("%s/views/remove_portfolio", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveProject - Remove a project from a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 1.0
// Changelog:
//   9.3: The `Browse` permission on a project is no longer required to remove it from a portfolio definition
//   8.3: Project to which user has `Browse` permission can be used in 'project'
//   7.3: This web service should not be used on application, api/applications/remove_project should be used instead
func (s *Views) RemoveProject(ctx context.Context, r views.RemoveProjectRequest) error {
	u := fmt.Sprintf("%s/views/remove_project", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// RemoveProjectBranch - Remove a branch of a project selected in a portfolio.
// Requires 'Administrator' permission on the portfolio and 'Browse' permission for the project.
// Since 9.2
// Changelog:
func (s *Views) RemoveProjectBranch(ctx context.Context, r views.RemoveProjectBranchRequest) error {
	u := fmt.Sprintf("%s/views/remove_project_branch", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetManualMode - Set the projects selection mode of a portfolio on manual selection.
// In order to add project, please use api/view/add_project.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
// Changelog:
func (s *Views) SetManualMode(ctx context.Context, r views.SetManualModeRequest) error {
	u := fmt.Sprintf("%s/views/set_manual_mode", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetNoneMode - Set the projects selection mode of a portfolio to none.
// After setting this mode portfolio will not have any projects assigned.
// Requires 'Administrator' permission on the portfolio.
// Since 9.1
// Changelog:
func (s *Views) SetNoneMode(ctx context.Context, r views.SetNoneModeRequest) error {
	u := fmt.Sprintf("%s/views/set_none_mode", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetRegexpMode - Set the projects selection mode of a portfolio on regular expression.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
// Changelog:
//   9.2: Param 'branch' added
func (s *Views) SetRegexpMode(ctx context.Context, r views.SetRegexpModeRequest) error {
	u := fmt.Sprintf("%s/views/set_regexp_mode", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetRemainingProjectsMode - Set the projects selection mode of a portfolio on unassociated projects in hierarchy.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
// Changelog:
//   9.2: Param 'branch' added
func (s *Views) SetRemainingProjectsMode(ctx context.Context, r views.SetRemainingProjectsModeRequest) error {
	u := fmt.Sprintf("%s/views/set_remaining_projects_mode", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetTagsMode - Set the projects selection mode of a portfolio on project tags.
// Requires 'Administrator' permission on the portfolio.
// Since 7.4
// Changelog:
//   9.2: Param 'branch' added
func (s *Views) SetTagsMode(ctx context.Context, r views.SetTagsModeRequest) error {
	u := fmt.Sprintf("%s/views/set_tags_mode", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Show - Show the details of a portfolio, including its hierarchy and project selection mode.
// Authentication is required for this API endpoint.
func (s *Views) Show(ctx context.Context, r views.ShowRequest) (*views.ShowResponse, error) {
	u := fmt.Sprintf("%s/views/show", API)
	v := new(views.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Update - Update a portfolio.
// Requires 'Administrator' permission on the portfolio.
// Since 1.0
// Changelog:
//   9.3: This web service can't update applications. Use 'applications/update' instead
//   7.3: This web service should not be used on application, api/applications/update should be used instead
func (s *Views) Update(ctx context.Context, r views.UpdateRequest) error {
	u := fmt.Sprintf("%s/views/update", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
