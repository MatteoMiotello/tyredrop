package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"os"
	"pillowww/titw/graph"
	"pillowww/titw/graph/graphErrors"
	"pillowww/titw/graph/resolvers"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/pkg/log"
	"runtime/debug"
	"strings"
)

type GraphqlController Controller

func (g *GraphqlController) buildConfig() graph.Config {
	c := graph.Config{
		Resolvers: resolvers.NewResolver(db.DB),
	}

	c.Directives.IsAdmin = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		auth := auth2.FromCtx(ctx)

		if auth.Role.Code != string(user.ADMIN_ROLE) {
			return nil, fmt.Errorf("Access denied")
		}

		return next(ctx)
	}

	c.Directives.NotEmpty = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		val, err := next(ctx)
		str, ok := val.(string)

		graphContext := graphql.GetPathContext(ctx)
		fieldName := graphContext.Field

		if !ok {
			return nil, graphErrors.NewGraphError(ctx, errors.New("Not empty can be used only on strings."), "2002")
		}

		str = strings.Replace(str, " ", "", -1)

		if len(str) == 0 {
			return nil, graphErrors.NewGraphError(ctx, errors.New(fmt.Sprintf("%s field can not be empty.", *fieldName)), "2001")
		}

		return val, err
	}

	c.Directives.EmptyStringToNull = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		val, err := next(ctx)
		str, ok := val.(*string)

		if !ok {
			return val, nil
		}

		if len(*str) == 0 {
			return nil, nil
		}

		return val, err
	}

	c.Directives.UserConfirmed = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		u, err := auth2.CurrentUser(ctx)

		if err != nil {
			return nil, graphErrors.NewUserNotFoundError(ctx, err)
		}

		if !u.Confirmed {
			return nil, graphErrors.NewNotAuthorizedError(ctx)
		}

		return next(ctx)
	}

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
