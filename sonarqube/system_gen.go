package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/system"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type System service

// ChangeLogLevel - Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.
// Since 5.2
// Changelog:
func (s *System) ChangeLogLevel(ctx context.Context, r system.ChangeLogLevelRequest) error {
	u := fmt.Sprintf("%s/system/change_log_level", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
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
func (s *System) DbMigrationStatus(ctx context.Context, r system.DbMigrationStatusRequest) (*system.DbMigrationStatusResponse, error) {
	u := fmt.Sprintf("%s/system/db_migration_status", API)
	v := new(system.DbMigrationStatusResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Health - Provide health status of SonarQube.<p>Although global health is calculated based on both application and search nodes, detailed information is returned only for application nodes.</p><p>
//   * GREEN: SonarQube is fully operational
//   * YELLOW: SonarQube is usable, but it needs attention in order to be fully operational
//   * RED: SonarQube is not operational
//  </p>
// Requires the 'Administer System' permission or system passcode (see WEB_SYSTEM_PASS_CODE in sonar.properties).
// When SonarQube is in safe mode (waiting or running a database upgrade), only the authentication with a system passcode is supported.
func (s *System) Health(ctx context.Context, r system.HealthRequest) (*system.HealthResponse, error) {
	u := fmt.Sprintf("%s/system/health", API)
	v := new(system.HealthResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Info - Get detailed information about system configuration.
// Requires 'Administer' permissions.
func (s *System) Info(ctx context.Context, r system.InfoRequest) (*system.InfoResponse, error) {
	u := fmt.Sprintf("%s/system/info", API)
	v := new(system.InfoResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Logs - Get system logs in plain-text format. Requires system administration permission.
func (s *System) Logs(ctx context.Context, r system.LogsRequest) (*system.LogsResponse, error) {
	u := fmt.Sprintf("%s/system/logs", API)
	v := new(system.LogsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
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
// Changelog:
func (s *System) MigrateDb(ctx context.Context, r system.MigrateDbRequest) (*system.MigrateDbResponse, error) {
	u := fmt.Sprintf("%s/system/migrate_db", API)
	v := new(system.MigrateDbResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Ping - Answers "pong" as plain-text
func (s *System) Ping(ctx context.Context, r system.PingRequest) (*system.PingResponse, error) {
	u := fmt.Sprintf("%s/system/ping", API)
	v := new(system.PingResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Restart - Restarts server. Requires 'Administer System' permission. Performs a full restart of the Web, Search and Compute Engine Servers processes. Does not reload sonar.properties.
// Since 4.3
// Changelog:
func (s *System) Restart(ctx context.Context, r system.RestartRequest) error {
	u := fmt.Sprintf("%s/system/restart", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Status - Get state information about SonarQube.<p>status: the running status
//   * STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing
//   * UP: SonarQube instance is up and running
//   * DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).
//   * RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).
//   * DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.
//   * DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)
//  </p>
func (s *System) Status(ctx context.Context, r system.StatusRequest) (*system.StatusResponse, error) {
	u := fmt.Sprintf("%s/system/status", API)
	v := new(system.StatusResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Upgrades - Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.
// Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
func (s *System) Upgrades(ctx context.Context, r system.UpgradesRequest) (*system.UpgradesResponse, error) {
	u := fmt.Sprintf("%s/system/upgrades", API)
	v := new(system.UpgradesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
