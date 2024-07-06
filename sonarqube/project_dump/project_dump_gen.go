package project_dump

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ExportRequest Triggers project dump so that the project can be imported to another SonarQube server (see api/project_dump/import, available in Enterprise Edition). Requires the 'Administer' permission.
type ExportRequest struct {
	Key string `form:"key"` //
}

// ExportResponse is the response for ExportRequest
type ExportResponse struct {
	ProjectId   string `json:"projectId,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	TaskId      string `json:"taskId,omitempty"`
}

// ImportRequest Triggers the import of a project dump. Permission 'Administer' is required. This feature is provided by the Governance plugin.
type ImportRequest struct {
	Key string `form:"key"` //
}

// ImportResponse is the response for ImportRequest
type ImportResponse struct {
	ProjectId   string `json:"projectId,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	TaskId      string `json:"taskId,omitempty"`
}
