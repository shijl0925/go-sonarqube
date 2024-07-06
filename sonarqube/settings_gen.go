package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/settings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Settings service

// ListDefinitions - List settings definitions.
// Requires 'Browse' permission when a component is specified
// To access licensed settings, authentication is required
// To access secured settings, one of the following permissions is required:
//  * 'Execute Analysis'
//  * 'Administer System'
//  * 'Administer' rights on the specified component
//
// Since 6.3
// Changelog:
//   10.1: The use of module keys in parameter 'component' is removed
//   7.6: The use of module keys in parameter 'component' is deprecated
func (s *Settings) ListDefinitions(ctx context.Context, r settings.ListDefinitionsRequest) (*settings.ListDefinitionsResponse, error) {
	u := fmt.Sprintf("%s/list_definitions", s.path)
	v := new(settings.ListDefinitionsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Reset - Remove a setting value.
// The settings defined in conf/sonar.properties are read-only and can't be changed.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified component
//
// Since 6.1
// Changelog:
//   10.1: Param 'component' now only accept keys for projects, applications, portfolios or subportfolios
//   10.1: Internal parameters 'branch' and 'pullRequest' were removed
//   8.8: Deprecated parameter 'componentKey' has been removed
//   7.6: The use of module keys in parameter 'component' is deprecated
//   7.1: The settings defined in conf/sonar.properties are read-only and can't be changed
func (s *Settings) Reset(ctx context.Context, r settings.ResetRequest) error {
	u := fmt.Sprintf("%s/reset", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Set - Update a setting value.
// Either 'value' or 'values' must be provided.
// The settings defined in conf/sonar.properties are read-only and can't be changed.
// Requires one of the following permissions:
//  * 'Administer System'
//  * 'Administer' rights on the specified component
//
// Since 6.1
// Changelog:
//   10.1: Param 'component' now only accept keys for projects, applications, portfolios or subportfolios
//   10.1: The use of module keys in parameter 'component' is removed
//   8.8: Deprecated parameter 'componentKey' has been removed
//   7.6: The use of module keys in parameter 'component' is deprecated
//   7.1: The settings defined in conf/sonar.properties are read-only and can't be changed
func (s *Settings) Set(ctx context.Context, r settings.SetRequest) error {
	u := fmt.Sprintf("%s/set", s.path)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Values - List settings values.
// If no value has been set for a setting, then the default value is returned.
// The settings from conf/sonar.properties are excluded from results.
// Requires 'Browse' or 'Execute Analysis' permission when a component is specified.
// Secured settings values are not returned by the endpoint.
//
// Since 6.3
// Changelog:
//   10.1: The use of module keys in parameter 'component' is removed
//   9.1: The secured settings values are no longer returned. Secured settings keys that have a value are now returned in setSecuredSettings array.
//   7.6: The use of module keys in parameter 'component' is deprecated
//   7.1: The settings from conf/sonar.properties are excluded from results.
func (s *Settings) Values(ctx context.Context, r settings.ValuesRequest) (*settings.ValuesResponse, error) {
	u := fmt.Sprintf("%s/values", s.path)
	v := new(settings.ValuesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
