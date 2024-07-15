package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/duplications"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Duplications service

// Show - Get duplications. Require Browse permission on file's project
// Since 4.4
// Changelog:
//   9.6: The fields 'subProject', 'subProjectName' were removed from the response.
//   8.8: Deprecated parameter 'uuid' was removed.
//   8.8: The fields 'uuid', 'projectUuid', 'subProjectUuid' were removed from the response.
//   6.5: Parameter 'uuid' is now deprecated.
//   6.5: The fields 'uuid', 'projectUuid', 'subProjectUuid' are now deprecated in the response.
func (s *Duplications) Show(ctx context.Context, r duplications.ShowRequest) (*duplications.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(duplications.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
