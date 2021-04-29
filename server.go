package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cadazab/gqlgenServer/graph"
	"github.com/cadazab/gqlgenServer/graph/generated"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "8080")
		port = "8080"
	}
	// Setting up Gin
	r := gin.New()
	r.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowOrigins = []string{"http://localhost", "https://routify-app.vercel.app/"}
	r.Use(cors.Default())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}
