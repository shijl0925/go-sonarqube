package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/alm_integrations"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AlmIntegrations service

// ImportAzureProject - Create a SonarQube project with the information from the provided Azure DevOps project.
// Autoconfigure pull request decoration mechanism.
// Requires the 'Create Projects' permission
// Since 8.6
// Deprecated since 10.5
// Changelog:
//
//	10.5: This endpoint is deprecated, please use its API v2 version /api/v2/dop-translation/bound-projects
//	10.3: Parameter almSetting becomes optional if you have only one configuration for Azure
//	10.3: Endpoint visibility change from internal to public
func (s *AlmIntegrations) ImportAzureProject(ctx context.Context, r alm_integrations.ImportAzureProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/import_azure_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ImportBitbucketcloudRepo - Create a SonarQube project with the information from the provided Bitbucket Cloud repository.
// Autoconfigure pull request decoration mechanism.
// Requires the 'Create Projects' permission
// Since 9.0
// Deprecated since 10.5
// Changelog:
//
//	10.5: This endpoint is deprecated, please use its API v2 version /api/v2/dop-translation/bound-projects
//	10.3: Parameter almSetting becomes optional if you have only one configuration for BitBucket Cloud
//	10.3: Endpoint visibility change from internal to public
func (s *AlmIntegrations) ImportBitbucketcloudRepo(ctx context.Context, r alm_integrations.ImportBitbucketcloudRepoRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/import_bitbucketcloud_repo", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ImportBitbucketserverProject - Create a SonarQube project with the information from the provided BitbucketServer project.
// Autoconfigure pull request decoration mechanism.
// Requires the 'Create Projects' permission
// Since 8.2
// Deprecated since 10.5
// Changelog:
//
//	10.5: This endpoint is deprecated, please use its API v2 version /api/v2/dop-translation/bound-projects
//	10.3: Parameter almSetting becomes optional if you have only one configuration for BitBucket Server
//	10.3: Endpoint visibility change from internal to public
func (s *AlmIntegrations) ImportBitbucketserverProject(ctx context.Context, r alm_integrations.ImportBitbucketserverProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/import_bitbucketserver_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ImportGithubProject - Create a SonarQube project with the information from the provided GitHub repository.
// Autoconfigure pull request decoration mechanism. If Automatic Provisioning is enable for GitHub, it will also synchronize permissions from the repository.
// Requires the 'Create Projects' permission
// Since 8.4
// Deprecated since 10.5
// Changelog:
//
//	10.5: This endpoint is deprecated, please use its API v2 version /api/v2/dop-translation/bound-projects
//	10.3: Parameter organization is not necessary anymore
//	10.3: Parameter almSetting becomes optional if you have only one configuration for GitHub
//	10.3: Endpoint visibility change from internal to public
func (s *AlmIntegrations) ImportGithubProject(ctx context.Context, r alm_integrations.ImportGithubProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/import_github_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ImportGitlabProject - Import a GitLab project to SonarQube, creating a new project and configuring MR decoration
// Requires the 'Create Projects' permission
// Since 8.5
// Deprecated since 10.5
// Changelog:
//
//	10.5: This endpoint is deprecated, please use its API v2 version /api/v2/dop-translation/bound-projects
//	10.3: Parameter almSetting becomes optional if you have only one configuration for GitLab
func (s *AlmIntegrations) ImportGitlabProject(ctx context.Context, r alm_integrations.ImportGitlabProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/import_gitlab_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListAzureProjects - List Azure projects
// Requires the 'Create Projects' permission
// Since 8.6
func (s *AlmIntegrations) ListAzureProjects(ctx context.Context, r alm_integrations.ListAzureProjectsRequest) (*alm_integrations.ListAzureProjectsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list_azure_projects", s.path)
	v := new(alm_integrations.ListAzureProjectsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// ListBitbucketserverProjects - List the Bitbucket Server projects
// Requires the 'Create Projects' permission
// Since 8.2
func (s *AlmIntegrations) ListBitbucketserverProjects(ctx context.Context, r alm_integrations.ListBitbucketserverProjectsRequest) (*alm_integrations.ListBitbucketserverProjectsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list_bitbucketserver_projects", s.path)
	v := new(alm_integrations.ListBitbucketserverProjectsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// SearchAzureRepos - Search the Azure repositories
// Requires the 'Create Projects' permission
// Since 8.6
func (s *AlmIntegrations) SearchAzureRepos(ctx context.Context, r alm_integrations.SearchAzureReposRequest) (*alm_integrations.SearchAzureReposResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search_azure_repos", s.path)
	v := new(alm_integrations.SearchAzureReposResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// SearchBitbucketcloudRepos - Search the Bitbucket Cloud repositories
// Requires the 'Create Projects' permission
// Since 9.0
func (s *AlmIntegrations) SearchBitbucketcloudRepos(ctx context.Context, r alm_integrations.SearchBitbucketcloudReposRequest, p paging.Params) (*alm_integrations.SearchBitbucketcloudReposResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search_bitbucketcloud_repos", s.path)
	v := new(alm_integrations.SearchBitbucketcloudReposResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *AlmIntegrations) SearchBitbucketcloudReposAll(ctx context.Context, r alm_integrations.SearchBitbucketcloudReposRequest) (*alm_integrations.SearchBitbucketcloudReposResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &alm_integrations.SearchBitbucketcloudReposResponseAll{}
	for {
		res, _, err := s.SearchBitbucketcloudRepos(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to alm_integrations.SearchBitbucketcloudRepos: %+v", err)
		}
		response.Repositories = append(response.Repositories, res.Repositories...)
		if res.GetPaging().End() {
			break
		}
		p.P++
	}
	return response, nil
}

// SearchBitbucketserverRepos - Search the Bitbucket Server repositories with REPO_ADMIN access
// Requires the 'Create Projects' permission
// Since 8.2
func (s *AlmIntegrations) SearchBitbucketserverRepos(ctx context.Context, r alm_integrations.SearchBitbucketserverReposRequest) (*alm_integrations.SearchBitbucketserverReposResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search_bitbucketserver_repos", s.path)
	v := new(alm_integrations.SearchBitbucketserverReposResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// SearchGitlabRepos - Search the GitLab projects.
// Requires the 'Create Projects' permission
// Since 8.5
func (s *AlmIntegrations) SearchGitlabRepos(ctx context.Context, r alm_integrations.SearchGitlabReposRequest, p paging.Params) (*alm_integrations.SearchGitlabReposResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search_gitlab_repos", s.path)
	v := new(alm_integrations.SearchGitlabReposResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *AlmIntegrations) SearchGitlabReposAll(ctx context.Context, r alm_integrations.SearchGitlabReposRequest) (*alm_integrations.SearchGitlabReposResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &alm_integrations.SearchGitlabReposResponseAll{}
	for {
		res, _, err := s.SearchGitlabRepos(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to alm_integrations.SearchGitlabRepos: %+v", err)
		}
		response.Repositories = append(response.Repositories, res.Repositories...)
		if res.GetPaging().End() {
			break
		}
		p.P++
	}
	return response, nil
}

// SetPat - Set a Personal Access Token for the given DevOps Platform setting
// Requires the 'Create Projects' permission
// Since 8.2
// Changelog:
//
//	10.3: Allow setting Personal Access Tokens for all DevOps platforms
//	10.3: Parameter almSetting becomes optional if you have only one DevOps Platform configuration
//	9.0: Bitbucket Cloud support and optional Username parameter were added
func (s *AlmIntegrations) SetPat(ctx context.Context, r alm_integrations.SetPatRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_pat", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
