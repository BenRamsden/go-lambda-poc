package main

import (
	"context"
	"github.com/jugo-io/go-poc/internal/api"
	"github.com/jugo-io/go-poc/internal/api/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// TODO: Once pulumi ENV / Secrets Manager is set up, re-add the panic
		//panic("failed to connect database")
	}

	repo := sql.NewSQLRepository(db)

	options := api.HandlerOptions{
		Repo: repo,
	}
	r := api.Handler(options)
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
