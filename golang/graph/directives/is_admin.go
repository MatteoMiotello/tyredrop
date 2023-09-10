package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/domain/user"
)

func IsAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	auth := auth2.FromCtx(ctx)

	if auth.Role.Code != string(user.ADMIN_ROLE) {
		return nil, fmt.Errorf("Access denied")
	}

	return next(ctx)
}
