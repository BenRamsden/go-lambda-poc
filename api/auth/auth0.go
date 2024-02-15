package auth

import (
	"context"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type Auth struct {
	ID string
}

func GetUser(ctx context.Context) (Auth, error) {
	val := ctx.Value(jwtmiddleware.ContextKey{})
	if val == nil {
		return Auth{}, ErrNotAuthenticated
	}
	token := val.(*validator.ValidatedClaims)
	return Auth{ID: token.RegisteredClaims.Subject}, nil
}
