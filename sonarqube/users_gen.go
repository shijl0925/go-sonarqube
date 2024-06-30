package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/users"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Users service

// Anonymize - Anonymize a deactivated user. Requires Administer System permission
// Since 9.7
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use DELETE api/v2/users-management/users/{id}?anonymize=true instead
func (s *Users) Anonymize(ctx context.Context, r users.AnonymizeRequest) error {
	u := fmt.Sprintf("%s/users/anonymize", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// ChangePassword - Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password.
// Since 5.2
// Changelog:
//   8.6: It's no more possible for the password to be the same as the previous one
func (s *Users) ChangePassword(ctx context.Context, r users.ChangePasswordRequest) error {
	u := fmt.Sprintf("%s/users/change_password", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Create - Create a user.
// If a deactivated user account exists with the given login, it will be reactivated.
// Requires Administer System permission
// Since 3.7
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use POST api/v2/users-management/users instead
//   6.3: The password is only mandatory when creating local users, and should not be set on non local users
//   6.3: The 'infos' message is no more returned when a user is reactivated
func (s *Users) Create(ctx context.Context, r users.CreateRequest) (*users.CreateResponse, error) {
	u := fmt.Sprintf("%s/users/create", API)
	v := new(users.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Deactivate - Deactivate a user. Requires Administer System permission
// Since 3.7
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use DELETE api/v2/users-management/users/{id} instead
func (s *Users) Deactivate(ctx context.Context, r users.DeactivateRequest) (*users.DeactivateResponse, error) {
	u := fmt.Sprintf("%s/users/deactivate", API)
	v := new(users.DeactivateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Groups - Lists the groups a user belongs to.
// Requires Administer System permission.
// Since 5.2
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use GET api/v2/authorizations/groups-memberships?userId={} instead
//   6.4: Paging response fields moved to a Paging object
//   6.4: 'default' response field has been added
func (s *Users) Groups(ctx context.Context, r users.GroupsRequest, p paging.Params) (*users.GroupsResponse, error) {
	u := fmt.Sprintf("%s/users/groups", API)
	v := new(users.GroupsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Users) GroupsAll(ctx context.Context, r users.GroupsRequest) (*users.GroupsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &users.GroupsResponseAll{}
	for {
		res, err := s.Groups(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to users.Groups: %+v", err)
		}
		response.Groups = append(response.Groups, res.Groups...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// Search - Get a list of users. By default, only active users are returned.
// The following fields are only returned when user has Administer System permission or for logged-in in user :
//     * 'email'
//     * 'externalIdentity'
//     * 'externalProvider'
//     * 'groups'
//     * 'lastConnectionDate'
//     * 'sonarLintLastConnectionDate'
//     * 'tokensCount'
// Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user authenticates many times in less than one hour.
// Since 3.6
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use GET api/v2/users-management/users instead
//   10.3: New optional parameters externalIdentity to find a user by its IdP login
//   10.1: New optional parameters slLastConnectedAfter and slLastConnectedBefore to filter users by SonarLint last connection date. Only available with Administer System permission.
//   10.1: New optional parameters lastConnectedAfter and lastConnectedBefore to filter users by SonarQube last connection date. Only available with Administer System permission.
//   10.1: New field 'sonarLintLastConnectionDate' is added to response
//   10.0: 'q' parameter values is now always performing a case insensitive match
//   10.0: New parameter 'managed' to optionally search by managed status
//   10.0: Response includes 'managed' field.
//   9.7: New parameter 'deactivated' to optionally search for deactivated users
//   7.7: New field 'lastConnectionDate' is added to response
//   7.4: External identity is only returned to system administrators
//   6.4: Paging response fields moved to a Paging object
//   6.4: Avatar has been added to the response
//   6.4: Email is only returned when user has Administer System permission
func (s *Users) Search(ctx context.Context, r users.SearchRequest, p paging.Params) (*users.SearchResponse, error) {
	u := fmt.Sprintf("%s/users/search", API)
	v := new(users.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Users) SearchAll(ctx context.Context, r users.SearchRequest) (*users.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &users.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to users.Search: %+v", err)
		}
		response.Users = append(response.Users, res.Users...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// Update - Update a user.
// Requires Administer System permission
// Since 3.7
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use PATCH api/v2/users-management/users/{id} instead
//   5.2: User's password can only be changed using the 'change_password' action.
func (s *Users) Update(ctx context.Context, r users.UpdateRequest) (*users.UpdateResponse, error) {
	u := fmt.Sprintf("%s/users/update", API)
	v := new(users.UpdateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// UpdateIdentityProvider - Update identity provider information.
// It's only possible to migrate to an installed identity provider. Be careful that as soon as this information has been updated for a user, the user will only be able to authenticate on the new identity provider. It is not possible to migrate external user to local one.
// Requires Administer System permission.
// Since 8.7
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use PATCH api/v2/users-management/users/{id} instead
//   9.8: Use of 'sonarqube' for the value of 'newExternalProvider' is deprecated.
func (s *Users) UpdateIdentityProvider(ctx context.Context, r users.UpdateIdentityProviderRequest) error {
	u := fmt.Sprintf("%s/users/update_identity_provider", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLogin - Update a user login. A login can be updated many times.
// Requires Administer System permission
// Since 7.6
// Deprecated since 10.4
// Changelog:
//   10.4: Deprecated. Use PATCH api/v2/users-management/users/{id} instead
func (s *Users) UpdateLogin(ctx context.Context, r users.UpdateLoginRequest) error {
	u := fmt.Sprintf("%s/users/update_login", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
