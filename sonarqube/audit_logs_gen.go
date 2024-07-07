package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/audit_logs"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type AuditLogs service

// Download - Returns security related audits of this SonarQube instance. Logs are returned in JSON format.
// Requires the system administration permission
//
// Since 9.1
// Changelog:
//
//	10.2: Fields 'permissionUuid', 'componentUuid', 'groupUuid', 'permissionTemplateUuid',
// 'devOpsPlatformSettingUuid', 'qualityGateUuid', 'patUuid', 'userUuid', 'pluginUuid', 'webhookUuid', 'tokenUuid' in response are
// now deprecated.
//	9.5: Field 'userTriggered' added to the response payload.
func (s *AuditLogs) Download(ctx context.Context, r audit_logs.DownloadRequest) (*audit_logs.DownloadResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/download", s.path)
	v := new(audit_logs.DownloadResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
