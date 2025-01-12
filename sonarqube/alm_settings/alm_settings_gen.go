package alm_settings

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// CountBindingRequest Count number of project bound to an DevOps Platform setting.<br/>Requires the 'Administer System' permission
type CountBindingRequest struct {
	AlmSetting string `url:"almSetting"` // DevOps Platform setting key
}

// CountBindingResponse is the response for CountBindingRequest
type CountBindingResponse struct {
	Key      string  `json:"key,omitempty"`
	Projects float64 `json:"projects,omitempty"`
}

// CreateAzureRequest Create Azure instance Setting. <br/>Requires the 'Administer System' permission
type CreateAzureRequest struct {
	Key                 string `form:"key"`                 // Unique key of the Azure Devops instance setting
	PersonalAccessToken string `form:"personalAccessToken"` // Azure Devops personal access token
	Url                 string `form:"url"`                 // Azure API URL
}

// CreateBitbucketRequest Create Bitbucket instance Setting. <br/>Requires the 'Administer System' permission
type CreateBitbucketRequest struct {
	Key                 string `form:"key"`                 // Unique key of the Bitbucket instance setting
	PersonalAccessToken string `form:"personalAccessToken"` // Bitbucket personal access token
	Url                 string `form:"url"`                 // BitBucket server API URL
}

// CreateBitbucketcloudRequest Configure a new instance of Bitbucket Cloud. <br/>Requires the 'Administer System' permission
type CreateBitbucketcloudRequest struct {
	ClientId     string `form:"clientId"`     // Bitbucket Cloud Client ID
	ClientSecret string `form:"clientSecret"` // Bitbucket Cloud Client Secret
	Key          string `form:"key"`          // Unique key of the Bitbucket Cloud setting
	Workspace    string `form:"workspace"`    // Bitbucket Cloud workspace ID
}

// CreateGithubRequest Create GitHub instance Setting. <br/>Requires the 'Administer System' permission
type CreateGithubRequest struct {
	AppId         string `form:"appId"`                   // GitHub App ID
	ClientId      string `form:"clientId"`                // GitHub App Client ID
	ClientSecret  string `form:"clientSecret"`            // GitHub App Client Secret
	Key           string `form:"key"`                     // Unique key of the GitHub instance setting
	PrivateKey    string `form:"privateKey"`              // GitHub App private key
	Url           string `form:"url"`                     // GitHub API URL
	WebhookSecret string `form:"webhookSecret,omitempty"` // GitHub App Webhook Secret
}

// CreateGitlabRequest Create GitLab instance Setting. <br/>Requires the 'Administer System' permission
type CreateGitlabRequest struct {
	Key                 string `form:"key"`                 // Unique key of the GitLab instance setting
	PersonalAccessToken string `form:"personalAccessToken"` // GitLab personal access token
	Url                 string `form:"url"`                 // GitLab API URL
}

// DeleteRequest Delete an DevOps Platform Setting.<br/>Requires the 'Administer System' permission
type DeleteRequest struct {
	Key string `form:"key"` // DevOps Platform Setting key
}

// DeleteBindingRequest Delete the DevOps Platform binding of a project.<br/>Requires the 'Administer' permission on the project
type DeleteBindingRequest struct {
	Project string `form:"project"` // Project key
}

// GetBindingRequest Get DevOps Platform binding of a given project.<br/>Requires the 'Browse' permission on the project
type GetBindingRequest struct {
	Project string `url:"project"` // Project key
}

// GetBindingResponse is the response for GetBindingRequest
type GetBindingResponse struct {
	Alm                   string `json:"alm,omitempty"`
	Key                   string `json:"key,omitempty"`
	Monorepo              bool   `json:"monorepo,omitempty"`
	Repository            string `json:"repository,omitempty"`
	SummaryCommentEnabled bool   `json:"summaryCommentEnabled,omitempty"`
	Url                   string `json:"url,omitempty"`
}

// ListRequest List DevOps Platform setting available for a given project, sorted by DevOps Platform key<br/>Requires the 'Administer project' permission if the 'project' parameter is provided, requires the 'Create Projects' permission otherwise.
type ListRequest struct {
	Project string `url:"project,omitempty"` // Project key
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	AlmSettings []struct {
		Alm string `json:"alm,omitempty"`
		Key string `json:"key,omitempty"`
		Url string `json:"url,omitempty"`
	} `json:"almSettings,omitempty"`
}

// ListDefinitionsRequest List DevOps Platform Settings, sorted by created date.<br/>Requires the 'Administer System' permission
type ListDefinitionsRequest struct{}

// ListDefinitionsResponse is the response for ListDefinitionsRequest
type ListDefinitionsResponse struct {
	Azure          []string `json:"azure,omitempty"`
	Bitbucket      []string `json:"bitbucket,omitempty"`
	Bitbucketcloud []string `json:"bitbucketcloud,omitempty"`
	Github         []struct {
		AppId    string `json:"appId,omitempty"`
		ClientId string `json:"clientId,omitempty"`
		Key      string `json:"key,omitempty"`
		Url      string `json:"url,omitempty"`
	} `json:"github,omitempty"`
	Gitlab []string `json:"gitlab,omitempty"`
}

// SetAzureBindingRequest Bind a Azure DevOps instance to a project.<br/>If the project was already bound to a previous Azure DevOps instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
type SetAzureBindingRequest struct {
	AlmSetting               string `form:"almSetting"`                         // Azure DevOps setting key
	InlineAnnotationsEnabled string `form:"inlineAnnotationsEnabled,omitempty"` // Since 2025.1;Enable inline annotations during Pull Request decoration for this project
	Monorepo                 string `form:"monorepo"`                           // Since 8.7;Is this project part of a monorepo
	Project                  string `form:"project"`                            // SonarQube project key
	ProjectName              string `form:"projectName"`                        // Since 8.6;Azure project name
	RepositoryName           string `form:"repositoryName"`                     // Since 8.6;Azure repository name
}

// SetBitbucketBindingRequest Bind a Bitbucket instance to a project.<br/>If the project was already bound to a previous Bitbucket instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
type SetBitbucketBindingRequest struct {
	AlmSetting string `form:"almSetting"` // Bitbucket Server setting key
	Monorepo   string `form:"monorepo"`   // Since 8.7;Is this project part of a monorepo
	Project    string `form:"project"`    // Project key
	Repository string `form:"repository"` // Bitbucket Server repository key
	Slug       string `form:"slug"`       // Bitbucket repository slug
}

// SetBitbucketcloudBindingRequest Bind a Bitbucket Cloud setting to a project.<br/>If the project was already bound to a different Bitbucket Cloud setting, the binding will be updated to the new one.Requires the 'Administer' permission on the project
type SetBitbucketcloudBindingRequest struct {
	AlmSetting string `form:"almSetting"` // Bitbucket Cloud setting key
	Monorepo   string `form:"monorepo"`   // Since 8.8;Is this project part of a monorepo
	Project    string `form:"project"`    // Project key
	Repository string `form:"repository"` // Bitbucket Cloud repository key
}

// SetGithubBindingRequest Bind a GitHub instance to a project.<br/>If the project was already bound to a previous GitHub instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
type SetGithubBindingRequest struct {
	AlmSetting            string `form:"almSetting"`                      // GitHub setting key
	Monorepo              string `form:"monorepo"`                        // Since 8.7;Is this project part of a monorepo
	Project               string `form:"project"`                         // Project key
	Repository            string `form:"repository"`                      // GitHub Repository
	SummaryCommentEnabled string `form:"summaryCommentEnabled,omitempty"` // Enable/disable summary in PR discussion tab
}

// SetGitlabBindingRequest Bind a GitLab instance to a project.<br/>If the project was already bound to a previous Gitlab instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
type SetGitlabBindingRequest struct {
	AlmSetting string `form:"almSetting"` // GitLab setting key
	Monorepo   string `form:"monorepo"`   // Since 8.7;Is this project part of a monorepo
	Project    string `form:"project"`    // Project key
	Repository string `form:"repository"` // GitLab project ID
}

// UpdateAzureRequest Update Azure instance Setting. <br/>Requires the 'Administer System' permission
type UpdateAzureRequest struct {
	Key                 string `form:"key"`                           // Unique key of the Azure instance setting
	NewKey              string `form:"newKey,omitempty"`              // Optional new value for an unique key of the Azure Devops instance setting
	PersonalAccessToken string `form:"personalAccessToken,omitempty"` // Azure Devops personal access token
	Url                 string `form:"url"`                           // Azure API URL
}

// UpdateBitbucketRequest Update Bitbucket instance Setting. <br/>Requires the 'Administer System' permission
type UpdateBitbucketRequest struct {
	Key                 string `form:"key"`                           // Unique key of the Bitbucket instance setting
	NewKey              string `form:"newKey,omitempty"`              // Optional new value for an unique key of the Bitbucket instance setting
	PersonalAccessToken string `form:"personalAccessToken,omitempty"` // Bitbucket personal access token
	Url                 string `form:"url"`                           // Bitbucket API URL
}

// UpdateBitbucketcloudRequest Update Bitbucket Cloud Setting. <br/>Requires the 'Administer System' permission
type UpdateBitbucketcloudRequest struct {
	ClientId     string `form:"clientId"`               // Bitbucket Cloud Client ID
	ClientSecret string `form:"clientSecret,omitempty"` // Optional new value for the Bitbucket Cloud client secret
	Key          string `form:"key"`                    // Unique key of the Bitbucket Cloud setting
	NewKey       string `form:"newKey,omitempty"`       // Optional new value for an unique key of the Bitbucket Cloud setting
	Workspace    string `form:"workspace"`              // Bitbucket Cloud workspace ID
}

// UpdateGithubRequest Update GitHub instance Setting. <br/>Requires the 'Administer System' permission
type UpdateGithubRequest struct {
	AppId         string `form:"appId"`                   // GitHub API ID
	ClientId      string `form:"clientId"`                // GitHub App Client ID
	ClientSecret  string `form:"clientSecret,omitempty"`  // GitHub App Client Secret
	Key           string `form:"key"`                     // Unique key of the GitHub instance setting
	NewKey        string `form:"newKey,omitempty"`        // Optional new value for an unique key of the GitHub instance setting
	PrivateKey    string `form:"privateKey,omitempty"`    // GitHub App private key
	Url           string `form:"url"`                     // GitHub API URL
	WebhookSecret string `form:"webhookSecret,omitempty"` // GitHub App Webhook Secret
}

// UpdateGitlabRequest Update GitLab instance Setting. <br/>Requires the 'Administer System' permission
type UpdateGitlabRequest struct {
	Key                 string `form:"key"`                           // Unique key of the GitLab instance setting
	NewKey              string `form:"newKey,omitempty"`              // Optional new value for an unique key of the GitLab instance setting
	PersonalAccessToken string `form:"personalAccessToken,omitempty"` // GitLab personal access token
	Url                 string `form:"url"`                           // GitLab API URL
}

// ValidateRequest Validate an DevOps Platform Setting by checking connectivity and permissions<br/>Requires the 'Administer System' permission
type ValidateRequest struct {
	Key string `url:"key"` // Unique key of the DevOps Platform settings
}

// ValidateResponse is the response for ValidateRequest
type ValidateResponse struct {
	Errors []struct {
		Msg string `json:"msg,omitempty"`
	} `json:"errors,omitempty"`
}
