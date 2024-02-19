//go:build lambda

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/dynamo"
	"github.com/jugo-io/go-poc/api/graph"
	"github.com/jugo-io/go-poc/api/service"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")

	r := gin.Default()

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Debug:            false,
		AttachStacktrace: true,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	} else {
		log.Println("Sentry initialized")
	}
	r.Use(sentrygin.New(sentrygin.Options{}))
	// Sentry Flush with lambda?

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

	graphSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	graphSrv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		sentry.CaptureException(err)
		return err
	})

	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/graphql")))
	r.POST("/graphql", auth.EnsureValidToken(), gin.WrapH(graphSrv))

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
