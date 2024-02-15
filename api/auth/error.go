package auth

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNotAuthenticated Error = "not authenticated"
	ErrNotAuthorized    Error = "not authorized"
)
