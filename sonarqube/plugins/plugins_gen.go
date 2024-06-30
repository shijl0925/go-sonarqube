package plugins

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AvailableRequest Get the list of all the plugins available for installation on the SonarQube instance, sorted by plugin name.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: <ul><li>COMPATIBLE: plugin is compatible with current SonarQube instance.</li><li>INCOMPATIBLE: plugin is not compatible with current SonarQube instance.</li><li>REQUIRES_SYSTEM_UPGRADE: plugin requires SonarQube to be upgraded before being installed.</li><li>DEPS_REQUIRE_SYSTEM_UPGRADE: at least one plugin on which the plugin is dependent requires SonarQube to be upgraded.</li></ul>Require 'Administer System' permission.
type AvailableRequest struct{}

// AvailableResponse is the response for AvailableRequest
type AvailableResponse struct {
	Plugins []struct {
		Category         string `json:"category,omitempty"`
		Description      string `json:"description,omitempty"`
		EditionBundled   bool   `json:"editionBundled,omitempty"`
		Key              string `json:"key,omitempty"`
		License          string `json:"license,omitempty"`
		Name             string `json:"name,omitempty"`
		OrganizationName string `json:"organizationName,omitempty"`
		OrganizationUrl  string `json:"organizationUrl,omitempty"`
		Release          struct {
			Date    string `json:"date,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"release,omitempty"`
		TermsAndConditionsUrl string `json:"termsAndConditionsUrl,omitempty"`
		Update                struct {
			Requires []string `json:"requires,omitempty"`
			Status   string   `json:"status,omitempty"`
		} `json:"update,omitempty"`
	} `json:"plugins,omitempty"`
	UpdateCenterRefresh string `json:"updateCenterRefresh,omitempty"`
}

// CancelAllRequest Cancels any operation pending on any plugin (install, update or uninstall)<br/>Requires user to be authenticated with Administer System permissions
type CancelAllRequest struct{}

// InstallRequest Installs the latest version of a plugin specified by its key.<br/>Plugin information is retrieved from Update Center.<br/>Fails if used on commercial editions or plugin risk consent has not been accepted.<br/>Requires user to be authenticated with Administer System permissions
type InstallRequest struct {
	Key string `json:"key"` // The key identifying the plugin to install
}

// InstalledRequest Get the list of all the plugins installed on the SonarQube instance, sorted by plugin name.<br/>Requires authentication.
type InstalledRequest struct {
	F string `url:"f,omitempty"` // Since 5.6;Comma-separated list of the additional fields to be returned in response. No additional field is returned by default. Possible values are:<ul><li>category - category as defined in the Update Center. A connection to the Update Center is needed</li></ul>
}

// InstalledResponse is the response for InstalledRequest
type InstalledResponse struct {
	Plugins []struct {
		Description          string   `json:"description,omitempty"`
		DocumentationPath    string   `json:"documentationPath,omitempty"`
		EditionBundled       bool     `json:"editionBundled,omitempty"`
		Filename             string   `json:"filename,omitempty"`
		Hash                 string   `json:"hash,omitempty"`
		HomepageUrl          string   `json:"homepageUrl,omitempty"`
		ImplementationBuild  string   `json:"implementationBuild,omitempty"`
		IssueTrackerUrl      string   `json:"issueTrackerUrl,omitempty"`
		Key                  string   `json:"key,omitempty"`
		License              string   `json:"license,omitempty"`
		Name                 string   `json:"name,omitempty"`
		OrganizationName     string   `json:"organizationName,omitempty"`
		OrganizationUrl      string   `json:"organizationUrl,omitempty"`
		RequiredForLanguages []string `json:"requiredForLanguages,omitempty"`
		SonarLintSupported   bool     `json:"sonarLintSupported,omitempty"`
		UpdatedAt            float64  `json:"updatedAt,omitempty"`
		Version              string   `json:"version,omitempty"`
	} `json:"plugins,omitempty"`
}

// PendingRequest Get the list of plugins which will either be installed or removed at the next startup of the SonarQube instance, sorted by plugin name.<br/>Require 'Administer System' permission.
type PendingRequest struct{}

// PendingResponse is the response for PendingRequest
type PendingResponse struct {
	Installing []struct {
		Category            string `json:"category,omitempty"`
		Description         string `json:"description,omitempty"`
		DocumentationPath   string `json:"documentationPath,omitempty"`
		HomepageUrl         string `json:"homepageUrl,omitempty"`
		ImplementationBuild string `json:"implementationBuild,omitempty"`
		IssueTrackerUrl     string `json:"issueTrackerUrl,omitempty"`
		Key                 string `json:"key,omitempty"`
		License             string `json:"license,omitempty"`
		Name                string `json:"name,omitempty"`
		OrganizationName    string `json:"organizationName,omitempty"`
		OrganizationUrl     string `json:"organizationUrl,omitempty"`
		Version             string `json:"version,omitempty"`
	} `json:"installing,omitempty"`
	Removing []struct {
		Category            string `json:"category,omitempty"`
		Description         string `json:"description,omitempty"`
		DocumentationPath   string `json:"documentationPath,omitempty"`
		HomepageUrl         string `json:"homepageUrl,omitempty"`
		ImplementationBuild string `json:"implementationBuild,omitempty"`
		IssueTrackerUrl     string `json:"issueTrackerUrl,omitempty"`
		Key                 string `json:"key,omitempty"`
		License             string `json:"license,omitempty"`
		Name                string `json:"name,omitempty"`
		OrganizationName    string `json:"organizationName,omitempty"`
		OrganizationUrl     string `json:"organizationUrl,omitempty"`
		Version             string `json:"version,omitempty"`
	} `json:"removing,omitempty"`
	Updating []struct {
		Category            string `json:"category,omitempty"`
		Description         string `json:"description,omitempty"`
		HomepageUrl         string `json:"homepageUrl,omitempty"`
		ImplementationBuild string `json:"implementationBuild,omitempty"`
		IssueTrackerUrl     string `json:"issueTrackerUrl,omitempty"`
		Key                 string `json:"key,omitempty"`
		License             string `json:"license,omitempty"`
		Name                string `json:"name,omitempty"`
		OrganizationName    string `json:"organizationName,omitempty"`
		OrganizationUrl     string `json:"organizationUrl,omitempty"`
		Version             string `json:"version,omitempty"`
	} `json:"updating,omitempty"`
}

// UninstallRequest Uninstalls the plugin specified by its key.<br/>Requires user to be authenticated with Administer System permissions.
type UninstallRequest struct {
	Key string `json:"key"` // The key identifying the plugin to uninstall
}

// UpdateRequest Updates a plugin specified by its key to the latest version compatible with the SonarQube instance.<br/>Plugin information is retrieved from Update Center.<br/>Requires user to be authenticated with Administer System permissions
type UpdateRequest struct {
	Key string `json:"key"` // The key identifying the plugin to update
}

// UpdatesRequest Lists plugins installed on the SonarQube instance for which at least one newer version is available, sorted by plugin name.<br/>Each newer version is listed, ordered from the oldest to the newest, with its own update/compatibility status.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: [COMPATIBLE, INCOMPATIBLE, REQUIRES_UPGRADE, DEPS_REQUIRE_UPGRADE].<br/>Require 'Administer System' permission.
type UpdatesRequest struct{}

// UpdatesResponse is the response for UpdatesRequest
type UpdatesResponse struct {
	Plugins []struct {
		Category              string `json:"category,omitempty"`
		Description           string `json:"description,omitempty"`
		EditionBundled        bool   `json:"editionBundled,omitempty"`
		Key                   string `json:"key,omitempty"`
		License               string `json:"license,omitempty"`
		Name                  string `json:"name,omitempty"`
		OrganizationName      string `json:"organizationName,omitempty"`
		OrganizationUrl       string `json:"organizationUrl,omitempty"`
		TermsAndConditionsUrl string `json:"termsAndConditionsUrl,omitempty"`
		Updates               []struct {
			Release struct {
				ChangeLogUrl string `json:"changeLogUrl,omitempty"`
				Date         string `json:"date,omitempty"`
				Description  string `json:"description,omitempty"`
				Version      string `json:"version,omitempty"`
			} `json:"release,omitempty"`
			Requires []string `json:"requires,omitempty"`
			Status   string   `json:"status,omitempty"`
		} `json:"updates,omitempty"`
	} `json:"plugins,omitempty"`
}
