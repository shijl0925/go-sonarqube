package sources

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// RawRequest Get source code as raw text. Require 'See Source Code' permission on file
type RawRequest struct {
	Key string `form:"key,omitempty"` // File key
}

// RawResponse is the response for RawRequest
type RawResponse string

// ScmRequest Get SCM information of source files. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Author of the commit</li><li>Datetime of the commit (before 5.2 it was only the Date)</li><li>Revision of the commit (added in 5.2)</li></ol>
type ScmRequest struct {
	CommitsByLine string `form:"commits_by_line,omitempty"` // Group lines by SCM commit if value is false, else display commits for each line, even if two consecutive lines relate to the same commit.
	From          string `form:"from,omitempty"`            // First line to return. Starts at 1
	Key           string `form:"key,omitempty"`             // File key
	To            string `form:"to,omitempty"`              // Last line to return (inclusive)
}

// ScmResponse is the response for ScmRequest
type ScmResponse struct {
	Scm [][]float64 `json:"scm,omitempty"`
}

// ShowRequest Get source code. Requires See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Content of the line</li></ol>
type ShowRequest struct {
	From string `form:"from,omitempty"` // First line to return. Starts at 1
	Key  string `form:"key,omitempty"`  // File key
	To   string `form:"to,omitempty"`   // Last line to return (inclusive)
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Sources [][]float64 `json:"sources,omitempty"`
}
