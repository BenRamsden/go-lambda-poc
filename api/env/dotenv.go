//go:build !lambda

package env

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	AUTH0_AUDIENCE = os.Getenv("AUTH0_AUDIENCE")
	AUTH0_DOMAIN = os.Getenv("AUTH0_DOMAIN")
}
