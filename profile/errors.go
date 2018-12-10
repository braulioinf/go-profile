package profile

import (
	"encoding/json"
	"fmt"
)

// Error struct
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func parseError(e interface{}) (message string) {
	errors, _ := json.Marshal(e)
	err := string(errors)

	message = err
	return message
}

// ErrorUnauthorized interface
func ErrorUnauthorized(res interface{}) *Error {
	message := parseError(res)
	return &Error{
		Code:    400,
		Message: message,
	}
}

// ErrorInternalServer func
func ErrorInternalServer(res interface{}) *Error {
	message := parseError(res)
	return &Error{
		Code:    500,
		Message: message,
	}
}

// ErrorUnauthorized func
func (e *Error) Error() string {
	return fmt.Sprintf("The server has responded with Code: %d, Detail: %s", e.Code, e.Message)
}
