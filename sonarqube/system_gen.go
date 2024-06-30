package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/system"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type System service

func (s *System) ChangeLogLevel(ctx context.Context, r system.ChangeLogLevelRequest) error {
	u := fmt.Sprintf("%s/system/change_log_level", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *System) DbMigrationStatus(ctx context.Context, r system.DbMigrationStatusRequest) (*system.DbMigrationStatusResponse, error) {
	u := fmt.Sprintf("%s/system/db_migration_status", API)
	v := new(system.DbMigrationStatusResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *System) Health(ctx context.Context, r system.HealthRequest) (*system.HealthResponse, error) {
	u := fmt.Sprintf("%s/system/health", API)
	v := new(system.HealthResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *System) Info(ctx context.Context, r system.InfoRequest) (*system.InfoResponse, error) {
	u := fmt.Sprintf("%s/system/info", API)
	v := new(system.InfoResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *System) Logs(ctx context.Context, r system.LogsRequest) (*system.LogsResponse, error) {
	u := fmt.Sprintf("%s/system/logs", API)
	v := new(system.LogsResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *System) MigrateDb(ctx context.Context, r system.MigrateDbRequest) (*system.MigrateDbResponse, error) {
	u := fmt.Sprintf("%s/system/migrate_db", API)
	v := new(system.MigrateDbResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *System) Ping(ctx context.Context, r system.PingRequest) (*system.PingResponse, error) {
	u := fmt.Sprintf("%s/system/ping", API)
	v := new(system.PingResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *System) Restart(ctx context.Context, r system.RestartRequest) error {
	u := fmt.Sprintf("%s/system/restart", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *System) Status(ctx context.Context, r system.StatusRequest) (*system.StatusResponse, error) {
	u := fmt.Sprintf("%s/system/status", API)
	v := new(system.StatusResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *System) Upgrades(ctx context.Context, r system.UpgradesRequest) (*system.UpgradesResponse, error) {
	u := fmt.Sprintf("%s/system/upgrades", API)
	v := new(system.UpgradesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
