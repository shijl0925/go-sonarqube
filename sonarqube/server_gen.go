package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/server"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Server service

// Version - Version of SonarQube in plain text
// Since 2.10
// Changelog:
func (s *Server) Version(ctx context.Context, r server.VersionRequest) (*server.VersionResponse, error) {
	u := fmt.Sprintf("%s/version", s.path)
	v := new(server.VersionResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
