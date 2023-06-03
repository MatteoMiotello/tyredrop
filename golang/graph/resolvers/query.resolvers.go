package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"errors"
	"fmt"
	"pillowww/titw/graph"
	"pillowww/titw/graph/aggregators"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/graphErrors"
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
func (r *queryResolver) TaxRates(ctx context.Context) ([]*model.Tax, error) {
	panic(fmt.Errorf("not implemented: TaxRates - taxRates"))
}

// Brands is the resolver for the brands field.
func (r *queryResolver) Brands(ctx context.Context) ([]*model.Brand, error) {
	brandModels, err := r.BrandDao.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var graphModels []*model.Brand

	for _, m := range brandModels {
		graphModels = append(graphModels, converters.BrandToGraphQL(m))
	}

	return graphModels, nil
}

// Carts is the resolver for the carts field.
func (r *queryResolver) Carts(ctx context.Context) (*model.CartResponse, error) {
	a := auth.FromCtx(ctx)
	user, err := a.GetUser(ctx)

	if err != nil {
		return nil, graphErrors.NewGraphError(ctx, errors.New("User not found in context"), "4004")
	}

	return aggregators.GetAllCartsByUserId(ctx, r.CartDao, user.ID)
}

// SearchBrands is the resolver for the searchBrands field.
func (r *queryResolver) SearchBrands(ctx context.Context, name string) ([]*model.Brand, error) {
	var graphModels []*model.Brand

	if len(name) < 2 {
		return graphModels, nil
	}

	models, err := r.BrandDao.FindByName(ctx, name)

	if err != nil {
		return nil, err
	}

	for _, model := range models {
		graphModels = append(graphModels, converters.BrandToGraphQL(model))
	}

	return graphModels, nil
}

// ProductCategories is the resolver for the productCategories field.
func (r *queryResolver) ProductCategories(ctx context.Context) ([]*model.ProductCategory, error) {
	defaultLang := auth.CurrentLanguage(ctx)

	dbModels, err := r.ProductCategoryDao.
		Load(
			models.ProductCategoryRels.ProductCategoryLanguages,
			models.ProductCategoryLanguageWhere.LanguageID.EQ(defaultLang.L.ID),
		).
		FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var graphModels []*model.ProductCategory

	for _, dbModel := range dbModels {
		graphModels = append(graphModels, converters.ProductCategoryToGraphQL(dbModel))
	}

	return graphModels, nil
}

// LegalEntityTypes is the resolver for the legalEntityTypes field.
func (r *queryResolver) LegalEntityTypes(ctx context.Context) ([]*model.LegalEntityType, error) {
	types, err := r.LegalEntityDao.GetAllTypes(ctx)

	if err != nil {
		return nil, err
	}

	var graphModels []*model.LegalEntityType

	for _, t := range types {
		graphModels = append(graphModels, converters.LegalEntityTypeToGraphQL(*t))
	}

	return graphModels, nil
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
func (r *queryResolver) ProductItems(ctx context.Context, pagination *model.PaginationInput, productSearchInput *model.ProductSearchInput) (*model.ProductItemPaginate, error) {
	currency, err := r.CurrencyDao.FindDefault(ctx)
	dao := r.ProductItemDao
	pWithoutPagination, err := dao.FindProductItems(ctx, productSearchInput, currency)

	if err != nil {
		return nil, err
	}

	if pagination != nil {
		dao = r.ProductItemDao.Paginate(pagination.Limit, pagination.Offset)
	}

	if err != nil {
		return nil, err
	}

	products, err := dao.FindProductItems(ctx, productSearchInput, currency)

	if err != nil {
		return nil, err
	}

	countInt := len(pWithoutPagination)

	var graphModels []*model.ProductItem

	for _, product := range products {
		graphModels = append(graphModels, converters.ProductItemToGraphQL(product))
	}

	return &model.ProductItemPaginate{
		ProductItems: graphModels,
		Pagination: &model.Pagination{
			Offset: &pagination.Offset,
			Limit:  &pagination.Limit,
			Totals: &countInt,
		},
	}, nil
}

// ProductItem is the resolver for the productItem field.
func (r *queryResolver) ProductItem(ctx context.Context, id int64) (*model.ProductItem, error) {
	dbModel, err := r.ProductItemDao.FindProductItemById(ctx, id)

	if err != nil {
		return nil, err
	}

	return converters.ProductItemToGraphQL(dbModel), nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, pagination *model.PaginationInput, productSearchInput *model.ProductSearchInput) (*model.ProductPaginate, error) {
	dao := r.ProductDao

	if pagination != nil {
		dao = r.ProductDao.Paginate(pagination.Limit, pagination.Offset)
	}

	currency, err := r.CurrencyDao.FindDefault(ctx)

	if err != nil {
		return nil, err
	}

	products, err := dao.Search(ctx, productSearchInput, currency)

	if err != nil {
		return nil, err
	}

	countInt := len(products)

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
