package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/duplications"
	"net/http"
)

type Duplications service

func (s *Duplications) Show(ctx context.Context, r duplications.ShowRequest) (*duplications.ShowResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(duplications.ShowResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
