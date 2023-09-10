package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"pillowww/titw/graph/graphErrors"
	auth2 "pillowww/titw/internal/auth"
)

func UserConfirmed(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	u, err := auth2.CurrentUser(ctx)

	if err != nil {
		return nil, graphErrors.NewUserNotFoundError(ctx, err)
	}

	if !u.Confirmed {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	if u.Rejected {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	return next(ctx)
}
