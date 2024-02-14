package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugo-io/go-poc/internal/api/auth"
	"github.com/jugo-io/go-poc/internal/api/graph"
	"github.com/jugo-io/go-poc/internal/api/model"
	"github.com/jugo-io/go-poc/internal/api/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type HandlerOptions struct {
	Repo sql.Repository
}

func Handler(options HandlerOptions) *gin.Engine {
	if options.Repo == nil {
		panic("missing options.repo")
	}

	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	assetService := model.NewAssetService(options.Repo)

	resolver := &graph.Resolver{
		AssetService: assetService,
	}

	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/graphql")))
	r.POST("/graphql", auth.EnsureValidToken(), gin.WrapH(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))))

	return r
}
