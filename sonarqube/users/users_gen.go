package users

import paging "github.com/shijl0925/go-sonarqube/sonarqube/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AnonymizeRequest Anonymize a deactivated user. Requires Administer System permission
// Deprecated: this action has been deprecated since version 10.4
type AnonymizeRequest struct {
	Login string `json:"login"` // User login
}

// ChangePasswordRequest Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password.
type ChangePasswordRequest struct {
	Login            string `json:"login"`                      // User login
	Password         string `json:"password"`                   // New password
	PreviousPassword string `json:"previousPassword,omitempty"` // Previous password. Required when changing one's own password.
}

// CreateRequest Create a user.<br/>If a deactivated user account exists with the given login, it will be reactivated.<br/>Requires Administer System permission
// Deprecated: this action has been deprecated since version 10.4
type CreateRequest struct {
	Email      string `json:"email,omitempty"`      // User email
	Local      string `json:"local,omitempty"`      // Since 6.3;Specify if the user should be authenticated from SonarQube server or from an external authentication system. Password should not be set when local is set to false.
	Login      string `json:"login"`                // User login
	Name       string `json:"name"`                 // User name
	Password   string `json:"password,omitempty"`   // User password. Only mandatory when creating local user, otherwise it should not be set
	ScmAccount string `json:"scmAccount,omitempty"` // List of SCM accounts. To set several values, the parameter must be called once for each value.
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	User struct {
		Active      bool     `json:"active,omitempty"`
		Email       string   `json:"email,omitempty"`
		Local       bool     `json:"local,omitempty"`
		Login       string   `json:"login,omitempty"`
		Name        string   `json:"name,omitempty"`
		ScmAccounts []string `json:"scmAccounts,omitempty"`
	} `json:"user,omitempty"`
}

// DeactivateRequest Deactivate a user. Requires Administer System permission
// Deprecated: this action has been deprecated since version 10.4
type DeactivateRequest struct {
	Anonymize string `json:"anonymize,omitempty"` // Since 9.7;Anonymize user in addition to deactivating it
	Login     string `json:"login"`               // User login
}

// DeactivateResponse is the response for DeactivateRequest
type DeactivateResponse struct {
	User struct {
		Active      bool     `json:"active,omitempty"`
		Groups      []string `json:"groups,omitempty"`
		Local       bool     `json:"local,omitempty"`
		Login       string   `json:"login,omitempty"`
		Name        string   `json:"name,omitempty"`
		ScmAccounts []string `json:"scmAccounts,omitempty"`
	} `json:"user,omitempty"`
}

// GroupsRequest Lists the groups a user belongs to. <br/>Requires Administer System permission.
// Deprecated: this action has been deprecated since version 10.4
type GroupsRequest struct {
	Login    string `url:"login"`              // A user login
	Q        string `url:"q,omitempty"`        // Limit search to group names that contain the supplied string.
	Selected string `url:"selected,omitempty"` // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// GroupsResponse is the response for GroupsRequest
type GroupsResponse struct {
	Groups []struct {
		Default     bool    `json:"default,omitempty"`
		Description string  `json:"description,omitempty"`
		Id          float64 `json:"id,omitempty"`
		Name        string  `json:"name,omitempty"`
		Selected    bool    `json:"selected,omitempty"`
	} `json:"groups,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from GroupsResponse
func (r *GroupsResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// GroupsResponseAll is the collection for GroupsRequest
type GroupsResponseAll struct {
	Groups []struct {
		Default     bool    `json:"default,omitempty"`
		Description string  `json:"description,omitempty"`
		Id          float64 `json:"id,omitempty"`
		Name        string  `json:"name,omitempty"`
		Selected    bool    `json:"selected,omitempty"`
	} `json:"groups,omitempty"`
}

// SearchRequest Get a list of users. By default, only active users are returned.<br/>The following fields are only returned when user has Administer System permission or for logged-in in user :<ul>   <li>'email'</li>   <li>'externalIdentity'</li>   <li>'externalProvider'</li>   <li>'groups'</li>   <li>'lastConnectionDate'</li>   <li>'sonarLintLastConnectionDate'</li>   <li>'tokensCount'</li></ul>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user authenticates many times in less than one hour.
// Deprecated: this action has been deprecated since version 10.4
type SearchRequest struct {
	Deactivated      string `url:"deactivated,omitempty"`      // Since 9.7;Return deactivated users instead of active users
	ExternalIdentity string `url:"externalIdentity,omitempty"` /*
	Since 10.3;Find a user by its external identity (ie. its login in the Identity Provider).
	This is case sensitive and only available with Administer System permission.
	*/
	LastConnectedAfter string `url:"lastConnectedAfter,omitempty"` /*
	Since 10.1;Filter the users based on the last connection date field.
	Only users who interacted with this instance at or after the date will be returned.
	The format must be ISO 8601 datetime format (YYYY-MM-DDThh:mm:ss±hhmm)
	*/
	LastConnectedBefore string `url:"lastConnectedBefore,omitempty"` /*
	Since 10.1;Filter the users based on the last connection date field.
	Only users that never connected or who interacted with this instance at or before the date will be returned.
	The format must be ISO 8601 datetime format (YYYY-MM-DDThh:mm:ss±hhmm)
	*/
	Managed              string `url:"managed,omitempty"`              // Since 10.0;Return managed or non-managed users. Only available for managed instances, throws for non-managed instances.
	Q                    string `url:"q,omitempty"`                    // Filter on login, name and email.<br />This parameter can either perform an exact match, or a partial match (contains), it is case insensitive.
	SlLastConnectedAfter string `url:"slLastConnectedAfter,omitempty"` /*
	Since 10.1;Filter the users based on the sonar lint last connection date field
	Only users who interacted with this instance using SonarLint at or after the date will be returned.
	The format must be ISO 8601 datetime format (YYYY-MM-DDThh:mm:ss±hhmm)
	*/
	SlLastConnectedBefore string `url:"slLastConnectedBefore,omitempty"` /*
	Since 10.1;Filter the users based on the sonar lint last connection date field.
	Only users that never connected or who interacted with this instance using SonarLint at or before the date will be returned.
	The format must be ISO 8601 datetime format (YYYY-MM-DDThh:mm:ss±hhmm)
	*/
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Paging paging.Paging `json:"paging,omitempty"`
	Users  []struct {
		Active                      bool     `json:"active,omitempty"`
		Avatar                      string   `json:"avatar,omitempty"`
		Email                       string   `json:"email,omitempty"`
		ExternalIdentity            string   `json:"externalIdentity,omitempty"`
		ExternalProvider            string   `json:"externalProvider,omitempty"`
		Groups                      []string `json:"groups,omitempty"`
		LastConnectionDate          string   `json:"lastConnectionDate,omitempty"`
		Local                       bool     `json:"local,omitempty"`
		Login                       string   `json:"login,omitempty"`
		Managed                     bool     `json:"managed,omitempty"`
		Name                        string   `json:"name,omitempty"`
		SonarLintLastConnectionDate string   `json:"sonarLintLastConnectionDate,omitempty"`
		TokensCount                 float64  `json:"tokensCount,omitempty"`
		ScmAccounts                 []string `json:"scmAccounts,omitempty"`
	} `json:"users,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Users []struct {
		Active                      bool     `json:"active,omitempty"`
		Avatar                      string   `json:"avatar,omitempty"`
		Email                       string   `json:"email,omitempty"`
		ExternalIdentity            string   `json:"externalIdentity,omitempty"`
		ExternalProvider            string   `json:"externalProvider,omitempty"`
		Groups                      []string `json:"groups,omitempty"`
		LastConnectionDate          string   `json:"lastConnectionDate,omitempty"`
		Local                       bool     `json:"local,omitempty"`
		Login                       string   `json:"login,omitempty"`
		Managed                     bool     `json:"managed,omitempty"`
		Name                        string   `json:"name,omitempty"`
		SonarLintLastConnectionDate string   `json:"sonarLintLastConnectionDate,omitempty"`
		TokensCount                 float64  `json:"tokensCount,omitempty"`
		ScmAccounts                 []string `json:"scmAccounts,omitempty"`
	} `json:"users,omitempty"`
}

// UpdateRequest Update a user.<br/>Requires Administer System permission
// Deprecated: this action has been deprecated since version 10.4
type UpdateRequest struct {
	Email      string `json:"email,omitempty"`      // User email
	Login      string `json:"login"`                // User login
	Name       string `json:"name,omitempty"`       // User name
	ScmAccount string `json:"scmAccount,omitempty"` // SCM accounts. To set several values, the parameter must be called once for each value.
}

// UpdateResponse is the response for UpdateRequest
type UpdateResponse struct {
	User struct {
		Active      bool     `json:"active,omitempty"`
		Email       string   `json:"email,omitempty"`
		Local       bool     `json:"local,omitempty"`
		Login       string   `json:"login,omitempty"`
		Name        string   `json:"name,omitempty"`
		ScmAccounts []string `json:"scmAccounts,omitempty"`
	} `json:"user,omitempty"`
}

// UpdateIdentityProviderRequest Update identity provider information. <br/>It's only possible to migrate to an installed identity provider. Be careful that as soon as this information has been updated for a user, the user will only be able to authenticate on the new identity provider. It is not possible to migrate external user to local one.<br/>Requires Administer System permission.
// Deprecated: this action has been deprecated since version 10.4
type UpdateIdentityProviderRequest struct {
	Login               string `json:"login"`                         // User login
	NewExternalIdentity string `json:"newExternalIdentity,omitempty"` // New external identity, usually the login used in the authentication system. If not provided previous identity will be used.
	NewExternalProvider string `json:"newExternalProvider"`           // New external provider. Only authentication system installed are available. Use 'LDAP' identity provider for single server LDAP setup.Use 'LDAP_{serverKey}' identity provider for multiple LDAP servers setup.
}

// UpdateLoginRequest Update a user login. A login can be updated many times.<br/>Requires Administer System permission
// Deprecated: this action has been deprecated since version 10.4
type UpdateLoginRequest struct {
	Login    string `json:"login"`    // The current login (case-sensitive)
	NewLogin string `json:"newLogin"` // The new login. It must not already exist.
}
