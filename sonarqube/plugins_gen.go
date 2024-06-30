package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/plugins"
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
func (s *Plugins) Available(ctx context.Context, r plugins.AvailableRequest) (*plugins.AvailableResponse, error) {
	u := fmt.Sprintf("%s/plugins/available", API)
	v := new(plugins.AvailableResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// CancelAll - Cancels any operation pending on any plugin (install, update or uninstall)
// Requires user to be authenticated with Administer System permissions
// Since 5.2
// Changelog:
func (s *Plugins) CancelAll(ctx context.Context, r plugins.CancelAllRequest) error {
	u := fmt.Sprintf("%s/plugins/cancel_all", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Install - Installs the latest version of a plugin specified by its key.
// Plugin information is retrieved from Update Center.
// Fails if used on commercial editions or plugin risk consent has not been accepted.
// Requires user to be authenticated with Administer System permissions
// Since 5.2
// Changelog:
func (s *Plugins) Install(ctx context.Context, r plugins.InstallRequest) error {
	u := fmt.Sprintf("%s/plugins/install", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Installed - Get the list of all the plugins installed on the SonarQube instance, sorted by plugin name.
// Requires authentication.
func (s *Plugins) Installed(ctx context.Context, r plugins.InstalledRequest) (*plugins.InstalledResponse, error) {
	u := fmt.Sprintf("%s/plugins/installed", API)
	v := new(plugins.InstalledResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Pending - Get the list of plugins which will either be installed or removed at the next startup of the SonarQube instance, sorted by plugin name.
// Require 'Administer System' permission.
func (s *Plugins) Pending(ctx context.Context, r plugins.PendingRequest) (*plugins.PendingResponse, error) {
	u := fmt.Sprintf("%s/plugins/pending", API)
	v := new(plugins.PendingResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Uninstall - Uninstalls the plugin specified by its key.
// Requires user to be authenticated with Administer System permissions.
// Since 5.2
// Changelog:
func (s *Plugins) Uninstall(ctx context.Context, r plugins.UninstallRequest) error {
	u := fmt.Sprintf("%s/plugins/uninstall", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Update - Updates a plugin specified by its key to the latest version compatible with the SonarQube instance.
// Plugin information is retrieved from Update Center.
// Requires user to be authenticated with Administer System permissions
// Since 5.2
// Changelog:
func (s *Plugins) Update(ctx context.Context, r plugins.UpdateRequest) error {
	u := fmt.Sprintf("%s/plugins/update", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Updates - Lists plugins installed on the SonarQube instance for which at least one newer version is available, sorted by plugin name.
// Each newer version is listed, ordered from the oldest to the newest, with its own update/compatibility status.
// Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
// Update status values are: [COMPATIBLE, INCOMPATIBLE, REQUIRES_UPGRADE, DEPS_REQUIRE_UPGRADE].
// Require 'Administer System' permission.
func (s *Plugins) Updates(ctx context.Context, r plugins.UpdatesRequest) (*plugins.UpdatesResponse, error) {
	u := fmt.Sprintf("%s/plugins/updates", API)
	v := new(plugins.UpdatesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
