//go:build !lambda

package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/dynamo"
	"github.com/jugo-io/go-poc/api/graph"
	"github.com/jugo-io/go-poc/api/service"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

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

	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/graphql")))
	r.POST("/graphql", auth.EnsureValidToken(), gin.WrapH(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))))

	r.Run(":4000")
}
