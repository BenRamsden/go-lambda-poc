package model

type Auth struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

const (
	ErrNotAuthenticated = "not authenticated"
)

type AuthService interface {
	// Find or create a new auth
	// If the user cannot be authenticated ErrNotAuthenticated will be returned
	FindOrCreate(jwt string) (Auth, error)
}
