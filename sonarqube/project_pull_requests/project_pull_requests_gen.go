package project_pull_requests

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// DeleteRequest Delete a pull request.<br/>Requires 'Administer' rights on the specified project.
type DeleteRequest struct {
	Project     string `json:"project"`     // Project key
	PullRequest string `json:"pullRequest"` // Pull request id
}

// ListRequest List the pull requests of a project.<br/>One of the following permissions is required: <ul><li>'Browse' rights on the specified project</li><li>'Execute Analysis' rights on the specified project</li></ul>
type ListRequest struct {
	Project string `url:"project"` // Project key
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	PullRequests []struct {
		AnalysisDate string `json:"analysisDate,omitempty"`
		Base         string `json:"base,omitempty"`
		Branch       string `json:"branch,omitempty"`
		Key          string `json:"key,omitempty"`
		Status       struct {
			QualityGateStatus string `json:"qualityGateStatus,omitempty"`
		} `json:"status,omitempty"`
		Target string `json:"target,omitempty"`
		Title  string `json:"title,omitempty"`
		Url    string `json:"url,omitempty"`
	} `json:"pullRequests,omitempty"`
}
