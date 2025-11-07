package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/new_code_periods"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type NewCodePeriods service

// List - Lists the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/setting-up-clean-as-you-code/#setting-a-new-code-definition" target="_blank" rel="noopener noreferrer">new code definition</a> for all branches in a project.
// Requires the permission to browse the project
// Since 8.0
func (s *NewCodePeriods) List(ctx context.Context, r new_code_periods.ListRequest) (*new_code_periods.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(new_code_periods.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Set - Updates the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/setting-up-clean-as-you-code/#setting-a-new-code-definition" target="_blank" rel="noopener noreferrer">new code definition</a> on different levels:
//
//  * Not providing a project key and a branch key will update the default value at global level. Existing projects or branches having a specific new code definition will not be impacted
//  * Project key must be provided to update the value for a project
//  * Both project and branch keys must be provided to update the value for a branch
// Requires one of the following permissions:
//  * 'Administer System' to change the global setting
//  * 'Administer' rights on the specified project to change the project setting
//
// Since 8.0
func (s *NewCodePeriods) Set(ctx context.Context, r new_code_periods.SetRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Show - Shows the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/setting-up-clean-as-you-code/#setting-a-new-code-definition" target="_blank" rel="noopener noreferrer">new code definition</a>.
// If the component requested doesn't exist or if no new code definition is set for it, a value is inherited from the project or from the global setting.Requires one of the following permissions if a component is specified:
//  * 'Administer' rights on the specified component
//  * 'Execute analysis' rights on the specified component
//
// Since 8.0
func (s *NewCodePeriods) Show(ctx context.Context, r new_code_periods.ShowRequest) (*new_code_periods.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(new_code_periods.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Unset - Unsets the <a href="https://sonar-documentations-preview.netlify.app/sonarqube-server/latest/project-administration/setting-up-clean-as-you-code/#setting-a-new-code-definition" target="_blank" rel="noopener noreferrer">new code definition</a> for a branch, project or global. It requires the inherited New Code Definition to be compatible with the Clean as You Code methodology, and one of the following permissions:
//  * 'Administer System' to change the global setting
//  * 'Administer' rights for a specified component
//
// Since 8.0
func (s *NewCodePeriods) Unset(ctx context.Context, r new_code_periods.UnsetRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/unset", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
