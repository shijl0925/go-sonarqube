package ce

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ActivityRequest Search for tasks.<br> Requires the system administration permission, or project administration permission if component is set.
type ActivityRequest struct {
	Component      string `url:"component,omitempty"`      // Since 8.0;Key of the component (project) to filter on
	MaxExecutedAt  string `url:"maxExecutedAt,omitempty"`  // Maximum date of end of task processing (inclusive)
	MinSubmittedAt string `url:"minSubmittedAt,omitempty"` // Minimum date of task submission (inclusive)
	OnlyCurrents   string `url:"onlyCurrents,omitempty"`   // Filter on the last tasks (only the most recent finished task by project)
	Q              string `url:"q,omitempty"`              // Since 5.5;Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li><li>task ids that are exactly the same as the supplied string</li></ul>
	Status         string `url:"status,omitempty"`         // Comma separated list of task statuses
	Type           string `url:"type,omitempty"`           // Task type
}

// ActivityResponse is the response for ActivityRequest
type ActivityResponse struct {
	Paging paging.Paging `json:"paging,omitempty"`
	Tasks  []struct {
		AnalysisId         string   `json:"analysisId,omitempty"`
		ComponentId        string   `json:"componentId,omitempty"`
		ComponentKey       string   `json:"componentKey,omitempty"`
		ComponentName      string   `json:"componentName,omitempty"`
		ComponentQualifier string   `json:"componentQualifier,omitempty"`
		ExecutedAt         string   `json:"executedAt,omitempty"`
		ExecutionTimeMs    float64  `json:"executionTimeMs,omitempty"`
		HasScannerContext  bool     `json:"hasScannerContext,omitempty"`
		Id                 string   `json:"id,omitempty"`
		InfoMessages       []string `json:"infoMessages,omitempty"`
		StartedAt          string   `json:"startedAt,omitempty"`
		Status             string   `json:"status,omitempty"`
		SubmittedAt        string   `json:"submittedAt,omitempty"`
		SubmitterLogin     string   `json:"submitterLogin,omitempty"`
		Type               string   `json:"type,omitempty"`
		WarningCount       float64  `json:"warningCount,omitempty"`
		Warnings           []string `json:"warnings,omitempty"`
	} `json:"tasks,omitempty"`
}

// GetPaging extracts the paging from ActivityResponse
func (r *ActivityResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// ActivityResponseAll is the collection for ActivityRequest
type ActivityResponseAll struct {
	Tasks []struct {
		AnalysisId         string   `json:"analysisId,omitempty"`
		ComponentId        string   `json:"componentId,omitempty"`
		ComponentKey       string   `json:"componentKey,omitempty"`
		ComponentName      string   `json:"componentName,omitempty"`
		ComponentQualifier string   `json:"componentQualifier,omitempty"`
		ExecutedAt         string   `json:"executedAt,omitempty"`
		ExecutionTimeMs    float64  `json:"executionTimeMs,omitempty"`
		HasScannerContext  bool     `json:"hasScannerContext,omitempty"`
		Id                 string   `json:"id,omitempty"`
		InfoMessages       []string `json:"infoMessages,omitempty"`
		StartedAt          string   `json:"startedAt,omitempty"`
		Status             string   `json:"status,omitempty"`
		SubmittedAt        string   `json:"submittedAt,omitempty"`
		SubmitterLogin     string   `json:"submitterLogin,omitempty"`
		Type               string   `json:"type,omitempty"`
		WarningCount       float64  `json:"warningCount,omitempty"`
		Warnings           []string `json:"warnings,omitempty"`
	} `json:"tasks,omitempty"`
}

// ActivityStatusRequest Returns CE activity related metrics.<br>Requires 'Administer System' permission or 'Administer' rights on the specified project.
type ActivityStatusRequest struct {
	Component string `url:"component,omitempty"` // Key of the component (project) to filter on
}

// ActivityStatusResponse is the response for ActivityStatusRequest
type ActivityStatusResponse struct {
	Failing     float64 `json:"failing,omitempty"`
	InProgress  float64 `json:"inProgress,omitempty"`
	Pending     float64 `json:"pending,omitempty"`
	PendingTime float64 `json:"pendingTime,omitempty"`
}

// ComponentRequest Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component.
type ComponentRequest struct {
	Component string `url:"component"` //
}

// ComponentResponse is the response for ComponentRequest
type ComponentResponse struct {
	Current struct {
		AnalysisId         string  `json:"analysisId,omitempty"`
		ComponentId        string  `json:"componentId,omitempty"`
		ComponentKey       string  `json:"componentKey,omitempty"`
		ComponentName      string  `json:"componentName,omitempty"`
		ComponentQualifier string  `json:"componentQualifier,omitempty"`
		ErrorMessage       string  `json:"errorMessage,omitempty"`
		ErrorType          string  `json:"errorType,omitempty"`
		ExecutionTimeMs    float64 `json:"executionTimeMs,omitempty"`
		FinishedAt         string  `json:"finishedAt,omitempty"`
		HasErrorStacktrace bool    `json:"hasErrorStacktrace,omitempty"`
		HasScannerContext  bool    `json:"hasScannerContext,omitempty"`
		Id                 string  `json:"id,omitempty"`
		Organization       string  `json:"organization,omitempty"`
		StartedAt          string  `json:"startedAt,omitempty"`
		Status             string  `json:"status,omitempty"`
		SubmittedAt        string  `json:"submittedAt,omitempty"`
		Type               string  `json:"type,omitempty"`
	} `json:"current,omitempty"`
	Queue []struct {
		ComponentId        string `json:"componentId,omitempty"`
		ComponentKey       string `json:"componentKey,omitempty"`
		ComponentName      string `json:"componentName,omitempty"`
		ComponentQualifier string `json:"componentQualifier,omitempty"`
		Id                 string `json:"id,omitempty"`
		Organization       string `json:"organization,omitempty"`
		Status             string `json:"status,omitempty"`
		SubmittedAt        string `json:"submittedAt,omitempty"`
		Type               string `json:"type,omitempty"`
	} `json:"queue,omitempty"`
}

// TaskRequest Give Compute Engine task details such as type, status, duration and associated component.<br/>Requires one of the following permissions: <ul><li>'Administer' at global or project level</li><li>'Execute Analysis' at global or project level</li></ul>Since 6.1, field "logs" is deprecated and its value is always false.
type TaskRequest struct {
	AdditionalFields string `url:"additionalFields,omitempty"` // Since 6.1;Comma-separated list of the optional fields to be returned in response.
	Id               string `url:"id"`                         // Id of task
}

// TaskResponse is the response for TaskRequest
type TaskResponse struct {
	Task struct {
		AnalysisId         string   `json:"analysisId,omitempty"`
		ComponentId        string   `json:"componentId,omitempty"`
		ComponentKey       string   `json:"componentKey,omitempty"`
		ComponentName      string   `json:"componentName,omitempty"`
		ComponentQualifier string   `json:"componentQualifier,omitempty"`
		ExecutedAt         string   `json:"executedAt,omitempty"`
		ExecutionTimeMs    float64  `json:"executionTimeMs,omitempty"`
		HasScannerContext  bool     `json:"hasScannerContext,omitempty"`
		Id                 string   `json:"id,omitempty"`
		StartedAt          string   `json:"startedAt,omitempty"`
		Status             string   `json:"status,omitempty"`
		SubmittedAt        string   `json:"submittedAt,omitempty"`
		SubmitterLogin     string   `json:"submitterLogin,omitempty"`
		Type               string   `json:"type,omitempty"`
		WarningCount       float64  `json:"warningCount,omitempty"`
		Warnings           []string `json:"warnings,omitempty"`
	} `json:"task,omitempty"`
}
