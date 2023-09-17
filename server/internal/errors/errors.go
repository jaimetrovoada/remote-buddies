package errors

type AuthError struct {
	Message string
	Err     error
}

func (e *AuthError) Error() string {
	return e.Message
}
