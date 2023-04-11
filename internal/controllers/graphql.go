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
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/pkg/log"
	"runtime/debug"
)

type GraphqlController Controller

func (g *GraphqlController) buildCoinfig() graph.Config {
	c := graph.Config{
		Resolvers: resolvers.NewResolver(db.DB)
	}

	c.Directives.IsAdmin = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		auth := auth2.FromCtx(ctx)

		if auth.Role != string(user.ADMIN_ROLE) {
			return nil, fmt.Errorf("Access denied")
		}

		return next(ctx), nil
	}

	return c
}

func (g *GraphqlController) Query(ctx *gin.Context) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(g.buildCoinfig()))
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
