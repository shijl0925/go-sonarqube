package user_groups

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddUserRequest Add a user to a group.<br />'name' must be provided.<br />Requires the following permission: 'Administer System'.
// Deprecated: this action has been deprecated since version 10.4
type AddUserRequest struct {
	Login string `form:"login,omitempty"` // User login
	Name  string `form:"name,omitempty"`  // Group name
}

// CreateRequest Create a group.<br>Requires the following permission: 'Administer System'.
// Deprecated: this action has been deprecated since version 10.4
type CreateRequest struct {
	Description string `form:"description,omitempty"` // Description for the new group. A group description cannot be larger than 200 characters.
	Name        string `form:"name,omitempty"`        // Name for the new group. A group name cannot be larger than 255 characters and must be unique. The value 'anyone' (whatever the case) is reserved and cannot be used.
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Group struct {
		Default      bool    `json:"default,omitempty"`
		Description  string  `json:"description,omitempty"`
		Id           float64 `json:"id,omitempty"`
		MembersCount float64 `json:"membersCount,omitempty"`
		Name         string  `json:"name,omitempty"`
		Organization string  `json:"organization,omitempty"`
	} `json:"group,omitempty"`
}

// DeleteRequest Delete a group. The default groups cannot be deleted.<br/>'name' must be provided.<br />Requires the following permission: 'Administer System'.
// Deprecated: this action has been deprecated since version 10.4
type DeleteRequest struct {
	Name string `form:"name,omitempty"` // Group name
}

// RemoveUserRequest Remove a user from a group.<br />'name' must be provided.<br>Requires the following permission: 'Administer System'.
// Deprecated: this action has been deprecated since version 10.4
type RemoveUserRequest struct {
	Login string `form:"login,omitempty"` // User login
	Name  string `form:"name,omitempty"`  // Group name
}

// SearchRequest Search for user groups.<br>Requires the following permission: 'Administer System'.
// Deprecated: this action has been deprecated since version 10.4
type SearchRequest struct {
	F       string `form:"f,omitempty"`       // Comma-separated list of the fields to be returned in response. All the fields are returned by default.
	Managed string `form:"managed,omitempty"` // Return managed or non-managed groups. Only available for managed instances, throws for non-managed instances.
	Q       string `form:"q,omitempty"`       // Limit search to names that contain the supplied string.
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Groups []struct {
		Default      bool    `json:"default,omitempty"`
		Description  string  `json:"description,omitempty"`
		Managed      bool    `json:"managed,omitempty"`
		MembersCount float64 `json:"membersCount,omitempty"`
		Name         string  `json:"name,omitempty"`
	} `json:"groups,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Groups []struct {
		Default      bool    `json:"default,omitempty"`
		Description  string  `json:"description,omitempty"`
		Managed      bool    `json:"managed,omitempty"`
		MembersCount float64 `json:"membersCount,omitempty"`
		Name         string  `json:"name,omitempty"`
	} `json:"groups,omitempty"`
}

// UpdateRequest Update a group.<br>Requires the following permission: 'Administer System'.
// Deprecated: this action has been deprecated since version 10.4
type UpdateRequest struct {
	CurrentName string `form:"currentName,omitempty"` // Name of the group to be updated.
	Description string `form:"description,omitempty"` // New optional description for the group. A group description cannot be larger than 200 characters. If value is not defined, then description is not changed.
	Name        string `form:"name,omitempty"`        // New optional name for the group. A group name cannot be larger than 255 characters and must be unique. Value 'anyone' (whatever the case) is reserved and cannot be used. If value is empty or not defined, then name is not changed.
}

// UsersRequest Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'.
// Deprecated: this action has been deprecated since version 10.4
type UsersRequest struct {
	Name     string `form:"name,omitempty"`     // Group name
	Q        string `form:"q,omitempty"`        // Limit search to names or logins that contain the supplied string.
	Selected string `form:"selected,omitempty"` // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// UsersResponse is the response for UsersRequest
type UsersResponse struct {
	Paging paging.Paging `json:"paging,omitempty"`
	Users  []struct {
		Login    string `json:"login,omitempty"`
		Managed  bool   `json:"managed,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"users,omitempty"`
}

// GetPaging extracts the paging from UsersResponse
func (r *UsersResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// UsersResponseAll is the collection for UsersRequest
type UsersResponseAll struct {
	Users []struct {
		Login    string `json:"login,omitempty"`
		Managed  bool   `json:"managed,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"users,omitempty"`
}
