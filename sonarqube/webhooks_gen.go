package sonarqube

import (
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/webhooks"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Webhooks service

func (s *Webhooks) Create(r webhooks.CreateRequest) (*webhooks.CreateResponse, error) {
	u := fmt.Sprintf("%s/webhooks/create", API)
	v := new(webhooks.CreateResponse)

	_, err := s.client.Call("POST", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Webhooks) Delete(r webhooks.DeleteRequest) error {
	u := fmt.Sprintf("%s/webhooks/delete", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *Webhooks) Deliveries(r webhooks.DeliveriesRequest, p paging.Params) (*webhooks.DeliveriesResponse, error) {
	u := fmt.Sprintf("%s/webhooks/deliveries", API)
	v := new(webhooks.DeliveriesResponse)

	_, err := s.client.Call("GET", u, v, r, p)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Webhooks) DeliveriesAll(r webhooks.DeliveriesRequest) (*webhooks.DeliveriesResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &webhooks.DeliveriesResponseAll{}
	for {
		res, err := s.Deliveries(r, p)
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

func (s *Webhooks) Delivery(r webhooks.DeliveryRequest) (*webhooks.DeliveryResponse, error) {
	u := fmt.Sprintf("%s/webhooks/delivery", API)
	v := new(webhooks.DeliveryResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Webhooks) List(r webhooks.ListRequest) (*webhooks.ListResponse, error) {
	u := fmt.Sprintf("%s/webhooks/list", API)
	v := new(webhooks.ListResponse)

	_, err := s.client.Call("GET", u, v, r)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Webhooks) Update(r webhooks.UpdateRequest) error {
	u := fmt.Sprintf("%s/webhooks/update", API)

	_, err := s.client.Call("POST", u, nil, r)
	if err != nil {
		return err
	}

	return nil
}
