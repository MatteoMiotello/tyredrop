package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"pillowww/titw/graph"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
)

// User is the resolver for the user field.
func (r *cartResolver) User(ctx context.Context, obj *model.Cart) (*model.User, error) {
	userModel, err := r.UserDao.FindOneById(ctx, obj.UserID)

	if err != nil {
		return nil, err
	}

	return converters.UserToGraphQL(userModel), nil
}

// ProductItem is the resolver for the productItem field.
func (r *cartResolver) ProductItem(ctx context.Context, obj *model.Cart) (*model.ProductItem, error) {
	itemModel, err := r.ProductItemDao.FindProductItemById(ctx, obj.ProductItemID)

	if err != nil {
		return nil, err
	}

	return converters.ProductItemToGraphQL(itemModel), err
}

// Cart returns graph.CartResolver implementation.
func (r *Resolver) Cart() graph.CartResolver { return &cartResolver{r} }

type cartResolver struct{ *Resolver }