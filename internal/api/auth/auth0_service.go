package auth

import (
	"context"
	"fmt"
)

type Auth0User struct {
	ID string
}

func GetUser(ctx context.Context) {
	fmt.Println("%v", ctx)
}
