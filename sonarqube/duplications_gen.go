package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/duplications"
)

type Duplications service

func (s *Duplications) Show(ctx context.Context, r duplications.ShowRequest) (*duplications.ShowResponse, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(duplications.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}
