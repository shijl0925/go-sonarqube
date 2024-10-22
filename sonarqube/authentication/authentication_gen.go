package authentication

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// LoginRequest Authenticate a user.
type LoginRequest struct {
	Login    string `form:"login"`    // Login of the user
	Password string `form:"password"` // Password of the user
}

// LogoutRequest Logout a user.
type LogoutRequest struct{}

// ValidateRequest Check credentials.
type ValidateRequest struct{}

// ValidateResponse is the response for ValidateRequest
type ValidateResponse struct {
	Valid bool `json:"valid,omitempty"`
}
