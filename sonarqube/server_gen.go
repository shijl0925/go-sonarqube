package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/server"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Server service

func (s *Server) Version(ctx context.Context, r server.VersionRequest) (*server.VersionResponse, error) {
	u := fmt.Sprintf("%s/server/version", API)
	v := new(server.VersionResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
