package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/audit_logs"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AuditLogs service

func (s *AuditLogs) Download(r audit_logs.DownloadRequest) (*audit_logs.DownloadResponse, error) {
	u := fmt.Sprintf("%s/audit_logs/download", API)
	v := new(audit_logs.DownloadResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
