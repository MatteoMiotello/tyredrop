package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/friendsofgo/errors"
	"pillowww/titw/graph/graphErrors"
	"strings"
)

func NotEmpty(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
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
