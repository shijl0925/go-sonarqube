package sonarqube

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const (
	basicAuth int = iota
	privateToken
	Anonymous
)

type Client struct {
	client              *http.Client
	host                string
	username            string
	password            string
	token               string
	authType            int
	AlmIntegrations     *AlmIntegrations
	AlmSettings         *AlmSettings
	AnalysisCache       *AnalysisCache
	Applications        *Applications
	AuditLogs           *AuditLogs
	Authentication      *Authentication
	Ce                  *Ce
	Components          *Components
	Duplications        *Duplications
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
	System              *System
	UserGroups          *UserGroups
	UserTokens          *UserTokens
	Users               *Users
	Views               *Views
	Webhooks            *Webhooks
	Webservices         *Webservices
}

type service struct {
	client *Client
	path   string
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

	c.host = sonarURL
	c.AlmIntegrations = &AlmIntegrations{client: c, path: "api/alm_integrations"}
	c.AlmSettings = &AlmSettings{client: c, path: "api/alm_settings"}
	c.AnalysisCache = &AnalysisCache{client: c, path: "api/analysis_cache"}
	c.Applications = &Applications{client: c, path: "api/applications"}
	c.AuditLogs = &AuditLogs{client: c, path: "api/audit_logs"}
	c.Authentication = &Authentication{client: c, path: "api/authentication"}
	c.Ce = &Ce{client: c, path: "api/ce"}
	c.Components = &Components{client: c, path: "api/components"}
	c.Duplications = &Duplications{client: c, path: "api/duplications"}
	c.Editions = &Editions{client: c, path: "api/editions"}
	c.Favorites = &Favorites{client: c, path: "api/favorites"}
	c.Hotspots = &Hotspots{client: c, path: "api/hotspots"}
	c.Issues = &Issues{client: c, path: "api/issues"}
	c.Languages = &Languages{client: c, path: "api/languages"}
	c.Measures = &Measures{client: c, path: "api/measures"}
	c.Metrics = &Metrics{client: c, path: "api/metrics"}
	c.Monitoring = &Monitoring{client: c, path: "api/monitoring"}
	c.NewCodePeriods = &NewCodePeriods{client: c, path: "api/new_code_periods"}
	c.Notifications = &Notifications{client: c, path: "api/notifications"}
	c.Permissions = &Permissions{client: c, path: "api/permissions"}
	c.Plugins = &Plugins{client: c, path: "api/plugins"}
	c.ProjectAnalyses = &ProjectAnalyses{client: c, path: "api/project_analyses"}
	c.ProjectBadges = &ProjectBadges{client: c, path: "api/project_badges"}
	c.ProjectBranches = &ProjectBranches{client: c, path: "api/project_branches"}
	c.ProjectDump = &ProjectDump{client: c, path: "api/project_dump"}
	c.ProjectLinks = &ProjectLinks{client: c, path: "api/project_links"}
	c.ProjectPullRequests = &ProjectPullRequests{client: c, path: "api/project_pull_requests"}
	c.ProjectTags = &ProjectTags{client: c, path: "api/project_tags"}
	c.Projects = &Projects{client: c, path: "api/projects"}
	c.Qualitygates = &Qualitygates{client: c, path: "api/qualitygates"}
	c.Qualityprofiles = &Qualityprofiles{client: c, path: "api/qualityprofiles"}
	c.Rules = &Rules{client: c, path: "api/rules"}
	c.Server = &Server{client: c, path: "api/server"}
	c.Settings = &Settings{client: c, path: "api/settings"}
	c.Sources = &Sources{client: c, path: "api/sources"}
	c.System = &System{client: c, path: "api/system"}
	c.UserGroups = &UserGroups{client: c, path: "api/user_groups"}
	c.UserTokens = &UserTokens{client: c, path: "api/user_tokens"}
	c.Users = &Users{client: c, path: "api/users"}
	c.Views = &Views{client: c, path: "api/views"}
	c.Webhooks = &Webhooks{client: c, path: "api/webhooks"}
	c.Webservices = &Webservices{client: c, path: "api/webservices"}

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

// 封装认证处理
func (c *Client) handleAuth(req *http.Request) {
	switch c.authType {
	case basicAuth:
		req.SetBasicAuth(c.username, c.password)
	case privateToken:
		req.SetBasicAuth(c.token, "")
	default:
		// do nothing
	}
}

func (c *Client) NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	// 认证处理
	c.handleAuth(req)

	// 设置通用请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)

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
	return resp, nil
}

func (c *Client) Call(ctx context.Context, method string, u string, v interface{}, opt ...interface{}) (*http.Response, error) {
	u = fmt.Sprintf("%s/%s", c.host, u)
	var req *http.Request
	var err error

	if method == http.MethodGet {
		for _, o := range opt {
			urlStr, err := addOptions(u, o)
			if err != nil {
				return nil, fmt.Errorf("could not Parse query values: %v", err)
			}
			u = urlStr
		}

		req, err = c.NewRequest(ctx, "GET", u, nil)
		if err != nil {
			return nil, fmt.Errorf("could not create request: %v", err)
		}
	} else {
		values := make(url.Values)
		encoder := form.NewEncoder()

		for _, o := range opt {
			vs, err := encoder.Encode(o)
			if err != nil {
				return nil, fmt.Errorf("could not encode form values: %v", err)
			}
			for k, v := range vs {
				values[k] = append(values[k], v...)
			}
		}

		req, err = c.NewRequest(ctx, "POST", u, strings.NewReader(values.Encode()))
		if err != nil {
			return nil, fmt.Errorf("could not create request: %v", err)
		}
	}

	isText := false

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute request: %+v", err)
	}

	if v != nil {
		defer resp.Body.Close()

		res := reflect.ValueOf(v).Elem()
		if res.Kind() == reflect.String {
			isText = true
		}

		if isText {
			buf := new(strings.Builder)
			_, err := io.Copy(buf, resp.Body)
			if err != nil {
				return resp, fmt.Errorf("could not read response body: %v", err)
			}
			res.SetString(buf.String())
		} else {
			if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
				return nil, fmt.Errorf("could not decode response: %v", err)
			}
		}
	}
	return resp, err
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()

	return origURL.String(), nil
}
