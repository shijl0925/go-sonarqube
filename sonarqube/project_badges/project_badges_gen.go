package project_badges

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AiCodeAssuranceRequest Generate a badge for project's AI assurance as an SVG.<br/>Requires 'Browse' permission on the specified project.
type AiCodeAssuranceRequest struct {
	Project string `url:"project"`         // Project or application key
	Token   string `url:"token,omitempty"` // Project badge token
}

// AiCodeAssuranceResponse is the response for AiCodeAssuranceRequest
type AiCodeAssuranceResponse string

// MeasureRequest Generate badge for project's measure as an SVG.<br/>Requires 'Browse' permission on the specified project.
type MeasureRequest struct {
	Branch  string `url:"branch,omitempty"` // Branch key
	Metric  string `url:"metric"`           // Metric key
	Project string `url:"project"`          // Project or application key
	Token   string `url:"token,omitempty"`  // Project badge token
}

// MeasureResponse is the response for MeasureRequest
type MeasureResponse string

// QualityGateRequest Generate badge for project's quality gate as an SVG.<br/>Requires 'Browse' permission on the specified project.
type QualityGateRequest struct {
	Branch  string `url:"branch,omitempty"` // Branch key
	Project string `url:"project"`          // Project or application key
	Token   string `url:"token,omitempty"`  // Project badge token
}

// QualityGateResponse is the response for QualityGateRequest
type QualityGateResponse string

// RenewTokenRequest Creates new token replacing any existing token for project or application badge access for private projects and applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Administer' permission on the specified project or application.
type RenewTokenRequest struct {
	Project string `form:"project"` // Project or application key
}

// TokenRequest Retrieve a token to use for project or application badge access for private projects or applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Browse' permission on the specified project or application.
type TokenRequest struct {
	Project string `url:"project"` // Project or application key
}

// TokenResponse is the response for TokenRequest
type TokenResponse struct {
	Token string `json:"token,omitempty"`
}
