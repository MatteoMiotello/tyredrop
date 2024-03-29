package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
)

// PriceMarkups is the resolver for the priceMarkups field.
func (r *queryResolver) PriceMarkups(ctx context.Context) ([]*model.ProductPriceMarkup, error) {
	pms, err := r.PriceMarkupDao.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var markups []*model.ProductPriceMarkup

	for _, pm := range pms {
		markups = append(markups, converters.PriceMarkupToGraphQL(pm))
	}

	return markups, nil
}
