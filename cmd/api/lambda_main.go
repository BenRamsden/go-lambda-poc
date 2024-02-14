//go:build lambda

package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/jugo-io/go-poc/internal/api"
	"github.com/jugo-io/go-poc/internal/api/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	repo := sql.NewSQLRepository(db)

	options := api.HandlerOptions{
		Repo: repo,
	}

	adapter := ginadapter.NewV2(api.Handler(options))
	lambda.Start(adapter.ProxyWithContext)
}
