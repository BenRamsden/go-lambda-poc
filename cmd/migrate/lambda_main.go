//go:build lambda

package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jugo-io/go-poc/internal/api/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func HandleRequest(ctx context.Context) error {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	repo := sql.NewSQLRepository(db)
	if err := repo.Migrate(); err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
