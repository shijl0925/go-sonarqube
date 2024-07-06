package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/components"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Components service

// Search - Search for components
// Since 6.3
// Changelog:
//   8.4: Param 'language' has been removed
//   8.4: The use of 'DIR','FIL','UTS' and 'BRC' as values for parameter 'qualifiers' is no longer supported
//   8.0: Field 'id' from response has been removed
//   7.6: The use of 'BRC' as value for parameter 'qualifiers' is deprecated
func (s *Components) Search(ctx context.Context, r components.SearchRequest, p paging.Params) (*components.SearchResponse, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(components.SearchResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Components) SearchAll(ctx context.Context, r components.SearchRequest) (*components.SearchResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &components.SearchResponseAll{}
	for {
		res, err := s.Search(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to components.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// Show - Returns a component (file, directory, project, portfolio…) and its ancestors. The ancestors are ordered from the parent to the root project. Requires the following permission: 'Browse' on the project of the specified component.
// Since 5.4
// Changelog:
//   10.1: The use of module keys in parameter 'component' is removed
//   7.6: The use of module keys in parameter 'component' is deprecated
func (s *Components) Show(ctx context.Context, r components.ShowRequest) (*components.ShowResponse, error) {
	u := fmt.Sprintf("%s/show", s.path)
	v := new(components.ShowResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Tree - Navigate through components based on the chosen strategy.
// Requires the following permission: 'Browse' on the specified project.
// When limiting search with the q parameter, directories are not returned.
// Since 5.4
// Changelog:
//   10.1: The use of module keys in parameter 'component' is removed
//   10.1: The use of 'BRC' as value for parameter 'qualifiers' is removed
//   7.6: The use of 'BRC' as value for parameter 'qualifiers' is deprecated
//   7.6: The use of module keys in parameter 'component' is deprecated
func (s *Components) Tree(ctx context.Context, r components.TreeRequest, p paging.Params) (*components.TreeResponse, error) {
	u := fmt.Sprintf("%s/tree", s.path)
	v := new(components.TreeResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Components) TreeAll(ctx context.Context, r components.TreeRequest) (*components.TreeResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &components.TreeResponseAll{}
	for {
		res, err := s.Tree(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to components.Tree: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
