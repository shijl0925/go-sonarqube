package settings

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ListDefinitionsRequest List settings definitions.<br>Requires 'Browse' permission when a component is specified<br/>To access licensed settings, authentication is required<br/>To access secured settings, one of the following permissions is required: <ul><li>'Execute Analysis'</li><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>
type ListDefinitionsRequest struct {
	Component string `url:"component,omitempty"` // Component key
}

// ListDefinitionsResponse is the response for ListDefinitionsRequest
type ListDefinitionsResponse struct {
	Definitions []struct {
		Category     string   `json:"category,omitempty"`
		DefaultValue string   `json:"defaultValue,omitempty"`
		Description  string   `json:"description,omitempty"`
		Fields       []string `json:"fields,omitempty"`
		Key          string   `json:"key,omitempty"`
		MultiValues  bool     `json:"multiValues,omitempty"`
		Name         string   `json:"name,omitempty"`
		Options      []string `json:"options,omitempty"`
		SubCategory  string   `json:"subCategory,omitempty"`
		Type         string   `json:"type,omitempty"`
	} `json:"definitions,omitempty"`
}

// ResetRequest Remove a setting value.<br>The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>
type ResetRequest struct {
	Component string `json:"component,omitempty"` // Component key. Only keys for projects, applications, portfolios or subportfolios are accepted.
	Keys      string `json:"keys"`                // Comma-separated list of keys
}

// SetRequest Update a setting value.<br>Either 'value' or 'values' must be provided.<br> The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>
type SetRequest struct {
	Component   string `json:"component,omitempty"`   // Component key. Only keys for projects, applications, portfolios or subportfolios are accepted.
	FieldValues string `json:"fieldValues,omitempty"` // Setting field values. To set several values, the parameter must be called once for each value.
	Key         string `json:"key"`                   // Setting key
	Value       string `json:"value,omitempty"`       // Setting value. To reset a value, please use the reset web service.
	Values      string `json:"values,omitempty"`      // Setting multi value. To set several values, the parameter must be called once for each value.
}

// ValuesRequest List settings values.<br>If no value has been set for a setting, then the default value is returned.<br>The settings from conf/sonar.properties are excluded from results.<br>Requires 'Browse' or 'Execute Analysis' permission when a component is specified.<br/>Secured settings values are not returned by the endpoint.<br/>
type ValuesRequest struct {
	Component string `url:"component,omitempty"` // Component key
	Keys      string `url:"keys,omitempty"`      // List of setting keys
}

// ValuesResponse is the response for ValuesRequest
type ValuesResponse struct {
	SetSecuredSettings []string `json:"setSecuredSettings,omitempty"`
	Settings           []struct {
		Inherited   bool     `json:"inherited,omitempty"`
		Key         string   `json:"key,omitempty"`
		Value       string   `json:"value,omitempty"`
		Values      []string `json:"values,omitempty"`
		FieldValues []struct {
			Boolean string `json:"boolean,omitempty"`
			Text    string `json:"text,omitempty"`
		} `json:"fieldValues,omitempty"`
	} `json:"settings,omitempty"`
}
