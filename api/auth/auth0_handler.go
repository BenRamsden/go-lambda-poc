package auth

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/jugo-io/go-poc/api/env"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) HasScope(expectedScope string) bool {
	result := c.GetScopes()
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

func (c CustomClaims) GetScopes() []string {
	return strings.Split(c.Scope, " ")
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken() gin.HandlerFunc {

	issuerURL, err := url.Parse(env.AUTH0_DOMAIN)
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{env.AUTH0_AUDIENCE},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{
					// Scopes: []string{"read:users", "create:users", "update:users", "read:users_app_metadata", "update:users_app_metadata", "create:users_app_metadata", "create:user_tickets"},
				}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("\n\nEncountered error while validating JWT:\n\t%v\n\n", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return adapter.Wrap(middleware.CheckJWT)
}
