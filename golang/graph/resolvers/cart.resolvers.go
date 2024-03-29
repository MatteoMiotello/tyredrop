package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"pillowww/titw/graph"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// User is the resolver for the user field.
func (r *cartResolver) User(ctx context.Context, obj *model.Cart) (*model.User, error) {
	userModel, err := r.UserDao.FindOneById(ctx, obj.UserID)

	if err != nil {
		return nil, err
	}

	return converters.UserToGraphQL(userModel), nil
}

// ProductItemPrice is the resolver for the productItemPrice field.
func (r *cartResolver) ProductItemPrice(ctx context.Context, obj *model.Cart) (*model.ProductItemPrice, error) {
	lang := auth2.CurrentLanguage(ctx)
	dbModel, err := r.ProductItemPriceDao.
		Load(
			qm.Rels(
				models.ProductItemPriceRels.Currency,
				models.CurrencyRels.CurrencyLanguages,
			),
			models.CurrencyLanguageWhere.LanguageID.EQ(lang.L.ID),
		).
		FindOneById(ctx, obj.ProductItemPriceID)

	if err != nil {
		return nil, err
	}

	return converters.ProductItemPriceToGraphQL(dbModel)
}

// Cart returns graph.CartResolver implementation.
func (r *Resolver) Cart() graph.CartResolver { return &cartResolver{r} }

type cartResolver struct{ *Resolver }
