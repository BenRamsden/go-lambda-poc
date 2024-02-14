package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	AUTH0_AUDIENCE string = ""
	AUTH0_DOMAIN   string = ""
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	AUTH0_AUDIENCE = os.Getenv("AUTH0_AUDIENCE")
	AUTH0_DOMAIN = os.Getenv("AUTH0_DOMAIN")
}
