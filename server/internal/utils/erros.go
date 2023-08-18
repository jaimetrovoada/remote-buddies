package utils

import "fmt"

type AuthError struct {
	Message string
	Code    int
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}
