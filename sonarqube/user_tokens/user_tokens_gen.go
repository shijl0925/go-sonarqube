package user_tokens

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// GenerateRequest Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />It requires administration permissions to specify a 'login' and generate a token for another user. Otherwise, a token is generated for the current user.
type GenerateRequest struct {
	ExpirationDate string `json:"expirationDate,omitempty"` // Since 9.6;The expiration date of the token being generated, in ISO 8601 format (YYYY-MM-DD). If not set, default to no expiration.
	Login          string `json:"login,omitempty"`          // User login. If not set, the token is generated for the authenticated user.
	Name           string `json:"name"`                     // Token name
	ProjectKey     string `json:"projectKey,omitempty"`     // Since 9.5;The key of the only project that can be analyzed by the PROJECT_ANALYSIS_TOKEN being generated.
	Type           string `json:"type,omitempty"`           // Since 9.5;Token Type. If this parameters is set to PROJECT_ANALYSIS_TOKEN, it is necessary to provide the projectKey parameter too.
}

// GenerateResponse is the response for GenerateRequest
type GenerateResponse struct {
	CreatedAt      string `json:"createdAt,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	Login          string `json:"login,omitempty"`
	Name           string `json:"name,omitempty"`
	Token          string `json:"token,omitempty"`
	Type           string `json:"type,omitempty"`
}

// RevokeRequest Revoke a user access token. <br/>It requires administration permissions to specify a 'login' and revoke a token for another user. Otherwise, the token for the current user is revoked.
type RevokeRequest struct {
	Login string `json:"login,omitempty"` // User login
	Name  string `json:"name"`            // Token name
}

// SearchRequest List the access tokens of a user.<br>The login must exist and active.<br>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user is using a token many times in less than one hour.<br> It requires administration permissions to specify a 'login' and list the tokens of another user. Otherwise, tokens for the current user are listed. <br> Authentication is required for this API endpoint
type SearchRequest struct {
	Login string `url:"login,omitempty"` // User login
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Login      string `json:"login,omitempty"`
	UserTokens []struct {
		CreatedAt      string `json:"createdAt,omitempty"`
		Name           string `json:"name,omitempty"`
		Type           string `json:"type,omitempty"`
		ExpirationDate string `json:"expirationDate,omitempty"`
		IsExpired      bool   `json:"isExpired,omitempty"`
		Project        struct {
			Key  string `json:"key,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"project,omitempty"`
	} `json:"userTokens,omitempty"`
}
