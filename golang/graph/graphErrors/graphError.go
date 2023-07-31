package graphErrors

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/friendsofgo/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func NewGraphError(ctx context.Context, err error, code string) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

func NewUserNotFoundError(ctx context.Context, err error) *gqlerror.Error {
	return NewGraphError(ctx, errors.WithMessage(err, "User not found in context"), "4004")
}

func NewNotAuthorizedError(ctx context.Context) *gqlerror.Error {
	return NewGraphError(ctx, errors.New("User not authorized to perform this action"), "4003")
}
