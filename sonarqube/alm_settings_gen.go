package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/alm_settings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AlmSettings service

func (s *AlmSettings) CountBinding(r alm_settings.CountBindingRequest) (*alm_settings.CountBindingResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/count_binding", API)
	v := new(alm_settings.CountBindingResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmSettings) CreateAzure(r alm_settings.CreateAzureRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_azure", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) CreateBitbucket(r alm_settings.CreateBitbucketRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_bitbucket", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) CreateBitbucketcloud(r alm_settings.CreateBitbucketcloudRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_bitbucketcloud", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) CreateGithub(r alm_settings.CreateGithubRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_github", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) CreateGitlab(r alm_settings.CreateGitlabRequest) error {
	u := fmt.Sprintf("%s/alm_settings/create_gitlab", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) Delete(r alm_settings.DeleteRequest) error {
	u := fmt.Sprintf("%s/alm_settings/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) DeleteBinding(r alm_settings.DeleteBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/delete_binding", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) GetBinding(r alm_settings.GetBindingRequest) (*alm_settings.GetBindingResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/get_binding", API)
	v := new(alm_settings.GetBindingResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmSettings) List(r alm_settings.ListRequest) (*alm_settings.ListResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/list", API)
	v := new(alm_settings.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmSettings) ListDefinitions(r alm_settings.ListDefinitionsRequest) (*alm_settings.ListDefinitionsResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/list_definitions", API)
	v := new(alm_settings.ListDefinitionsResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *AlmSettings) SetAzureBinding(r alm_settings.SetAzureBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_azure_binding", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) SetBitbucketBinding(r alm_settings.SetBitbucketBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_bitbucket_binding", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) SetBitbucketcloudBinding(r alm_settings.SetBitbucketcloudBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_bitbucketcloud_binding", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) SetGithubBinding(r alm_settings.SetGithubBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_github_binding", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) SetGitlabBinding(r alm_settings.SetGitlabBindingRequest) error {
	u := fmt.Sprintf("%s/alm_settings/set_gitlab_binding", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) UpdateAzure(r alm_settings.UpdateAzureRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_azure", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) UpdateBitbucket(r alm_settings.UpdateBitbucketRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_bitbucket", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) UpdateBitbucketcloud(r alm_settings.UpdateBitbucketcloudRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_bitbucketcloud", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) UpdateGithub(r alm_settings.UpdateGithubRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_github", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) UpdateGitlab(r alm_settings.UpdateGitlabRequest) error {
	u := fmt.Sprintf("%s/alm_settings/update_gitlab", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *AlmSettings) Validate(r alm_settings.ValidateRequest) (*alm_settings.ValidateResponse, error) {
	u := fmt.Sprintf("%s/alm_settings/validate", API)
	v := new(alm_settings.ValidateResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
