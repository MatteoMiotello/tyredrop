package controllers

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"os"
	"pillowww/titw/graph"
	"pillowww/titw/graph/resolvers"
	"pillowww/titw/internal/db"
	"pillowww/titw/pkg/log"
	"runtime/debug"
)

type GraphqlController Controller

func (g *GraphqlController) Query(ctx *gin.Context) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolvers.NewResolver(db.DB)}))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		log.
			WithField("path", graphql.GetPath(ctx)).
			Error(err, os.Stderr)

		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr)
		debug.PrintStack()

		return gqlerror.Errorf("internal system error")
	})
	srv.ServeHTTP(ctx.Writer, ctx.Request)
}

func (g *GraphqlController) Playground(ctx *gin.Context) {
	srv := playground.Handler("GraphQL playground", "/query")
	srv.ServeHTTP(ctx.Writer, ctx.Request)
}
