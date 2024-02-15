package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugo-io/go-poc/api/auth"
	"github.com/jugo-io/go-poc/api/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type HandlerOptions struct {
}

func Handler(options HandlerOptions) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	resolver := &graph.Resolver{}

	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/graphql")))
	r.POST("/graphql", auth.EnsureValidToken(), gin.WrapH(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))))

	return r
}
