package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

func EmptyStringToNull(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
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
