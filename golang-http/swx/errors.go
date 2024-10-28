package swx

import "fmt"

// ErrorBody contains the information about the error response.
type ErrorBody struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ErrorResponse defines the structure of the error response used by the
// SmartWorks API.
type ErrorResponse struct {
	Err ErrorBody `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Err.Status, e.Err.Message)
}

type ResponseStatus struct {
	Status int
}

// OAuth2Error defines an error response returned by the OAuth2 server.
type OAuth2Error struct {
	ErrorMessage     string         `json:"error"`
	ErrorDescription string         `json:"error_description"`
	Err              ResponseStatus `json:"-"` // For compatibility reasons
}

func (e *OAuth2Error) Error() string {
	return e.ErrorMessage + " -> " + e.ErrorDescription
}

// TokenRevokeError defines an error response returned when trying to revoke
// a Token that was not obtained successfully.
type TokenRevokeError struct {
	Message string
}

func (e *TokenRevokeError) Error() string {
	return e.Message
}
