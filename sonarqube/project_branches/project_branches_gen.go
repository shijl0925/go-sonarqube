package project_branches

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// DeleteRequest Delete a non-main branch of a project or application.<br/>Requires 'Administer' rights on the specified project or application.
type DeleteRequest struct {
	Branch  string `form:"branch"`  // Branch key
	Project string `form:"project"` // Project key
}

// ListRequest List the branches of a project or application.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project or application.
type ListRequest struct {
	Project string `url:"project"` // Project key
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Branches []struct {
		AnalysisDate      string `json:"analysisDate,omitempty"`
		BranchId          string `json:"branchId,omitempty"`
		ExcludedFromPurge bool   `json:"excludedFromPurge,omitempty"`
		IsMain            bool   `json:"isMain,omitempty"`
		Name              string `json:"name,omitempty"`
		Status            struct {
			QualityGateStatus string `json:"qualityGateStatus,omitempty"`
		} `json:"status,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"branches,omitempty"`
}

// RenameRequest Rename the main branch of a project or application.<br/>Requires 'Administer' permission on the specified project or application.
type RenameRequest struct {
	Name    string `form:"name"`    // New name of the main branch
	Project string `form:"project"` // Project key
}

// SetAutomaticDeletionProtectionRequest Protect a specific branch from automatic deletion. Protection can't be disabled for the main branch.<br/>Requires 'Administer' permission on the specified project or application.
type SetAutomaticDeletionProtectionRequest struct {
	Branch  string `form:"branch"`  // Branch key
	Project string `form:"project"` // Project key
	Value   string `form:"value"`   // Sets whether the branch should be protected from automatic deletion when it hasn't been analyzed for a set period of time.
}

// SetMainRequest Allow to set a new main branch.<br/>. Caution, only applicable on projects.<br>Requires 'Administer' rights on the specified project or application.
type SetMainRequest struct {
	Branch  string `form:"branch"`  // Branch key
	Project string `form:"project"` // Project key
}
