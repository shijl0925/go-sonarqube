package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/audit_logs"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AuditLogs service

// Download - Returns security related audits of this SonarQube instance. Logs are returned in JSON format.
// Requires the system administration permission
//
func (s *AuditLogs) Download(ctx context.Context, r audit_logs.DownloadRequest) (*audit_logs.DownloadResponse, error) {
	u := fmt.Sprintf("%s/audit_logs/download", API)
	v := new(audit_logs.DownloadResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
