package analysis_cache

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// GetRequest Get the scanner's cached data for a branch. Requires scan permission on the project. Data is returned gzipped if the corresponding 'Accept-Encoding' header is set in the request.
type GetRequest struct {
	Branch  string `form:"branch,omitempty"`  // Branch key. If not provided, main branch will be used.
	Project string `form:"project,omitempty"` // Project key
}
