//go:build !lambda

package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	AUTH0_AUDIENCE = os.Getenv("AUTH0_AUDIENCE")
	AUTH0_DOMAIN = os.Getenv("AUTH0_DOMAIN")
}
