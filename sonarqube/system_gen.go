package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/system"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type System service

// ChangeLogLevel - Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.
// Since 5.2
func (s *System) ChangeLogLevel(ctx context.Context, r system.ChangeLogLevelRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/change_log_level", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DbMigrationStatus - Display the database migration status of SonarQube.
// State values are:
//  * NO_MIGRATION: DB is up to date with current version of SonarQube.
//  * NOT_SUPPORTED: Migration is not supported on embedded databases.
//  * MIGRATION_RUNNING: DB migration is under go.
//  * MIGRATION_SUCCEEDED: DB migration has run and has been successful.
//  * MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).
//  * MIGRATION_REQUIRED: DB migration is required.
//
// Since 5.2
// Deprecated since 10.6
// Changelog:
//
//	10.6: This endpoint is deprecated, please use its API v2 version /api/v2/system/migrations-status instead.
func (s *System) DbMigrationStatus(ctx context.Context, r system.DbMigrationStatusRequest) (*system.DbMigrationStatusResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/db_migration_status", s.path)
	v := new(system.DbMigrationStatusResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Health - Provide health status of SonarQube.<p>Although global health is calculated based on both application and search nodes, detailed information is returned only for application nodes.</p><p>
//   * GREEN: SonarQube is fully operational
//   * YELLOW: SonarQube is usable, but it needs attention in order to be fully operational
//   * RED: SonarQube is not operational
//  </p>
// Requires the 'Administer System' permission or system passcode (see WEB_SYSTEM_PASS_CODE in sonar.properties).
// When SonarQube is in safe mode (waiting or running a database upgrade), only the authentication with a system passcode is supported.
// Since 6.6
func (s *System) Health(ctx context.Context, r system.HealthRequest) (*system.HealthResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/health", s.path)
	v := new(system.HealthResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Info - Get detailed information about system configuration.
// Requires 'Administer' permissions.
// Since 5.1
// Changelog:
//
//	9.8: 'Edition' field added to the response under the 'System' section
//	9.7: 'Statistics' field has been removed from response
//	8.3: Becomes public
//	5.5: Becomes internal to easily update result
func (s *System) Info(ctx context.Context, r system.InfoRequest) (*system.InfoResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/info", s.path)
	v := new(system.InfoResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Logs - Get system logs in plain-text format. Requires system administration permission.
// Since 5.2
// Changelog:
//
//	10.4: Add support for deprecation logs in process property.
//	10.4: Deprecate property 'process' in favor of 'name'.
func (s *System) Logs(ctx context.Context, r system.LogsRequest) (*system.LogsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/logs", s.path)
	v := new(system.LogsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// MigrateDb - Migrate the database to match the current version of SonarQube.
// Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.
// State values are:
//  * NO_MIGRATION: DB is up to date with current version of SonarQube.
//  * NOT_SUPPORTED: Migration is not supported on embedded databases.
//  * MIGRATION_RUNNING: DB migration is under go.
//  * MIGRATION_SUCCEEDED: DB migration has run and has been successful.
//  * MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).
//  * MIGRATION_REQUIRED: DB migration is required.
//
// Since 5.2
func (s *System) MigrateDb(ctx context.Context, r system.MigrateDbRequest) (*system.MigrateDbResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/migrate_db", s.path)
	v := new(system.MigrateDbResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Ping - Answers "pong" as plain-text
// Since 6.3
func (s *System) Ping(ctx context.Context, r system.PingRequest) (*system.PingResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/ping", s.path)
	v := new(system.PingResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Restart - Restarts server. Requires 'Administer System' permission. Performs a full restart of the Web, Search and Compute Engine Servers processes. Does not reload sonar.properties.
// Since 4.3
func (s *System) Restart(ctx context.Context, r system.RestartRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/restart", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Status - Get state information about SonarQube.<p>status: the running status
//   * STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing
//   * UP: SonarQube instance is up and running
//   * DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).
//   * RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).
//   * DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.
//   * DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)
//  </p>
// Since 5.2
func (s *System) Status(ctx context.Context, r system.StatusRequest) (*system.StatusResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/status", s.path)
	v := new(system.StatusResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Upgrades - Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.
// Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
// Since 5.2
// Changelog:
//
//	10.5: The field 'ltsVersion' is deprecated from the response
//	10.5: The field 'ltaVersion' is added to indicate the Long-Term Active Version
//	10.5: The field 'installedVersionActive' is added to indicate if the installed version is an active version
func (s *System) Upgrades(ctx context.Context, r system.UpgradesRequest) (*system.UpgradesResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/upgrades", s.path)
	v := new(system.UpgradesResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
