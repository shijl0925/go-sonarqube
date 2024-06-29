package webservices

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ListRequest List web services
type ListRequest struct {
	IncludeInternals string `form:"include_internals,omitempty"` // Include web services that are implemented for internal use only. Their forward-compatibility is not assured
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	WebServices []struct {
		Actions []struct {
			Changelog []struct {
				Description string `json:"description,omitempty"`
				Version     string `json:"version,omitempty"`
			} `json:"changelog,omitempty"`
			DeprecatedSince    string `json:"deprecatedSince,omitempty"`
			Description        string `json:"description,omitempty"`
			HasResponseExample bool   `json:"hasResponseExample,omitempty"`
			Internal           bool   `json:"internal,omitempty"`
			Key                string `json:"key,omitempty"`
			Params             []struct {
				Internal           bool     `json:"internal,omitempty"`
				Key                string   `json:"key,omitempty"`
				MaximumValue       float64  `json:"maximumValue,omitempty"`
				Required           bool     `json:"required,omitempty"`
				MaximumLength      float64  `json:"maximumLength,omitempty"`
				MinimumLength      float64  `json:"minimumLength,omitempty"`
				DefaultValue       string   `json:"defaultValue,omitempty"`
				DeprecatedKey      string   `json:"deprecatedKey,omitempty"`
				DeprecatedKeySince string   `json:"deprecatedKeySince,omitempty"`
				DeprecatedSince    string   `json:"deprecatedSince,omitempty"`
				Description        string   `json:"description,omitempty"`
				ExampleValue       string   `json:"exampleValue,omitempty"`
				PossibleValues     []string `json:"possibleValues,omitempty"`
				Since              string   `json:"since,omitempty"`
			} `json:"params,omitempty"`
			Post  bool   `json:"post,omitempty"`
			Since string `json:"since,omitempty"`
		} `json:"actions,omitempty"`
		Description string `json:"description,omitempty"`
		Path        string `json:"path,omitempty"`
		Since       string `json:"since,omitempty"`
	} `json:"webServices,omitempty"`
}

// ResponseExampleRequest Display web service response example
type ResponseExampleRequest struct {
	Action     string `form:"action,omitempty"`     // Action of the web service
	Controller string `form:"controller,omitempty"` // Controller of the web service
}

// ResponseExampleResponse is the response for ResponseExampleRequest
type ResponseExampleResponse string
