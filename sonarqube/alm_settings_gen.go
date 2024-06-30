package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/alm_settings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AlmSettings service

// CountBinding - Count number of project bound to an DevOps Platform setting.
// Requires the 'Administer System' permission
func (s *AlmSettings) CountBinding(ctx context.Context, r alm_settings.CountBindingRequest) (*alm_settings.CountBindingResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/count_binding", API)
	v := new(alm_settings.CountBindingResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// CreateAzure - Create Azure instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//   8.6: Parameter 'URL' was added
func (s *AlmSettings) CreateAzure(ctx context.Context, r alm_settings.CreateAzureRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_azure", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// CreateBitbucket - Create Bitbucket instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
func (s *AlmSettings) CreateBitbucket(ctx context.Context, r alm_settings.CreateBitbucketRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_bitbucket", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// CreateBitbucketcloud - Configure a new instance of Bitbucket Cloud.
// Requires the 'Administer System' permission
// Since 8.7
// Changelog:
func (s *AlmSettings) CreateBitbucketcloud(ctx context.Context, r alm_settings.CreateBitbucketcloudRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_bitbucketcloud", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// CreateGithub - Create GitHub instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//   9.7: Optional parameter 'webhookSecret' was added
func (s *AlmSettings) CreateGithub(ctx context.Context, r alm_settings.CreateGithubRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_github", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// CreateGitlab - Create GitLab instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//   8.2: Parameter 'URL' was added
func (s *AlmSettings) CreateGitlab(ctx context.Context, r alm_settings.CreateGitlabRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_gitlab", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Delete - Delete an DevOps Platform Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
func (s *AlmSettings) Delete(ctx context.Context, r alm_settings.DeleteRequest) error {
	u := fmt.Sprintf("%s/alm_settings/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// DeleteBinding - Delete the DevOps Platform binding of a project.
// Requires the 'Administer' permission on the project
// Since 8.1
// Changelog:
func (s *AlmSettings) DeleteBinding(ctx context.Context, r alm_settings.DeleteBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/delete_binding", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// GetBinding - Get DevOps Platform binding of a given project.
// Requires the 'Browse' permission on the project
func (s *AlmSettings) GetBinding(ctx context.Context, r alm_settings.GetBindingRequest) (*alm_settings.GetBindingResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/get_binding", API)
	v := new(alm_settings.GetBindingResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// List - List DevOps Platform setting available for a given project, sorted by DevOps Platform key
// Requires the 'Administer project' permission if the 'project' parameter is provided, requires the 'Create Projects' permission otherwise.
func (s *AlmSettings) List(ctx context.Context, r alm_settings.ListRequest) (*alm_settings.ListResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/list", API)
	v := new(alm_settings.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ListDefinitions - List DevOps Platform Settings, sorted by created date.
// Requires the 'Administer System' permission
func (s *AlmSettings) ListDefinitions(ctx context.Context, r alm_settings.ListDefinitionsRequest) (*alm_settings.ListDefinitionsResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/list_definitions", API)
	v := new(alm_settings.ListDefinitionsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// SetAzureBinding - Bind a Azure DevOps instance to a project.
// If the project was already bound to a previous Azure DevOps instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
// Changelog:
func (s *AlmSettings) SetAzureBinding(ctx context.Context, r alm_settings.SetAzureBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_azure_binding", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetBitbucketBinding - Bind a Bitbucket instance to a project.
// If the project was already bound to a previous Bitbucket instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
// Changelog:
func (s *AlmSettings) SetBitbucketBinding(ctx context.Context, r alm_settings.SetBitbucketBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_bitbucket_binding", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetBitbucketcloudBinding - Bind a Bitbucket Cloud setting to a project.
// If the project was already bound to a different Bitbucket Cloud setting, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.7
// Changelog:
func (s *AlmSettings) SetBitbucketcloudBinding(ctx context.Context, r alm_settings.SetBitbucketcloudBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_bitbucketcloud_binding", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetGithubBinding - Bind a GitHub instance to a project.
// If the project was already bound to a previous GitHub instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
// Changelog:
//   8.3: Add 'summaryCommentEnabled' param to enable/disable of putting analysis summary in a conversation tab of GitHub
func (s *AlmSettings) SetGithubBinding(ctx context.Context, r alm_settings.SetGithubBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_github_binding", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// SetGitlabBinding - Bind a GitLab instance to a project.
// If the project was already bound to a previous Gitlab instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
// Since 8.1
// Changelog:
//   8.2: Parameter 'repository' was added
func (s *AlmSettings) SetGitlabBinding(ctx context.Context, r alm_settings.SetGitlabBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_gitlab_binding", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAzure - Update Azure instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//   8.7: Parameter 'personalAccessToken' is no longer required
//   8.6: Parameter 'URL' was added
func (s *AlmSettings) UpdateAzure(ctx context.Context, r alm_settings.UpdateAzureRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_azure", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBitbucket - Update Bitbucket instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//   8.7: Parameter 'personalAccessToken' is no longer required
func (s *AlmSettings) UpdateBitbucket(ctx context.Context, r alm_settings.UpdateBitbucketRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_bitbucket", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBitbucketcloud - Update Bitbucket Cloud Setting.
// Requires the 'Administer System' permission
// Since 8.7
// Changelog:
func (s *AlmSettings) UpdateBitbucketcloud(ctx context.Context, r alm_settings.UpdateBitbucketcloudRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_bitbucketcloud", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGithub - Update GitHub instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//   9.7: Optional parameter 'webhookSecret' was added
//   8.7: Parameter 'privateKey' is no longer required
//   8.7: Parameter 'clientSecret' is no longer required
func (s *AlmSettings) UpdateGithub(ctx context.Context, r alm_settings.UpdateGithubRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_github", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGitlab - Update GitLab instance Setting.
// Requires the 'Administer System' permission
// Since 8.1
// Changelog:
//   8.7: Parameter 'personalAccessToken' is no longer required
//   8.2: Parameter 'URL' was added
func (s *AlmSettings) UpdateGitlab(ctx context.Context, r alm_settings.UpdateGitlabRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_gitlab", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Validate - Validate an DevOps Platform Setting by checking connectivity and permissions
// Requires the 'Administer System' permission
func (s *AlmSettings) Validate(ctx context.Context, r alm_settings.ValidateRequest) (*alm_settings.ValidateResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/validate", API)
	v := new(alm_settings.ValidateResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
