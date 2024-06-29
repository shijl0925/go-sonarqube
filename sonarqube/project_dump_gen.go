package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/project_dump"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectDump service

func (s *ProjectDump) Export(r project_dump.ExportRequest) (*project_dump.ExportResponse, error) {
	u := fmt.Sprintf("%s/project_dump/export", API)
	v := new(project_dump.ExportResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *ProjectDump) Import(r project_dump.ImportRequest) (*project_dump.ImportResponse, error) {
	u := fmt.Sprintf("%s/project_dump/import", API)
	v := new(project_dump.ImportResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
