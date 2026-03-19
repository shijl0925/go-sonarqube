package sonarqube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ErrorResponse represents an error returned by the SonarQube API.
// It implements the error interface.
type ErrorResponse struct {
	// Response is the HTTP response that caused this error.
	Response *http.Response `json:"-"`
	// StatusCode is the HTTP status code returned by the server.
	StatusCode int
	// Errors contains the individual error messages from the API.
	Errors []struct {
		Msg string `json:"msg"`
	} `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	messages := make([]string, len(e.Errors))
	for i := 0; i < len(e.Errors); i++ {
		messages[i] = e.Errors[i].Msg
	}
	messagesString := strings.Join(messages, ", ")

	if e.Response != nil {
		return fmt.Sprintf("%s %s: %d %s",
			e.Response.Request.Method,
			e.Response.Request.URL,
			e.StatusCode,
			messagesString,
		)
	}
	return fmt.Sprintf("received non 2xx status code (%d): %s", e.StatusCode, messagesString)
}

func ErrorResponseFrom(res *http.Response) (*ErrorResponse, error) {
	errorResponse := &ErrorResponse{}
	err := json.NewDecoder(res.Body).Decode(&errorResponse)
	if err != nil {
		return nil, fmt.Errorf("could not decode response into ErrorResponse: %+v", err)
	}
	errorResponse.StatusCode = res.StatusCode
	errorResponse.Response = res
	return errorResponse, nil
}
