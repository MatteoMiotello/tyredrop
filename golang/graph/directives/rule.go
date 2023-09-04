package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
	"pillowww/titw/graph/graphErrors"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func Rule(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}

	err = validate.Var(val, constraint)
	if err != nil {
		return nil, graphErrors.NewGraphError(ctx, err, "FIELD_VALIDATION")
	}

	return val, nil
}
