package helpers

import "fmt"

type HTTPError struct {
	Code    int
	Message string
}

func NewHTTPError(code int, message string) error {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

func (err *HTTPError) Error() string {
	return fmt.Sprintf("status %d: %s", err.Code, err.Message)
}
