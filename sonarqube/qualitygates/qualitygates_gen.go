package qualitygates

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddGroupRequest Allow a group of users to edit a Quality Gate.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Gates'</li>  <li>Edit right on the specified quality gate</li></ul>
type AddGroupRequest struct {
	GateName  string `json:"gateName"`  // Quality Gate name
	GroupName string `json:"groupName"` // Group name or 'anyone' (case insensitive)
}

// AddUserRequest Allow a user to edit a Quality Gate.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Gates'</li>  <li>Edit right on the specified quality gate</li></ul>
type AddUserRequest struct {
	GateName string `json:"gateName"` // Quality Gate name
	Login    string `json:"login"`    // User login
}

// CopyRequest Copy a Quality Gate.<br>'sourceName' must be provided. Requires the 'Administer Quality Gates' permission.
type CopyRequest struct {
	Name       string `json:"name"`       // The name of the quality gate to create
	SourceName string `json:"sourceName"` // Since 8.4;The name of the quality gate to copy
}

// CreateRequest Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
type CreateRequest struct {
	Name string `json:"name"` // The name of the quality gate to create
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// CreateConditionRequest Add a new condition to a quality gate.<br>Parameter 'gateName' must be provided. Requires the 'Administer Quality Gates' permission.
type CreateConditionRequest struct {
	Error    string `json:"error"`        // Condition error threshold
	GateName string `json:"gateName"`     // Name of the quality gate
	Metric   string `json:"metric"`       // Condition metric.<br/> Only metric of the following types are allowed:<ul><li>INT</li><li>MILLISEC</li><li>RATING</li><li>WORK_DUR</li><li>FLOAT</li><li>PERCENT</li><li>LEVEL</li></ul>Following metrics are forbidden:<ul><li>alert_status</li><li>security_hotspots</li><li>new_security_hotspots</li></ul>
	Op       string `json:"op,omitempty"` // Condition operator:<br/><ul><li>LT = is lower than</li><li>GT = is greater than</li></ul>
}

// CreateConditionResponse is the response for CreateConditionRequest
type CreateConditionResponse struct {
	Error   string `json:"error,omitempty"`
	Id      string `json:"id,omitempty"`
	Metric  string `json:"metric,omitempty"`
	Op      string `json:"op,omitempty"`
	Warning string `json:"warning,omitempty"`
}

// DeleteConditionRequest Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission.
type DeleteConditionRequest struct {
	Id string `json:"id"` // Condition UUID
}

// DeselectRequest Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul>
type DeselectRequest struct {
	ProjectKey string `json:"projectKey"` // Since 6.1;Project key
}

// DestroyRequest Delete a Quality Gate.<br>Parameter 'name' must be specified. Requires the 'Administer Quality Gates' permission.
type DestroyRequest struct {
	Name string `json:"name"` // Since 8.4;Name of the quality gate to delete
}

// GetByProjectRequest Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
type GetByProjectRequest struct {
	Project string `url:"project"` // Project key
}

// GetByProjectResponse is the response for GetByProjectRequest
type GetByProjectResponse struct {
	QualityGate struct {
		Default bool   `json:"default,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"qualityGate,omitempty"`
}

// ListRequest Get a list of quality gates
type ListRequest struct{}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Actions struct {
		Create bool `json:"create,omitempty"`
	} `json:"actions,omitempty"`
	Qualitygates []struct {
		Actions struct {
			AssociateProjects bool `json:"associateProjects,omitempty"`
			Copy              bool `json:"copy,omitempty"`
			Delegate          bool `json:"delegate,omitempty"`
			Delete            bool `json:"delete,omitempty"`
			ManageConditions  bool `json:"manageConditions,omitempty"`
			Rename            bool `json:"rename,omitempty"`
			SetAsDefault      bool `json:"setAsDefault,omitempty"`
		} `json:"actions,omitempty"`
		CaycStatus string `json:"caycStatus,omitempty"`
		IsBuiltIn  bool   `json:"isBuiltIn,omitempty"`
		IsDefault  bool   `json:"isDefault,omitempty"`
		Name       string `json:"name,omitempty"`
	} `json:"qualitygates,omitempty"`
}

// ProjectStatusRequest Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided <br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li><li>'Execute Analysis' on the specified project</li></ul>
type ProjectStatusRequest struct {
	AnalysisId  string `url:"analysisId,omitempty"`  // Analysis id
	Branch      string `url:"branch,omitempty"`      // Since 7.7;Branch key
	ProjectId   string `url:"projectId,omitempty"`   // Since 5.4;Project UUID. Doesn't work with branches or pull requests
	ProjectKey  string `url:"projectKey,omitempty"`  // Since 5.4;Project key
	PullRequest string `url:"pullRequest,omitempty"` // Since 7.7;Pull request id
}

// ProjectStatusResponse is the response for ProjectStatusRequest
type ProjectStatusResponse struct {
	ProjectStatus struct {
		CaycStatus string `json:"caycStatus,omitempty"`
		Conditions []struct {
			ActualValue    string `json:"actualValue,omitempty"`
			Comparator     string `json:"comparator,omitempty"`
			ErrorThreshold string `json:"errorThreshold,omitempty"`
			MetricKey      string `json:"metricKey,omitempty"`
			Status         string `json:"status,omitempty"`
		} `json:"conditions,omitempty"`
		IgnoredConditions bool `json:"ignoredConditions,omitempty"`
		Period            struct {
			Date      string `json:"date,omitempty"`
			Mode      string `json:"mode,omitempty"`
			Parameter string `json:"parameter,omitempty"`
		} `json:"period,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"projectStatus,omitempty"`
}

// RemoveGroupRequest Remove the ability from a group to edit a Quality Gate.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Gates'</li>  <li>Edit right on the specified quality gate</li></ul>
type RemoveGroupRequest struct {
	GateName  string `json:"gateName"`  // Quality Gate name
	GroupName string `json:"groupName"` // Group name or 'anyone' (case insensitive)
}

// RemoveUserRequest Remove the ability from an user to edit a Quality Gate.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Gates'</li>  <li>Edit right on the specified quality gate</li></ul>
type RemoveUserRequest struct {
	GateName string `json:"gateName"` // Quality Gate name
	Login    string `json:"login"`    // User login
}

// RenameRequest Rename a Quality Gate.<br>'currentName' must be specified. Requires the 'Administer Quality Gates' permission.
type RenameRequest struct {
	CurrentName string `json:"currentName"` // Since 8.4;Current name of the quality gate
	Name        string `json:"name"`        // New name of the quality gate
}

// SearchRequest Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for the current user will be returned.
type SearchRequest struct {
	GateName string `url:"gateName"`           // Since 8.4;Quality Gate name
	Page     string `url:"page,omitempty"`     // Page number
	PageSize string `url:"pageSize,omitempty"` // Page size
	Query    string `url:"query,omitempty"`    // To search for projects containing this string. If this parameter is set, "selected" is set to "all".
	Selected string `url:"selected,omitempty"` // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Paging struct {
		PageIndex float64 `json:"pageIndex,omitempty"`
		PageSize  float64 `json:"pageSize,omitempty"`
		Total     float64 `json:"total,omitempty"`
	} `json:"paging,omitempty"`
	Results []struct {
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"results,omitempty"`
}

// SearchGroupsRequest List the groups that are allowed to edit a Quality Gate.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Gates'</li>  <li>Edit right on the specified quality gate</li></ul>
type SearchGroupsRequest struct {
	GateName string `url:"gateName"`           // Quality Gate name
	Q        string `url:"q,omitempty"`        // Limit search to group names that contain the supplied string.
	Selected string `url:"selected,omitempty"` // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// SearchGroupsResponse is the response for SearchGroupsRequest
type SearchGroupsResponse struct {
	Groups []struct {
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Selected    bool   `json:"selected,omitempty"`
	} `json:"groups,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchGroupsResponse
func (r *SearchGroupsResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchGroupsResponseAll is the collection for SearchGroupsRequest
type SearchGroupsResponseAll struct {
	Groups []struct {
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Selected    bool   `json:"selected,omitempty"`
	} `json:"groups,omitempty"`
}

// SearchUsersRequest List the users that are allowed to edit a Quality Gate.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Gates'</li>  <li>Edit right on the specified quality gate</li></ul>
type SearchUsersRequest struct {
	GateName string `url:"gateName"`           // Quality Gate name
	Q        string `url:"q,omitempty"`        // Limit search to names or logins that contain the supplied string.
	Selected string `url:"selected,omitempty"` // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// SearchUsersResponse is the response for SearchUsersRequest
type SearchUsersResponse struct {
	Paging paging.Paging `json:"paging,omitempty"`
	Users  []struct {
		Avatar   string `json:"avatar,omitempty"`
		Login    string `json:"login,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"users,omitempty"`
}

// GetPaging extracts the paging from SearchUsersResponse
func (r *SearchUsersResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchUsersResponseAll is the collection for SearchUsersRequest
type SearchUsersResponseAll struct {
	Users []struct {
		Avatar   string `json:"avatar,omitempty"`
		Login    string `json:"login,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"users,omitempty"`
}

// SelectRequest Associate a project to a quality gate.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Gates'</li>  <li>'Administer' right on the specified project</li></ul>
type SelectRequest struct {
	GateName   string `json:"gateName"`   // Since 8.4;Name of the quality gate
	ProjectKey string `json:"projectKey"` // Since 6.1;Project key
}

// SetAsDefaultRequest Set a quality gate as the default quality gate.<br>Parameter 'name' must be specified. Requires the 'Administer Quality Gates' permission.
type SetAsDefaultRequest struct {
	Name string `json:"name"` // Since 8.4;Name of the quality gate to set as default
}

// ShowRequest Display the details of a quality gate
type ShowRequest struct {
	Name string `url:"name"` // Name of the quality gate. Either id or name must be set
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Actions struct {
		AssociateProjects bool `json:"associateProjects,omitempty"`
		Copy              bool `json:"copy,omitempty"`
		Delegate          bool `json:"delegate,omitempty"`
		Delete            bool `json:"delete,omitempty"`
		ManageConditions  bool `json:"manageConditions,omitempty"`
		Rename            bool `json:"rename,omitempty"`
		SetAsDefault      bool `json:"setAsDefault,omitempty"`
	} `json:"actions,omitempty"`
	CaycStatus string `json:"caycStatus,omitempty"`
	Conditions []struct {
		Error  string `json:"error,omitempty"`
		Id     string `json:"id,omitempty"`
		Metric string `json:"metric,omitempty"`
		Op     string `json:"op,omitempty"`
	} `json:"conditions,omitempty"`
	IsBuiltIn bool   `json:"isBuiltIn,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
	Name      string `json:"name,omitempty"`
}

// UpdateConditionRequest Update a condition attached to a quality gate.<br>Requires the 'Administer Quality Gates' permission.
type UpdateConditionRequest struct {
	Error  string `json:"error"`        // Condition error threshold
	Id     string `json:"id"`           // Condition ID
	Metric string `json:"metric"`       // Condition metric.<br/> Only metric of the following types are allowed:<ul><li>INT</li><li>MILLISEC</li><li>RATING</li><li>WORK_DUR</li><li>FLOAT</li><li>PERCENT</li><li>LEVEL</li></ul>Following metrics are forbidden:<ul><li>alert_status</li><li>security_hotspots</li><li>new_security_hotspots</li></ul>
	Op     string `json:"op,omitempty"` // Condition operator:<br/><ul><li>LT = is lower than</li><li>GT = is greater than</li></ul>
}
