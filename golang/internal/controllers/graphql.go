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
	"pillowww/titw/graph/directives"
	"pillowww/titw/graph/resolvers"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/db"
	"pillowww/titw/pkg/log"
	"runtime/debug"
)

type GraphqlController Controller

func (g *GraphqlController) buildConfig() graph.Config {
	c := graph.Config{
		Resolvers: resolvers.NewResolver(db.DB),
	}

	c.Directives.IsAdmin = directives.IsAdmin
	c.Directives.NotEmpty = directives.NotEmpty
	c.Directives.EmptyStringToNull = directives.EmptyStringToNull
	c.Directives.UserConfirmed = directives.UserConfirmed
	c.Directives.Rule = directives.Rule

	return c
}

func (g *GraphqlController) Query(ctx *gin.Context) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(g.buildConfig()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		log.
			WithField("path", graphql.GetPath(ctx)).
			Error(err, os.Stderr)

		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr)
		debug.PrintStack()

		return gqlerror.Errorf("internal system error")
	})

	auth := auth2.FromCtx(ctx)

	c := context.WithValue(ctx, "auth", auth) //todo
	r := ctx.Request.WithContext(c)
	srv.ServeHTTP(ctx.Writer, r)
}

func (g *GraphqlController) Playground(ctx *gin.Context) {
	srv := playground.Handler("GraphQL playground", "/query")
	srv.ServeHTTP(ctx.Writer, ctx.Request)
}
