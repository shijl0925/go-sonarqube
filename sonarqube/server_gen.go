package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/server"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Server service

// Version - Version of SonarQube in plain text
// Since 2.10
func (s *Server) Version(ctx context.Context, r server.VersionRequest) (*server.VersionResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/version", s.path)
	v := new(server.VersionResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
