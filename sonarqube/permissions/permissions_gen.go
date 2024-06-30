package permissions

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddGroupRequest Add a permission to a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name must be provided. <br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
type AddGroupRequest struct {
	GroupName  string `json:"groupName"`            // Group name or 'anyone' (case insensitive)
	Permission string `json:"permission"`           // The permission you would like to grant to the group.<ul><li>Possible values for global permissions: admin, gateadmin, profileadmin, provisioning, scan, applicationcreator, portfoliocreator</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `json:"projectId,omitempty"`  // Project id
	ProjectKey string `json:"projectKey,omitempty"` // Project key
}

// AddGroupToTemplateRequest Add a group to a permission template.<br /> The group name must be provided. <br />Requires the following permission: 'Administer System'.
type AddGroupToTemplateRequest struct {
	GroupName    string `json:"groupName"`              // Group name or 'anyone' (case insensitive)
	Permission   string `json:"permission"`             // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// AddProjectCreatorToTemplateRequest Add a project creator to a permission template.<br>Requires the following permission: 'Administer System'.
type AddProjectCreatorToTemplateRequest struct {
	Permission   string `json:"permission"`             // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// AddUserRequest Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
type AddUserRequest struct {
	Login      string `json:"login"`                // User login
	Permission string `json:"permission"`           // The permission you would like to grant to the user<ul><li>Possible values for global permissions: admin, gateadmin, profileadmin, provisioning, scan, applicationcreator, portfoliocreator</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `json:"projectId,omitempty"`  // Project id
	ProjectKey string `json:"projectKey,omitempty"` // Project key
}

// AddUserToTemplateRequest Add a user to a permission template.<br /> Requires the following permission: 'Administer System'.
type AddUserToTemplateRequest struct {
	Login        string `json:"login"`                  // User login
	Permission   string `json:"permission"`             // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// ApplyTemplateRequest Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the following permission: 'Administer System'.
type ApplyTemplateRequest struct {
	ProjectId    string `json:"projectId,omitempty"`    // Project id
	ProjectKey   string `json:"projectKey,omitempty"`   // Project key
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// BulkApplyTemplateRequest Apply a permission template to several components. Managed projects will be ignored.<br />The template id or name must be provided.<br />Requires the following permission: 'Administer System'.
type BulkApplyTemplateRequest struct {
	AnalyzedBefore    string `json:"analyzedBefore,omitempty"`    // Since 6.6;Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.
	OnProvisionedOnly string `json:"onProvisionedOnly,omitempty"` // Since 6.6;Filter the projects that are provisioned
	Projects          string `json:"projects,omitempty"`          // Since 6.6;Comma-separated list of project keys
	Q                 string `json:"q,omitempty"`                 // Limit search to: <ul><li>project names that contain the supplied string</li><li>project keys that are exactly the same as the supplied string</li></ul>
	Qualifiers        string `json:"qualifiers,omitempty"`        // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>APP - Applications</li><li>TRK - Projects</li><li>VW - Portfolios</li></ul>
	TemplateId        string `json:"templateId,omitempty"`        // Template id
	TemplateName      string `json:"templateName,omitempty"`      // Template name
}

// CreateTemplateRequest Create a permission template.<br />Requires the following permission: 'Administer System'.
type CreateTemplateRequest struct {
	Description       string `json:"description,omitempty"`       // Description
	Name              string `json:"name"`                        // Name
	ProjectKeyPattern string `json:"projectKeyPattern,omitempty"` // Project key pattern. Must be a valid Java regular expression
}

// CreateTemplateResponse is the response for CreateTemplateRequest
type CreateTemplateResponse struct {
	PermissionTemplate struct {
		Description       string `json:"description,omitempty"`
		Name              string `json:"name,omitempty"`
		ProjectKeyPattern string `json:"projectKeyPattern,omitempty"`
	} `json:"permissionTemplate,omitempty"`
}

// DeleteTemplateRequest Delete a permission template.<br />Requires the following permission: 'Administer System'.
type DeleteTemplateRequest struct {
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// RemoveGroupRequest Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name must be provided.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
type RemoveGroupRequest struct {
	GroupName  string `json:"groupName"`            // Group name or 'anyone' (case insensitive)
	Permission string `json:"permission"`           // The permission you would like to revoke from the group.<ul><li>Possible values for global permissions: admin, gateadmin, profileadmin, provisioning, scan, applicationcreator, portfoliocreator</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `json:"projectId,omitempty"`  // Project id
	ProjectKey string `json:"projectKey,omitempty"` // Project key
}

// RemoveGroupFromTemplateRequest Remove a group from a permission template.<br /> The group name must be provided. <br />Requires the following permission: 'Administer System'.
type RemoveGroupFromTemplateRequest struct {
	GroupName    string `json:"groupName"`              // Group name or 'anyone' (case insensitive)
	Permission   string `json:"permission"`             // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// RemoveProjectCreatorFromTemplateRequest Remove a project creator from a permission template.<br>Requires the following permission: 'Administer System'.
type RemoveProjectCreatorFromTemplateRequest struct {
	Permission   string `json:"permission"`             // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// RemoveUserRequest Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
type RemoveUserRequest struct {
	Login      string `json:"login"`                // User login
	Permission string `json:"permission"`           // The permission you would like to revoke from the user.<ul><li>Possible values for global permissions: admin, gateadmin, profileadmin, provisioning, scan, applicationcreator, portfoliocreator</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `json:"projectId,omitempty"`  // Project id
	ProjectKey string `json:"projectKey,omitempty"` // Project key
}

// RemoveUserFromTemplateRequest Remove a user from a permission template.<br /> Requires the following permission: 'Administer System'.
type RemoveUserFromTemplateRequest struct {
	Login        string `json:"login"`                  // User login
	Permission   string `json:"permission"`             // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// SearchTemplatesRequest List permission templates.<br />Requires the following permission: 'Administer System'.
type SearchTemplatesRequest struct {
	Q string `url:"q,omitempty"` // Limit search to permission template names that contain the supplied string.
}

// SearchTemplatesResponse is the response for SearchTemplatesRequest
type SearchTemplatesResponse struct {
	DefaultTemplates []struct {
		Qualifier  string `json:"qualifier,omitempty"`
		TemplateId string `json:"templateId,omitempty"`
	} `json:"defaultTemplates,omitempty"`
	PermissionTemplates []struct {
		CreatedAt   string `json:"createdAt,omitempty"`
		Description string `json:"description,omitempty"`
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Permissions []struct {
			GroupsCount        float64 `json:"groupsCount,omitempty"`
			Key                string  `json:"key,omitempty"`
			UsersCount         float64 `json:"usersCount,omitempty"`
			WithProjectCreator bool    `json:"withProjectCreator,omitempty"`
		} `json:"permissions,omitempty"`
		UpdatedAt string `json:"updatedAt,omitempty"`
	} `json:"permissionTemplates,omitempty"`
}

// SetDefaultTemplateRequest Set a permission template as default.<br />Requires the following permission: 'Administer System'.
type SetDefaultTemplateRequest struct {
	Qualifier    string `json:"qualifier,omitempty"`    // Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>APP - Applications</li><li>TRK - Projects</li><li>VW - Portfolios</li></ul>
	TemplateId   string `json:"templateId,omitempty"`   // Template id
	TemplateName string `json:"templateName,omitempty"` // Template name
}

// UpdateTemplateRequest Update a permission template.<br />Requires the following permission: 'Administer System'.
type UpdateTemplateRequest struct {
	Description       string `json:"description,omitempty"`       // Description
	Id                string `json:"id"`                          // Id
	Name              string `json:"name,omitempty"`              // Name
	ProjectKeyPattern string `json:"projectKeyPattern,omitempty"` // Project key pattern. Must be a valid Java regular expression
}

// UpdateTemplateResponse is the response for UpdateTemplateRequest
type UpdateTemplateResponse struct {
	PermissionTemplate struct {
		CreatedAt         string `json:"createdAt,omitempty"`
		Description       string `json:"description,omitempty"`
		Id                string `json:"id,omitempty"`
		Name              string `json:"name,omitempty"`
		ProjectKeyPattern string `json:"projectKeyPattern,omitempty"`
		UpdatedAt         string `json:"updatedAt,omitempty"`
	} `json:"permissionTemplate,omitempty"`
}
