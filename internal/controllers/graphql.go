package controllers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"pillowww/titw/graph"
	"pillowww/titw/graph/resolvers"
)

type GraphqlController Controller

func (g *GraphqlController) Query(ctx *gin.Context) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolvers.Resolver{}}))
	srv.ServeHTTP(ctx.Writer, ctx.Request)
}

func (g *GraphqlController) Playground(ctx *gin.Context) {
	srv := playground.Handler("GraphQL playground", "/query")
	srv.ServeHTTP(ctx.Writer, ctx.Request)
}
