//go:build lambda

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/dynamo"
	"github.com/jugo-io/go-poc/api/graph"
	"github.com/jugo-io/go-poc/api/service"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")

	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	repo := dynamo.NewFromEnv()
	assets := service.NewAssetService(repo)

	resolver := &graph.Resolver{
		AssetService: assets,
	}

	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/graphql")))
	r.POST("/graphql", auth.EnsureValidToken(), gin.WrapH(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))))

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
