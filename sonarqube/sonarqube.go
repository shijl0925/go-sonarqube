// Copyright 2013 The go-github AUTHORS. All rights reserved.
// Copyright 2021 Reinoud Kruithof
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sonarqube

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

var API string

const (
	basicAuth int = iota
	privateToken
	Anonymous int = iota
)

type Client struct {
	client   *http.Client
	API      string
	username string
	password string
	token    string
	authType int

	AlmIntegrations     *AlmIntegrations
	AlmSettings         *AlmSettings
	AnalysisCache       *AnalysisCache
	Applications        *Applications
	AuditLogs           *AuditLogs
	Authentication      *Authentication
	Ce                  *Ce
	Components          *Components
	Editions            *Editions
	Favorites           *Favorites
	Hotspots            *Hotspots
	Issues              *Issues
	Languages           *Languages
	Measures            *Measures
	Metrics             *Metrics
	Monitoring          *Monitoring
	NewCodePeriods      *NewCodePeriods
	Notifications       *Notifications
	Permissions         *Permissions
	Plugins             *Plugins
	ProjectAnalyses     *ProjectAnalyses
	ProjectBadges       *ProjectBadges
	ProjectBranches     *ProjectBranches
	ProjectDump         *ProjectDump
	ProjectLinks        *ProjectLinks
	ProjectPullRequests *ProjectPullRequests
	ProjectTags         *ProjectTags
	Projects            *Projects
	Qualitygates        *Qualitygates
	Qualityprofiles     *Qualityprofiles
	Rules               *Rules
	Server              *Server
	Settings            *Settings
	Sources             *Sources
	UserGroups          *UserGroups
	UserTokens          *UserTokens
	Users               *Users
	Views               *Views
	Webhooks            *Webhooks
	Webservices         *Webservices
}

type service struct {
	client *Client
}

func NewClient(sonarURL string, username string, password string, client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}

	var authType int
	if len(username) != 0 && len(password) != 0 {
		authType = basicAuth
	} else {
		authType = Anonymous
	}

	c := &Client{
		client:   client,
		username: username,
		password: password,
		authType: authType,
	}

	API = sonarURL

	c.AlmIntegrations = &AlmIntegrations{client: c}
	c.AlmSettings = &AlmSettings{client: c}
	c.AnalysisCache = &AnalysisCache{client: c}
	c.Applications = &Applications{client: c}
	c.AuditLogs = &AuditLogs{client: c}
	c.Authentication = &Authentication{client: c}
	c.Ce = &Ce{client: c}
	c.Components = &Components{client: c}
	c.Editions = &Editions{client: c}
	c.Favorites = &Favorites{client: c}
	c.Hotspots = &Hotspots{client: c}
	c.Issues = &Issues{client: c}
	c.Languages = &Languages{client: c}
	c.Measures = &Measures{client: c}
	c.Metrics = &Metrics{client: c}
	c.Monitoring = &Monitoring{client: c}
	c.NewCodePeriods = &NewCodePeriods{client: c}
	c.Notifications = &Notifications{client: c}
	c.Permissions = &Permissions{client: c}
	c.Plugins = &Plugins{client: c}
	c.ProjectAnalyses = &ProjectAnalyses{client: c}
	c.ProjectBadges = &ProjectBadges{client: c}
	c.ProjectBranches = &ProjectBranches{client: c}
	c.ProjectDump = &ProjectDump{client: c}
	c.ProjectLinks = &ProjectLinks{client: c}
	c.ProjectPullRequests = &ProjectPullRequests{client: c}
	c.ProjectTags = &ProjectTags{client: c}
	c.Projects = &Projects{client: c}
	c.Qualitygates = &Qualitygates{client: c}
	c.Qualityprofiles = &Qualityprofiles{client: c}
	c.Rules = &Rules{client: c}
	c.Server = &Server{client: c}
	c.Settings = &Settings{client: c}
	c.Sources = &Sources{client: c}
	c.UserGroups = &UserGroups{client: c}
	c.UserTokens = &UserTokens{client: c}
	c.Users = &Users{client: c}
	c.Views = &Views{client: c}
	c.Webhooks = &Webhooks{client: c}
	c.Webservices = &Webservices{client: c}

	return c
}

func NewClientByToken(sonarURL string, token string, client *http.Client) *Client {
	c := NewClient(sonarURL, "", "", client)
	c.token = token

	var authType int
	if len(token) != 0 {
		authType = privateToken
	} else {
		authType = Anonymous
	}
	c.authType = authType
	return c
}

func (c *Client) PostRequest(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	switch c.authType {
	case basicAuth:
		req.SetBasicAuth(c.username, c.password)
	case privateToken:
		req.SetBasicAuth(c.token, "")
	default:
		// do nothing
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) GetRequest(url string, params ...string) (*http.Request, error) {
	if l := len(params); l%2 != 0 {
		return nil, fmt.Errorf("params must be an even number, %d given", l)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()

	for i := 0; i < len(params); i++ {
		q.Add(params[i], params[i+1])
		i++
	}
	req.URL.RawQuery = q.Encode()

	switch c.authType {
	case basicAuth:
		req.SetBasicAuth(c.username, c.password)
	case privateToken:
		req.SetBasicAuth(c.token, "")
	default:
		// do nothing
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func (c *Client) Call(method string, u string, v interface{}, opt ...interface{}) (*http.Response, error) {
	var req *http.Request
	var err error
	if method == http.MethodGet {
		params := paramsFrom(opt...)
		req, err = c.GetRequest(u, params...)
		if err != nil {
			return nil, fmt.Errorf("could not create request: %+v", err)
		}
	} else {
		encoder := form.NewEncoder()
		values, err := encoder.Encode(opt)
		if err != nil {
			return nil, fmt.Errorf("could not encode form values: %+v", err)
		}
		req, err = c.PostRequest(u, strings.NewReader(values.Encode()))
		if err != nil {
			return nil, fmt.Errorf("could not create request: %+v", err)
		}
	}

	isText := false
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("v must be a pointer type")
	}
	res := val.Elem()
	if res.Kind() == reflect.String {
		req.Header.Set("Accept", "text/plain")
		isText = true
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute request: %+v", err)
	}

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return nil, fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return nil, errorResponse
		}
	}
	if v != nil {
		defer resp.Body.Close()

		if isText {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return resp, err
			}
			w := val.Elem()
			w.SetString(string(body))
		} else {

			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, fmt.Errorf("could not decode response: %+v", err)
			}
		}
	}
	return resp, err
}

// paramsFrom creates a slice with interleaving param and value entries, i.e. ["key1", "value1", "key2, "value2"]
func paramsFrom(items ...interface{}) []string {
	allParams := make([]string, 0)

	for _, item := range items {
		v := reflect.ValueOf(item)
		t := v.Type()

		params := make([]string, 2*v.NumField())

		for i := 0; i < v.NumField(); i++ {
			j := i * 2
			k := j + 1

			if v.Field(i).IsZero() {
				continue
			}

			// Convert some basic types to strings for convenience.
			// Note: other types should not be used as parameter values.
			fieldValue := ""
			switch t.Field(i).Type.Name() {
			case "int":
				fieldValue = strconv.Itoa(v.Field(i).Interface().(int))
			case "string":
				fieldValue = v.Field(i).Interface().(string)
			case "bool":
				fieldValue = strconv.FormatBool(v.Field(i).Interface().(bool))
			}

			params[j] = strcase.ToLowerCamel(t.Field(i).Name)
			params[k] = fieldValue
		}

		allParams = append(allParams, params...)
	}

	return allParams
}
