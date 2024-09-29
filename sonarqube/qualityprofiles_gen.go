package sonarqube

import (
	"context"
	"fmt"
	"github.com/shijl0925/go-sonarqube/sonarqube/paging"
	"github.com/shijl0925/go-sonarqube/sonarqube/qualityprofiles"
	"net/http"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Qualityprofiles service

// ActivateRule - Activate a rule on a Quality Profile.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//
// Since 4.4
// Changelog:
//
//	10.6: Add parameter 'prioritizedRule'.
//	10.2: Parameter 'severity' is now deprecated.
func (s *Qualityprofiles) ActivateRule(ctx context.Context, r qualityprofiles.ActivateRuleRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/activate_rule", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ActivateRules - Bulk-activate rules on one quality profile.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//
// Since 4.4
// Changelog:
//
//	10.6: Add parameter 'prioritizedRule'.
//	10.2: Parameters 'severities', 'targetSeverity', 'active_severities', and 'types' are now deprecated.
//	10.0: Parameter 'sansTop25' is deprecated
func (s *Qualityprofiles) ActivateRules(ctx context.Context, r qualityprofiles.ActivateRulesRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/activate_rules", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// AddProject - Associate a project with a quality profile.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Administer right on the specified project
//
// Since 5.2
func (s *Qualityprofiles) AddProject(ctx context.Context, r qualityprofiles.AddProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/add_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Backup - Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore.
// Since 5.2
// Changelog:
//
//	10.3: The 'priority' and 'type' fields of the rule XML object are deprecated.
func (s *Qualityprofiles) Backup(ctx context.Context, r qualityprofiles.BackupRequest) (*qualityprofiles.BackupResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/backup", s.path)
	v := new(qualityprofiles.BackupResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// ChangeParent - Change a quality profile's parent.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//
// Since 5.2
func (s *Qualityprofiles) ChangeParent(ctx context.Context, r qualityprofiles.ChangeParentRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/change_parent", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Changelog - Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first).
// Since 5.2
// Changelog:
//
//	10.3: Added fields 'cleanCodeAttributeCategory', 'impacts' to response
//	10.3: Added fields 'oldCleanCodeAttribute', 'newCleanCodeAttribute', 'oldCleanCodeAttributeCategory', 'newCleanCodeAttributeCategory' and 'impactChanges' to 'params' section of response
//	10.3: Added field 'sonarQubeVersion' to 'params' section of response
//	9.8: response fields 'total', 's', 'ps' have been deprecated, please use 'paging' object instead
//	9.8: The field 'paging' has been added to the response
func (s *Qualityprofiles) Changelog(ctx context.Context, r qualityprofiles.ChangelogRequest, p paging.Params) (*qualityprofiles.ChangelogResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/changelog", s.path)
	v := new(qualityprofiles.ChangelogResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Qualityprofiles) ChangelogAll(ctx context.Context, r qualityprofiles.ChangelogRequest) (*qualityprofiles.ChangelogResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualityprofiles.ChangelogResponseAll{}
	for {
		res, _, err := s.Changelog(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to qualityprofiles.Changelog: %+v", err)
		}
		response.Events = append(response.Events, res.Events...)
		if res.GetPaging().End() {
			break
		}
		p.P++
	}
	return response, nil
}

// Copy - Copy a quality profile.
// Requires to be logged in and the 'Administer Quality Profiles' permission.
// Since 5.2
func (s *Qualityprofiles) Copy(ctx context.Context, r qualityprofiles.CopyRequest) (*qualityprofiles.CopyResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/copy", s.path)
	v := new(qualityprofiles.CopyResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Create - Create a quality profile.
// Requires to be logged in and the 'Administer Quality Profiles' permission.
// Since 5.2
func (s *Qualityprofiles) Create(ctx context.Context, r qualityprofiles.CreateRequest) (*qualityprofiles.CreateResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/create", s.path)
	v := new(qualityprofiles.CreateResponse)

	resp, err := s.client.Call(ctx, "POST", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// DeactivateRule - Deactivate a rule on a quality profile.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//
// Since 4.4
// Changelog:
//
//	10.3: Inherited rules can be deactivated (if the global admin setting is enabled)
func (s *Qualityprofiles) DeactivateRule(ctx context.Context, r qualityprofiles.DeactivateRuleRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/deactivate_rule", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeactivateRules - Bulk deactivate rules on Quality profiles.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//
// Since 4.4
// Changelog:
//
//	10.3: Inherited rules can be deactivated (if the global admin setting is enabled)
//	10.2: Parameters 'severities', 'active_severities', and 'types' are now deprecated.
//	10.0: Parameter 'sansTop25' is deprecated
func (s *Qualityprofiles) DeactivateRules(ctx context.Context, r qualityprofiles.DeactivateRulesRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/deactivate_rules", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Delete - Delete a quality profile and all its descendants. The default quality profile cannot be deleted.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//
// Since 5.2
func (s *Qualityprofiles) Delete(ctx context.Context, r qualityprofiles.DeleteRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/delete", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Export - Export a quality profile.
// Since 5.2
func (s *Qualityprofiles) Export(ctx context.Context, r qualityprofiles.ExportRequest) (*qualityprofiles.ExportResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/export", s.path)
	v := new(qualityprofiles.ExportResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Exporters - Lists available profile export formats.
// Since 5.2
func (s *Qualityprofiles) Exporters(ctx context.Context, r qualityprofiles.ExportersRequest) (*qualityprofiles.ExportersResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/exporters", s.path)
	v := new(qualityprofiles.ExportersResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Importers - List supported importers.
// Since 5.2
func (s *Qualityprofiles) Importers(ctx context.Context, r qualityprofiles.ImportersRequest) (*qualityprofiles.ImportersResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/importers", s.path)
	v := new(qualityprofiles.ImportersResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Inheritance - Show a quality profile's ancestors and children.
// Since 5.2
// Changelog:
//
//	10.3: Field 'inactiveRuleCount' added to the response
func (s *Qualityprofiles) Inheritance(ctx context.Context, r qualityprofiles.InheritanceRequest) (*qualityprofiles.InheritanceResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/inheritance", s.path)
	v := new(qualityprofiles.InheritanceResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Projects - List projects with their association status regarding a quality profile.
// Only projects explicitly bound to the profile are returned, those associated with the profile because it is the default one are not.
// See api/qualityprofiles/search in order to get the Quality Profile Key.
// Since 5.2
// Changelog:
//
//	10.0: deprecated 'more' response field has been removed
//	8.8: deprecated 'id' response field has been removed
//	8.8: deprecated 'uuid' response field has been removed
//	7.2: 'more' response field is deprecated
//	6.5: 'id' response field is deprecated
//	6.0: 'uuid' response field is deprecated and replaced by 'id'
//	6.0: 'key' response field has been added to return the project key
func (s *Qualityprofiles) Projects(ctx context.Context, r qualityprofiles.ProjectsRequest, p paging.Params) (*qualityprofiles.ProjectsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/projects", s.path)
	v := new(qualityprofiles.ProjectsResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r, p)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

func (s *Qualityprofiles) ProjectsAll(ctx context.Context, r qualityprofiles.ProjectsRequest) (*qualityprofiles.ProjectsResponseAll, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	response := &qualityprofiles.ProjectsResponseAll{}
	for {
		res, _, err := s.Projects(ctx, r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to qualityprofiles.Projects: %+v", err)
		}
		response.Results = append(response.Results, res.Results...)
		if res.GetPaging().End() {
			break
		}
		p.P++
	}
	return response, nil
}

// RemoveProject - Remove a project's association with a quality profile.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//    * Administer right on the specified project
//
// Since 5.2
func (s *Qualityprofiles) RemoveProject(ctx context.Context, r qualityprofiles.RemoveProjectRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/remove_project", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Rename - Rename a quality profile.
// Requires one of the following permissions:
//    * 'Administer Quality Profiles'
//    * Edit right on the specified quality profile
//
// Since 5.2
func (s *Qualityprofiles) Rename(ctx context.Context, r qualityprofiles.RenameRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/rename", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Restore - Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.
// Requires to be logged in and the 'Administer Quality Profiles' permission.
// Since 5.2
// Changelog:
//
//	10.3: The 'priority' and 'type' fields of the rule XML object are deprecated.
func (s *Qualityprofiles) Restore(ctx context.Context, r qualityprofiles.RestoreRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/restore", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Search - Search quality profiles
// Since 5.2
// Changelog:
//
//	10.0: Remove deprecated parameter 'project_key'. Please use 'project' instead.
//	7.0: Add available actions 'delete' and 'associateProjects'
//	6.6: Add available actions 'edit', 'copy' and 'setAsDefault' and global action 'create'
//	6.5: The parameters 'defaults', 'project' and 'language' can be combined without any constraint
func (s *Qualityprofiles) Search(ctx context.Context, r qualityprofiles.SearchRequest) (*qualityprofiles.SearchResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/search", s.path)
	v := new(qualityprofiles.SearchResponse)

	resp, err := s.client.Call(ctx, "GET", u, v, r)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// SetDefault - Select the default profile for a given language.
// Requires to be logged in and the 'Administer Quality Profiles' permission.
// Since 5.2
func (s *Qualityprofiles) SetDefault(ctx context.Context, r qualityprofiles.SetDefaultRequest) (*http.Response, error) {
	u := fmt.Sprintf("%s/set_default", s.path)

	resp, err := s.client.Call(ctx, "POST", u, nil, r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
