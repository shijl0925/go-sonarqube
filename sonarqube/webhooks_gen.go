package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/webhooks"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Webhooks service

// Create - Create a Webhook.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 7.1
// Changelog:
//
//	10.6: The minimum length of parameter 'secret' increased to 16.
func (s *Webhooks) Create(ctx context.Context, r webhooks.CreateRequest) (*webhooks.CreateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)
	v := new(webhooks.CreateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Delete - Delete a Webhook.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 7.1
func (s *Webhooks) Delete(ctx context.Context, r webhooks.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deliveries - Get the recent deliveries for a specified project or Compute Engine task.
// Require 'Administer' permission on the related project.
// Note that additional information are returned by api/webhooks/delivery.
// Since 6.2
// Changelog:
//
//	10.7: 'ceTaskId' and 'componentKey' parameters are now deprecated. These parameters won't be replaced, the deliveries related to a specific project can be obtained by fetching the webhook first, and then fetching the associated deliveries.
//	10.7: 'ceTaskId' response field is now deprecated.
func (s *Webhooks) Deliveries(ctx context.Context, r webhooks.DeliveriesRequest, p paging.Params) (*webhooks.DeliveriesResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/deliveries", s.path)
	v := new(webhooks.DeliveriesResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Webhooks) DeliveriesAll(ctx context.Context, r webhooks.DeliveriesRequest) (*webhooks.DeliveriesResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &webhooks.DeliveriesResponseAll{}
	for {
		res, _, err := s.Deliveries(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to webhooks.Deliveries: %+v", err)
		}
		response.Deliveries = append(response.Deliveries, res.Deliveries...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

// Delivery - Get a webhook delivery by its id.
// Require 'Administer System' permission.
// Note that additional information are returned by api/webhooks/delivery.
// Since 6.2
func (s *Webhooks) Delivery(ctx context.Context, r webhooks.DeliveryRequest) (*webhooks.DeliveryResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/delivery", s.path)
	v := new(webhooks.DeliveryResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// List - Search for global webhooks or project webhooks. Webhooks are ordered by name.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 7.1
// Changelog:
//
//	10.1: Field 'secret' replaced by flag 'hasSecret' in response
//	7.8: Field 'secret' added to response
func (s *Webhooks) List(ctx context.Context, r webhooks.ListRequest) (*webhooks.ListResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/list", s.path)
	v := new(webhooks.ListResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Update - Update a Webhook.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 7.1
func (s *Webhooks) Update(ctx context.Context, r webhooks.UpdateRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/update", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
