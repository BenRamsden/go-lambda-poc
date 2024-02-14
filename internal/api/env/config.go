package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	AUTH0_AUDIENCE      string = ""
	AUTH0_CLIENT_ID     string = ""
	AUTH0_DOMAIN        string = ""
	AUTH0_ISSUER_URL    string = ""
	AUTH0_WEB_CLIENT_ID string = ""
	AUTH0_CLIENT_SECRET string = ""
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	AUTH0_AUDIENCE = os.Getenv("AUTH0_AUDIENCE")
	AUTH0_CLIENT_ID = os.Getenv("AUTH0_CLIENT_ID")
	AUTH0_DOMAIN = os.Getenv("AUTH0_DOMAIN")
	AUTH0_ISSUER_URL = os.Getenv("AUTH0_ISSUER_URL")
	AUTH0_WEB_CLIENT_ID = os.Getenv("AUTH0_WEB_CLIENT_ID")
	AUTH0_CLIENT_SECRET = os.Getenv("AUTH0_CLIENT_SECRET")
}
