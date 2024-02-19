//go:build !lambda

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/dynamo"
	"github.com/jugo-io/go-poc/api/graph"
	"github.com/jugo-io/go-poc/api/service"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
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
	defer sentry.Flush(2 * time.Second)

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})) // Allow all origins

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	repo := dynamo.NewFromLocal()
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

	r.Run(":4000")
}
