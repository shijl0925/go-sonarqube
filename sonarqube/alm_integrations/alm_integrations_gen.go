package alm_integrations

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ImportAzureProjectRequest Create a SonarQube project with the information from the provided Azure DevOps project.<br/>Autoconfigure pull request decoration mechanism.<br/>Requires the 'Create Projects' permission
// Deprecated: this action has been deprecated since version 10.5
type ImportAzureProjectRequest struct {
	AlmSetting             string `form:"almSetting,omitempty"`             // DevOps Platform configuration key. This parameter is optional if you have only one Azure integration.
	NewCodeDefinitionType  string `form:"newCodeDefinitionType,omitempty"`  // Project New Code Definition Type<br/>New code definitions of the following types are allowed:<ul><li>PREVIOUS_VERSION</li><li>NUMBER_OF_DAYS</li><li>REFERENCE_BRANCH - will default to the main branch.</li></ul>
	NewCodeDefinitionValue string `form:"newCodeDefinitionValue,omitempty"` // Project New Code Definition Value<br/>For each new code definition type, a different value is expected:<ul><li>no value, when the new code definition type is PREVIOUS_VERSION and REFERENCE_BRANCH</li><li>a number between 1 and 90, when the new code definition type is NUMBER_OF_DAYS</li></ul>
	ProjectName            string `form:"projectName,omitempty"`            // Azure project name
	RepositoryName         string `form:"repositoryName,omitempty"`         // Azure repository name
}

// ImportBitbucketcloudRepoRequest Create a SonarQube project with the information from the provided Bitbucket Cloud repository.<br/>Autoconfigure pull request decoration mechanism.<br/>Requires the 'Create Projects' permission
// Deprecated: this action has been deprecated since version 10.5
type ImportBitbucketcloudRepoRequest struct {
	AlmSetting             string `form:"almSetting,omitempty"`             // DevOps Platform configuration key. This parameter is optional if you have only one BitBucket Cloud integration.
	NewCodeDefinitionType  string `form:"newCodeDefinitionType,omitempty"`  // Project New Code Definition Type<br/>New code definitions of the following types are allowed:<ul><li>PREVIOUS_VERSION</li><li>NUMBER_OF_DAYS</li><li>REFERENCE_BRANCH - will default to the main branch.</li></ul>
	NewCodeDefinitionValue string `form:"newCodeDefinitionValue,omitempty"` // Project New Code Definition Value<br/>For each new code definition type, a different value is expected:<ul><li>no value, when the new code definition type is PREVIOUS_VERSION and REFERENCE_BRANCH</li><li>a number between 1 and 90, when the new code definition type is NUMBER_OF_DAYS</li></ul>
	RepositorySlug         string `form:"repositorySlug,omitempty"`         // Bitbucket Cloud repository slug
}

// ImportBitbucketserverProjectRequest Create a SonarQube project with the information from the provided BitbucketServer project.<br/>Autoconfigure pull request decoration mechanism.<br/>Requires the 'Create Projects' permission
// Deprecated: this action has been deprecated since version 10.5
type ImportBitbucketserverProjectRequest struct {
	AlmSetting             string `form:"almSetting,omitempty"`             // DevOps Platform configuration key. This parameter is optional if you have only one BitBucket Server integration.
	NewCodeDefinitionType  string `form:"newCodeDefinitionType,omitempty"`  // Project New Code Definition Type<br/>New code definitions of the following types are allowed:<ul><li>PREVIOUS_VERSION</li><li>NUMBER_OF_DAYS</li><li>REFERENCE_BRANCH - will default to the main branch.</li></ul>
	NewCodeDefinitionValue string `form:"newCodeDefinitionValue,omitempty"` // Project New Code Definition Value<br/>For each new code definition type, a different value is expected:<ul><li>no value, when the new code definition type is PREVIOUS_VERSION and REFERENCE_BRANCH</li><li>a number between 1 and 90, when the new code definition type is NUMBER_OF_DAYS</li></ul>
	ProjectKey             string `form:"projectKey,omitempty"`             // BitbucketServer project key
	RepositorySlug         string `form:"repositorySlug,omitempty"`         // BitbucketServer repository slug
}

// ImportGithubProjectRequest Create a SonarQube project with the information from the provided GitHub repository.<br/>Autoconfigure pull request decoration mechanism. If Automatic Provisioning is enable for GitHub, it will also synchronize permissions from the repository.<br/>Requires the 'Create Projects' permission
// Deprecated: this action has been deprecated since version 10.5
type ImportGithubProjectRequest struct {
	AlmSetting             string `form:"almSetting,omitempty"`             // DevOps Platform configuration key. This parameter is optional if you have only one GitHub integration.
	NewCodeDefinitionType  string `form:"newCodeDefinitionType,omitempty"`  // Project New Code Definition Type<br/>New code definitions of the following types are allowed:<ul><li>PREVIOUS_VERSION</li><li>NUMBER_OF_DAYS</li><li>REFERENCE_BRANCH - will default to the main branch.</li></ul>
	NewCodeDefinitionValue string `form:"newCodeDefinitionValue,omitempty"` // Project New Code Definition Value<br/>For each new code definition type, a different value is expected:<ul><li>no value, when the new code definition type is PREVIOUS_VERSION and REFERENCE_BRANCH</li><li>a number between 1 and 90, when the new code definition type is NUMBER_OF_DAYS</li></ul>
	RepositoryKey          string `form:"repositoryKey,omitempty"`          // GitHub repository key (organization/repoSlug
}

// ImportGitlabProjectRequest Import a GitLab project to SonarQube, creating a new project and configuring MR decoration<br/>Requires the 'Create Projects' permission
// Deprecated: this action has been deprecated since version 10.5
type ImportGitlabProjectRequest struct {
	AlmSetting             string `form:"almSetting,omitempty"`             // DevOps Platform configuration key. This parameter is optional if you have only one GitLab integration.
	GitlabProjectId        string `form:"gitlabProjectId,omitempty"`        // GitLab project ID
	NewCodeDefinitionType  string `form:"newCodeDefinitionType,omitempty"`  // Project New Code Definition Type<br/>New code definitions of the following types are allowed:<ul><li>PREVIOUS_VERSION</li><li>NUMBER_OF_DAYS</li><li>REFERENCE_BRANCH - will default to the main branch.</li></ul>
	NewCodeDefinitionValue string `form:"newCodeDefinitionValue,omitempty"` // Project New Code Definition Value<br/>For each new code definition type, a different value is expected:<ul><li>no value, when the new code definition type is PREVIOUS_VERSION and REFERENCE_BRANCH</li><li>a number between 1 and 90, when the new code definition type is NUMBER_OF_DAYS</li></ul>
}

// ListAzureProjectsRequest List Azure projects<br/>Requires the 'Create Projects' permission
type ListAzureProjectsRequest struct {
	AlmSetting string `form:"almSetting,omitempty"` // DevOps Platform setting key
}

// ListAzureProjectsResponse is the response for ListAzureProjectsRequest
type ListAzureProjectsResponse struct {
	Projects []struct {
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"projects,omitempty"`
}

// ListBitbucketserverProjectsRequest List the Bitbucket Server projects<br/>Requires the 'Create Projects' permission
type ListBitbucketserverProjectsRequest struct {
	AlmSetting string `form:"almSetting,omitempty"` // DevOps Platform setting key
}

// ListBitbucketserverProjectsResponse is the response for ListBitbucketserverProjectsRequest
type ListBitbucketserverProjectsResponse struct {
	Projects []struct {
		Key  string `json:"key,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"projects,omitempty"`
}

// SearchAzureReposRequest Search the Azure repositories<br/>Requires the 'Create Projects' permission
type SearchAzureReposRequest struct {
	AlmSetting  string `form:"almSetting,omitempty"`  // DevOps Platform setting key
	ProjectName string `form:"projectName,omitempty"` // Project name filter
	SearchQuery string `form:"searchQuery,omitempty"` // Search query filter
}

// SearchAzureReposResponse is the response for SearchAzureReposRequest
type SearchAzureReposResponse struct {
	Repositories []struct {
		Name        string `json:"name,omitempty"`
		ProjectName string `json:"projectName,omitempty"`
	} `json:"repositories,omitempty"`
}

// SearchBitbucketcloudReposRequest Search the Bitbucket Cloud repositories<br/>Requires the 'Create Projects' permission
type SearchBitbucketcloudReposRequest struct {
	AlmSetting     string `form:"almSetting,omitempty"`     // DevOps Platform setting key
	RepositoryName string `form:"repositoryName,omitempty"` // Repository name filter
}

// SearchBitbucketcloudReposResponse is the response for SearchBitbucketcloudReposRequest
type SearchBitbucketcloudReposResponse struct {
	IsLastPage   bool          `json:"isLastPage,omitempty"`
	Paging       paging.Paging `json:"paging,omitempty"`
	Repositories []struct {
		Name         string `json:"name,omitempty"`
		ProjectKey   string `json:"projectKey,omitempty"`
		Slug         string `json:"slug,omitempty"`
		Uuid         string `json:"uuid,omitempty"`
		Workspace    string `json:"workspace,omitempty"`
		SqProjectKey string `json:"sqProjectKey,omitempty"`
	} `json:"repositories,omitempty"`
}

// GetPaging extracts the paging from SearchBitbucketcloudReposResponse
func (r *SearchBitbucketcloudReposResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchBitbucketcloudReposResponseAll is the collection for SearchBitbucketcloudReposRequest
type SearchBitbucketcloudReposResponseAll struct {
	IsLastPage   bool `json:"isLastPage,omitempty"`
	Repositories []struct {
		Name         string `json:"name,omitempty"`
		ProjectKey   string `json:"projectKey,omitempty"`
		Slug         string `json:"slug,omitempty"`
		Uuid         string `json:"uuid,omitempty"`
		Workspace    string `json:"workspace,omitempty"`
		SqProjectKey string `json:"sqProjectKey,omitempty"`
	} `json:"repositories,omitempty"`
}

// SearchBitbucketserverReposRequest Search the Bitbucket Server repositories with REPO_ADMIN access<br/>Requires the 'Create Projects' permission
type SearchBitbucketserverReposRequest struct {
	AlmSetting     string `form:"almSetting,omitempty"`     // DevOps Platform setting key
	ProjectName    string `form:"projectName,omitempty"`    // Project name filter
	RepositoryName string `form:"repositoryName,omitempty"` // Repository name filter
}

// SearchBitbucketserverReposResponse is the response for SearchBitbucketserverReposRequest
type SearchBitbucketserverReposResponse struct {
	IsLastPage   bool `json:"isLastPage,omitempty"`
	Repositories []struct {
		Name         string `json:"name,omitempty"`
		ProjectKey   string `json:"projectKey,omitempty"`
		Slug         string `json:"slug,omitempty"`
		Uuid         string `json:"uuid,omitempty"`
		Workspace    string `json:"workspace,omitempty"`
		SqProjectKey string `json:"sqProjectKey,omitempty"`
	} `json:"repositories,omitempty"`
}

// SearchGitlabReposRequest Search the GitLab projects.<br/>Requires the 'Create Projects' permission
type SearchGitlabReposRequest struct {
	AlmSetting  string `form:"almSetting,omitempty"`  // DevOps Platform setting key
	ProjectName string `form:"projectName,omitempty"` // Project name filter
}

// SearchGitlabReposResponse is the response for SearchGitlabReposRequest
type SearchGitlabReposResponse struct {
	Paging       paging.Paging `json:"paging,omitempty"`
	Repositories []struct {
		Id       float64 `json:"id,omitempty"`
		Name     string  `json:"name,omitempty"`
		PathName string  `json:"pathName,omitempty"`
		PathSlug string  `json:"pathSlug,omitempty"`
		Slug     string  `json:"slug,omitempty"`
		Url      string  `json:"url,omitempty"`
	} `json:"repositories,omitempty"`
}

// GetPaging extracts the paging from SearchGitlabReposResponse
func (r *SearchGitlabReposResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchGitlabReposResponseAll is the collection for SearchGitlabReposRequest
type SearchGitlabReposResponseAll struct {
	Repositories []struct {
		Id       float64 `json:"id,omitempty"`
		Name     string  `json:"name,omitempty"`
		PathName string  `json:"pathName,omitempty"`
		PathSlug string  `json:"pathSlug,omitempty"`
		Slug     string  `json:"slug,omitempty"`
		Url      string  `json:"url,omitempty"`
	} `json:"repositories,omitempty"`
}

// SetPatRequest Set a Personal Access Token for the given DevOps Platform setting<br/>Requires the 'Create Projects' permission
type SetPatRequest struct {
	AlmSetting string `form:"almSetting,omitempty"` // DevOps Platform configuration key. This parameter is optional if you have only one single DevOps Platform integration.
	Pat        string `form:"pat,omitempty"`        // Personal Access Token
	Username   string `form:"username,omitempty"`   // Username
}
