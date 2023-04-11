package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"pillowww/titw/graph"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/auth"
	"pillowww/titw/models"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int64) (*model.User, error) {
	user, err := r.UserDao.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	}

	return converters.UserToGraphQL(user), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// TaxRates is the resolver for the taxRates field.
func (r *queryResolver) TaxRates(ctx context.Context) ([]*model.TaxRate, error) {
	panic(fmt.Errorf("not implemented: TaxRates - taxRates"))
}

// LegalEntityTypes is the resolver for the legalEntityTypes field.
func (r *queryResolver) LegalEntityTypes(ctx context.Context) ([]*model.LegalEntityType, error) {
	panic(fmt.Errorf("not implemented: LegalEntityTypes - legalEntityTypes"))
}

// ProductsItemsByCode is the resolver for the productsItemsByCode field.
func (r *queryResolver) ProductsItemsByCode(ctx context.Context, code string) (*model.ProductItem, error) {
	pItem, err := r.ProductItemDao.FindLessExpensiveByProductCode(ctx, code)

	if err != nil {
		return nil, err
	}

	return converters.ProductItemToGraphQL(pItem), nil
}

// ProductItems is the resolver for the productItems field.
func (r *queryResolver) ProductItems(ctx context.Context, input []*model.ProductSpecificationInput) ([]*model.ProductItem, error) {
	panic(fmt.Errorf("not implemented: ProductItems - productItems"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, pagination *model.PaginationInput) (*model.ProductPaginate, error) {
	count, err := r.ProductDao.CountAll(ctx)
	countInt := int(count)
	if err != nil {
		return nil, err
	}

	dao := r.ProductDao

	if pagination != nil {
		dao = r.ProductDao.Paginate(pagination.Limit, pagination.Offset)
	}

	products, err := dao.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var graphModels []*model.Product

	for _, product := range products {
		graphModels = append(graphModels, converters.ProductToGraphQL(product))
	}

	return &model.ProductPaginate{
		Products: graphModels,
		Pagination: &model.Pagination{
			Offset: &pagination.Offset,
			Limit:  &pagination.Limit,
			Totals: &countInt,
		},
	}, nil
}

// Currency is the resolver for the currency field.
func (r *queryResolver) Currency(ctx context.Context, id int64) (*model.Currency, error) {
	lang := auth.CurrentLanguage(ctx)

	c, err := r.CurrencyDao.
		Load(models.CurrencyRels.CurrencyLanguages, models.CurrencyLanguageWhere.LanguageID.EQ(lang.L.ID)).
		FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return converters.CurrencyToGraphQL(c), nil
}

// Currencies is the resolver for the currencies field.
func (r *queryResolver) Currencies(ctx context.Context) ([]*model.Currency, error) {
	leng := auth.CurrentLanguage(ctx)

	allCurrencies, err := r.CurrencyDao.
		Load(models.CurrencyRels.CurrencyLanguages, models.CurrencyLanguageWhere.LanguageID.EQ(leng.L.ID)).
		FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var cAll []*model.Currency

	for _, c := range allCurrencies {
		cAll = append(cAll, converters.CurrencyToGraphQL(c))
	}

	return cAll, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
