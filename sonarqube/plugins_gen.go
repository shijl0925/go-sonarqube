package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/plugins"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Plugins service

// Available - Get the list of all the plugins available for installation on the SonarQube instance, sorted by plugin name.
// Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
// Update status values are:
//  * COMPATIBLE: plugin is compatible with current SonarQube instance.
//  * INCOMPATIBLE: plugin is not compatible with current SonarQube instance.
//  * REQUIRES_SYSTEM_UPGRADE: plugin requires SonarQube to be upgraded before being installed.
//  * DEPS_REQUIRE_SYSTEM_UPGRADE: at least one plugin on which the plugin is dependent requires SonarQube to be upgraded.
// Require 'Administer System' permission.
// Since 5.2
// Changelog:
func (s *Plugins) Available(ctx context.Context, r plugins.AvailableRequest) (*plugins.AvailableResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/available", s.path)
	v := new(plugins.AvailableResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// CancelAll - Cancels any operation pending on any plugin (install, update or uninstall)
// Requires user to be authenticated with Administer System permissions
// Since 5.2
// Changelog:
func (s *Plugins) CancelAll(ctx context.Context, r plugins.CancelAllRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/cancel_all", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Install - Installs the latest version of a plugin specified by its key.
// Plugin information is retrieved from Update Center.
// Fails if used on commercial editions or plugin risk consent has not been accepted.
// Requires user to be authenticated with Administer System permissions
// Since 5.2
// Changelog:
func (s *Plugins) Install(ctx context.Context, r plugins.InstallRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/install", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Installed - Get the list of all the plugins installed on the SonarQube instance, sorted by plugin name.
// Requires authentication.
// Since 5.2
// Changelog:
//   10.4: The response field 'requiredForLanguages' is added for plugins that support it
//   9.8: The 'documentationPath' field is deprecated
//   9.7: Authentication check added
//   8.0: The 'documentationPath' field is added
//   7.0: The fields 'compressedHash' and 'compressedFilename' are added
//   6.6: The 'filename' field is added
//   6.6: The 'fileHash' field is added
//   6.6: The 'sonarLintSupported' field is added
//   6.6: The 'updatedAt' field is added
func (s *Plugins) Installed(ctx context.Context, r plugins.InstalledRequest) (*plugins.InstalledResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/installed", s.path)
	v := new(plugins.InstalledResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Pending - Get the list of plugins which will either be installed or removed at the next startup of the SonarQube instance, sorted by plugin name.
// Require 'Administer System' permission.
// Since 5.2
// Changelog:
//   9.8: The 'documentationPath' field is deprecated
//   8.0: The 'documentationPath' field is added
func (s *Plugins) Pending(ctx context.Context, r plugins.PendingRequest) (*plugins.PendingResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/pending", s.path)
	v := new(plugins.PendingResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Uninstall - Uninstalls the plugin specified by its key.
// Requires user to be authenticated with Administer System permissions.
// Since 5.2
// Changelog:
func (s *Plugins) Uninstall(ctx context.Context, r plugins.UninstallRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/uninstall", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Update - Updates a plugin specified by its key to the latest version compatible with the SonarQube instance.
// Plugin information is retrieved from Update Center.
// Requires user to be authenticated with Administer System permissions
// Since 5.2
// Changelog:
func (s *Plugins) Update(ctx context.Context, r plugins.UpdateRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates - Lists plugins installed on the SonarQube instance for which at least one newer version is available, sorted by plugin name.
// Each newer version is listed, ordered from the oldest to the newest, with its own update/compatibility status.
// Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
// Update status values are: [COMPATIBLE, INCOMPATIBLE, REQUIRES_UPGRADE, DEPS_REQUIRE_UPGRADE].
// Require 'Administer System' permission.
// Since 5.2
// Changelog:
func (s *Plugins) Updates(ctx context.Context, r plugins.UpdatesRequest) (*plugins.UpdatesResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/updates", s.path)
	v := new(plugins.UpdatesResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
