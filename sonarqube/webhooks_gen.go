package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/webhooks"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Webhooks service

// Create - Create a Webhook.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 7.1
// Changelog:
//   10.6: The minimum length of parameter 'secret' increased to 16.
func (s *Webhooks) Create(ctx context.Context, r webhooks.CreateRequest) (*webhooks.CreateResponse, error) {
	u := fmt.Sprintf("%s/webhooks/create", API)
	v := new(webhooks.CreateResponse)

	_, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Delete - Delete a Webhook.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 7.1
// Changelog:
func (s *Webhooks) Delete(ctx context.Context, r webhooks.DeleteRequest) error {
	u := fmt.Sprintf("%s/webhooks/delete", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

// Deliveries - Get the recent deliveries for a specified project or Compute Engine task.
// Require 'Administer' permission on the related project.
// Note that additional information are returned by api/webhooks/delivery.
func (s *Webhooks) Deliveries(ctx context.Context, r webhooks.DeliveriesRequest, p paging.Params) (*webhooks.DeliveriesResponse, error) {
	u := fmt.Sprintf("%s/webhooks/deliveries", API)
	v := new(webhooks.DeliveriesResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Webhooks) DeliveriesAll(ctx context.Context, r webhooks.DeliveriesRequest) (*webhooks.DeliveriesResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &webhooks.DeliveriesResponseAll{}
	for {
		res, err := s.Deliveries(ctx, r, p)
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
func (s *Webhooks) Delivery(ctx context.Context, r webhooks.DeliveryRequest) (*webhooks.DeliveryResponse, error) {
	u := fmt.Sprintf("%s/webhooks/delivery", API)
	v := new(webhooks.DeliveryResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// List - Search for global webhooks or project webhooks. Webhooks are ordered by name.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *Webhooks) List(ctx context.Context, r webhooks.ListRequest) (*webhooks.ListResponse, error) {
	u := fmt.Sprintf("%s/webhooks/list", API)
	v := new(webhooks.ListResponse)

	_, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Update - Update a Webhook.
// Requires 'Administer' permission on the specified project, or global 'Administer' permission.
// Since 7.1
// Changelog:
func (s *Webhooks) Update(ctx context.Context, r webhooks.UpdateRequest) error {
	u := fmt.Sprintf("%s/webhooks/update", API)

	_, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
