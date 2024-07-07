package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/alm_settings"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AlmSettings service

// CountBinding - Count number of project bound to an DevOps Platform setting.
// Requires the 'Administer System' permission
// Since 8.1
func (s *AlmSettings) CountBinding(ctx context.Context, r alm_settings.CountBindingRequest) (*alm_settings.CountBindingResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/count_binding", s.path)
	v := new(alm_settings.CountBindingResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// CreateAzure - Create Azure instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	8.6: Parameter 'URL' was added
func (s *AlmSettings) CreateAzure(ctx context.Context, r alm_settings.CreateAzureRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/create_azure", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateBitbucket - Create Bitbucket instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
func (s *AlmSettings) CreateBitbucket(ctx context.Context, r alm_settings.CreateBitbucketRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/create_bitbucket", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateBitbucketcloud - Configure a new instance of Bitbucket Cloud.
// Requires the 'Administer System' permission
// Since 8.7
func (s *AlmSettings) CreateBitbucketcloud(ctx context.Context, r alm_settings.CreateBitbucketcloudRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/create_bitbucketcloud", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateGithub - Create GitHub instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	9.7: Optional parameter 'webhookSecret' was added
func (s *AlmSettings) CreateGithub(ctx context.Context, r alm_settings.CreateGithubRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/create_github", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateGitlab - Create GitLab instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	8.2: Parameter 'URL' was added
func (s *AlmSettings) CreateGitlab(ctx context.Context, r alm_settings.CreateGitlabRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/create_gitlab", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Delete - Delete an DevOps Platform Setting.
// Requires the 'Administer System' permission
// Since 8.1
func (s *AlmSettings) Delete(ctx context.Context, r alm_settings.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteBinding - Delete the DevOps Platform binding of a project.
// Requires the 'Administer' permission on the project
// Since 8.1
func (s *AlmSettings) DeleteBinding(ctx context.Context, r alm_settings.DeleteBindingRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete_binding", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetBinding - Get DevOps Platform binding of a given project.
// Requires the 'Browse' permission on the project
// Since 8.1
// Changelog:
//
//	10.1: Permission needed changed from 'Administer' to 'Browse'
//	8.7: Azure binding now contains a monorepo flag for monorepo feature in Enterprise Edition and above
//	8.6: Azure binding now contains the project and repository names
func (s *AlmSettings) GetBinding(ctx context.Context, r alm_settings.GetBindingRequest) (*alm_settings.GetBindingResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/get_binding", s.path)
	v := new(alm_settings.GetBindingResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// List - List DevOps Platform setting available for a given project, sorted by DevOps Platform key
// Requires the 'Administer project' permission if the 'project' parameter is provided, requires the 'Create Projects' permission otherwise.
// Since 8.1
// Changelog:
//
//	8.6: Field 'URL' added for Azure definitions
//	8.3: Permission needed changed to 'Administer project' or 'Create Projects'
//	8.2: Permission needed changed from 'Administer project' to 'Create Projects'
func (s *AlmSettings) List(ctx context.Context, r alm_settings.ListRequest) (*alm_settings.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(alm_settings.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// ListDefinitions - List DevOps Platform Settings, sorted by created date.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	8.7: Fields 'personalAccessToken', 'privateKey', and 'clientSecret' are no longer returned
//	8.6: Field 'URL' added for Azure definitions
//	8.2: Field 'URL' added for GitLab definitions
func (s *AlmSettings) ListDefinitions(ctx context.Context, r alm_settings.ListDefinitionsRequest) (*alm_settings.ListDefinitionsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list_definitions", s.path)
	v := new(alm_settings.ListDefinitionsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// SetAzureBinding - Bind a Azure DevOps instance to a project.
// If the project was already bound to a previous Azure DevOps instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
func (s *AlmSettings) SetAzureBinding(ctx context.Context, r alm_settings.SetAzureBindingRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_azure_binding", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetBitbucketBinding - Bind a Bitbucket instance to a project.
// If the project was already bound to a previous Bitbucket instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
func (s *AlmSettings) SetBitbucketBinding(ctx context.Context, r alm_settings.SetBitbucketBindingRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_bitbucket_binding", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetBitbucketcloudBinding - Bind a Bitbucket Cloud setting to a project.
// If the project was already bound to a different Bitbucket Cloud setting, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.7
func (s *AlmSettings) SetBitbucketcloudBinding(ctx context.Context, r alm_settings.SetBitbucketcloudBindingRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_bitbucketcloud_binding", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetGithubBinding - Bind a GitHub instance to a project.
// If the project was already bound to a previous GitHub instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
// Changelog:
//
//	8.3: Add 'summaryCommentEnabled' param to enable/disable of putting analysis summary in a conversation tab of GitHub
func (s *AlmSettings) SetGithubBinding(ctx context.Context, r alm_settings.SetGithubBindingRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_github_binding", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetGitlabBinding - Bind a GitLab instance to a project.
// If the project was already bound to a previous Gitlab instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
// Changelog:
//
//	8.2: Parameter 'repository' was added
func (s *AlmSettings) SetGitlabBinding(ctx context.Context, r alm_settings.SetGitlabBindingRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_gitlab_binding", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateAzure - Update Azure instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	8.7: Parameter 'personalAccessToken' is no longer required
//	8.6: Parameter 'URL' was added
func (s *AlmSettings) UpdateAzure(ctx context.Context, r alm_settings.UpdateAzureRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update_azure", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateBitbucket - Update Bitbucket instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	8.7: Parameter 'personalAccessToken' is no longer required
func (s *AlmSettings) UpdateBitbucket(ctx context.Context, r alm_settings.UpdateBitbucketRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update_bitbucket", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateBitbucketcloud - Update Bitbucket Cloud Setting.
// Requires the 'Administer System' permission
// Since 8.7
func (s *AlmSettings) UpdateBitbucketcloud(ctx context.Context, r alm_settings.UpdateBitbucketcloudRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update_bitbucketcloud", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateGithub - Update GitHub instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	9.7: Optional parameter 'webhookSecret' was added
//	8.7: Parameter 'privateKey' is no longer required
//	8.7: Parameter 'clientSecret' is no longer required
func (s *AlmSettings) UpdateGithub(ctx context.Context, r alm_settings.UpdateGithubRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update_github", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateGitlab - Update GitLab instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//
//	8.7: Parameter 'personalAccessToken' is no longer required
//	8.2: Parameter 'URL' was added
func (s *AlmSettings) UpdateGitlab(ctx context.Context, r alm_settings.UpdateGitlabRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update_gitlab", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Validate - Validate an DevOps Platform Setting by checking connectivity and permissions
// Requires the 'Administer System' permission
// Since 8.6
func (s *AlmSettings) Validate(ctx context.Context, r alm_settings.ValidateRequest) (*alm_settings.ValidateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/validate", s.path)
	v := new(alm_settings.ValidateResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
