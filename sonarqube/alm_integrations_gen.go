package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/alm_integrations"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AlmIntegrations service

func (s *AlmIntegrations) ImportAzureProject(r alm_integrations.ImportAzureProjectRequest) error {
	u := fmt.Sprintf("%s/alm_integrations/import_azure_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmIntegrations) ImportBitbucketcloudRepo(r alm_integrations.ImportBitbucketcloudRepoRequest) error {
	u := fmt.Sprintf("%s/alm_integrations/import_bitbucketcloud_repo", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmIntegrations) ImportBitbucketserverProject(r alm_integrations.ImportBitbucketserverProjectRequest) error {
	u := fmt.Sprintf("%s/alm_integrations/import_bitbucketserver_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmIntegrations) ImportGithubProject(r alm_integrations.ImportGithubProjectRequest) error {
	u := fmt.Sprintf("%s/alm_integrations/import_github_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmIntegrations) ImportGitlabProject(r alm_integrations.ImportGitlabProjectRequest) error {
	u := fmt.Sprintf("%s/alm_integrations/import_gitlab_project", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmIntegrations) ListAzureProjects(r alm_integrations.ListAzureProjectsRequest) (*alm_integrations.ListAzureProjectsResponse, error) {
	u := fmt.Sprintf("%s/alm_integrations/list_azure_projects", API)
	v := new(alm_integrations.ListAzureProjectsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmIntegrations) ListBitbucketserverProjects(r alm_integrations.ListBitbucketserverProjectsRequest) (*alm_integrations.ListBitbucketserverProjectsResponse, error) {
	u := fmt.Sprintf("%s/alm_integrations/list_bitbucketserver_projects", API)
	v := new(alm_integrations.ListBitbucketserverProjectsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmIntegrations) SearchAzureRepos(r alm_integrations.SearchAzureReposRequest) (*alm_integrations.SearchAzureReposResponse, error) {
	u := fmt.Sprintf("%s/alm_integrations/search_azure_repos", API)
	v := new(alm_integrations.SearchAzureReposResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmIntegrations) SearchBitbucketcloudRepos(r alm_integrations.SearchBitbucketcloudReposRequest, p paging.Params) (*alm_integrations.SearchBitbucketcloudReposResponse, error) {
	u := fmt.Sprintf("%s/alm_integrations/search_bitbucketcloud_repos", API)
	v := new(alm_integrations.SearchBitbucketcloudReposResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmIntegrations) SearchBitbucketcloudReposAll(r alm_integrations.SearchBitbucketcloudReposRequest) (*alm_integrations.SearchBitbucketcloudReposResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &alm_integrations.SearchBitbucketcloudReposResponseAll{}
	for {
		res, err := s.SearchBitbucketcloudRepos(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to alm_integrations.SearchBitbucketcloudRepos: %+v", err)
		}
		response.Repositories = append(response.Repositories, res.Repositories...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *AlmIntegrations) SearchBitbucketserverRepos(r alm_integrations.SearchBitbucketserverReposRequest) (*alm_integrations.SearchBitbucketserverReposResponse, error) {
	u := fmt.Sprintf("%s/alm_integrations/search_bitbucketserver_repos", API)
	v := new(alm_integrations.SearchBitbucketserverReposResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmIntegrations) SearchGitlabRepos(r alm_integrations.SearchGitlabReposRequest, p paging.Params) (*alm_integrations.SearchGitlabReposResponse, error) {
	u := fmt.Sprintf("%s/alm_integrations/search_gitlab_repos", API)
	v := new(alm_integrations.SearchGitlabReposResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmIntegrations) SearchGitlabReposAll(r alm_integrations.SearchGitlabReposRequest) (*alm_integrations.SearchGitlabReposResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &alm_integrations.SearchGitlabReposResponseAll{}
	for {
		res, err := s.SearchGitlabRepos(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to alm_integrations.SearchGitlabRepos: %+v", err)
		}
		response.Repositories = append(response.Repositories, res.Repositories...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *AlmIntegrations) SetPat(r alm_integrations.SetPatRequest) error {
	u := fmt.Sprintf("%s/alm_integrations/set_pat", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
