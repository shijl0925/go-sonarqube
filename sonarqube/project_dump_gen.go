package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_dump"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectDump service

// Export - Triggers project dump so that the project can be imported to another SonarQube server (see api/project_dump/import, available in Enterprise Edition). Requires the 'Administer' permission.
// Since 1.0
// Changelog:
//
//	9.2: Moved from Enterprise Edition to Community Edition
func (s *ProjectDump) Export(ctx context.Context, r project_dump.ExportRequest) (*project_dump.ExportResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/export", s.path)
	v := new(project_dump.ExportResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Import - Triggers the import of a project dump. Permission 'Administer' is required. This feature is provided by the Governance plugin.
// Since 1.0
func (s *ProjectDump) Import(ctx context.Context, r project_dump.ImportRequest) (*project_dump.ImportResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/import", s.path)
	v := new(project_dump.ImportResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
