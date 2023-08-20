package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"errors"
	"fmt"
	"pillowww/titw/graph"
	"pillowww/titw/graph/aggregators"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/graphErrors"
	"pillowww/titw/graph/model"
	"pillowww/titw/graph/policies"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
	"time"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int64) (*model.User, error) {
	user, err := r.UserDao.FindOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	policy := policies.NewUserPolicy(user, r.UserDao)

	if !policy.CanRead(ctx) {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	return converters.UserToGraphQL(user), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, pagination *model.PaginationInput, filter *model.UserFilterInput) (*model.UserPaginator, error) {
	userDao := r.UserDao

	if pagination != nil {
		userDao = userDao.Paginate(pagination.Limit, pagination.Offset)
	}

	var users models.UserSlice
	var allUsers models.UserSlice
	var err error

	if filter != nil {
		users, err = userDao.FindAll(ctx, filter.Email, filter.Name, filter.Confirmed)
		allUsers, _ = r.UserDao.FindAll(ctx, filter.Email, filter.Name, filter.Confirmed)
	} else {
		users, err = userDao.FindAll(ctx, nil, nil, nil)
		allUsers, _ = r.UserDao.FindAll(ctx, nil, nil, nil)
	}

	if err != nil {
		return nil, err
	}

	return &model.UserPaginator{
		Data: aggregators.UsersToGraphql(users),
		Pagination: converters.PaginationToGraphql(
			pagination,
			len(allUsers),
		),
	}, nil
}

// UserAddress is the resolver for the userAddress field.
func (r *queryResolver) UserAddress(ctx context.Context, userID int64) ([]*model.UserAddress, error) {
	u, err := r.UserDao.FindOneById(ctx, userID)

	if err != nil {
		return nil, graphErrors.NewGraphError(ctx, errors.New("User not found in context"), "4004")
	}

	policy := policies.NewUserPolicy(u, r.UserDao)

	if !policy.CanRead(ctx) {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	return aggregators.NewUserAggregator(r.UserAddressDao).GetAllAddressesByUser(ctx, u.ID)
}

// UserBilling is the resolver for the userBilling field.
func (r *queryResolver) UserBilling(ctx context.Context, userID int64) (*model.UserBilling, error) {
	user, err := r.UserDao.FindOneById(ctx, userID)

	if err != nil {
		return nil, graphErrors.NewUserNotFoundError(ctx, err)
	}

	policy := policies.NewUserPolicy(user, r.UserDao)

	if !policy.CanRead(ctx) {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	billing, err := r.UserDao.GetUserBilling(ctx, user)

	if err != nil {
		return nil, graphErrors.NewGraphError(ctx, err, "Billing not found")
	}

	return converters.UserBillingToGraphQL(billing), err
}

// TaxRates is the resolver for the taxRates field.
func (r *queryResolver) TaxRates(ctx context.Context) ([]*model.Tax, error) {
	panic(fmt.Errorf("not implemented: TaxRates - taxRates"))
}

// Brands is the resolver for the brands field.
func (r *queryResolver) Brands(ctx context.Context) ([]*model.Brand, error) {
	brandModels, err := r.BrandDao.
		Paginate(0, 10).
		FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var graphModels []*model.Brand

	for _, m := range brandModels {
		graphModels = append(graphModels, converters.BrandToGraphQL(m))
	}

	return graphModels, nil
}

// SearchBrands is the resolver for the searchBrands field.
func (r *queryResolver) SearchBrands(ctx context.Context, name string) ([]*model.Brand, error) {
	var graphModels []*model.Brand

	if len(name) < 2 {
		all, err := r.BrandDao.FindAll(ctx)
		if err != nil {
			return nil, err
		}

		for _, model := range all {
			graphModels = append(graphModels, converters.BrandToGraphQL(model))
		}

		return graphModels, nil
	}

	models, err := r.BrandDao.Paginate(10, 0).FindByName(ctx, name)

	if err != nil {
		return nil, err
	}

	for _, model := range models {
		graphModels = append(graphModels, converters.BrandToGraphQL(model))
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

// UserOrders is the resolver for the userOrders field.
func (r *queryResolver) UserOrders(ctx context.Context, userID int64, pagination *model.PaginationInput, filter *model.OrderFilterInput, ordering []*model.OrderingInput) (*model.OrdersPaginator, error) {
	user, err := r.UserDao.FindOneById(ctx, userID)

	if err != nil {
		return nil, err
	}

	billing, err := r.UserDao.GetUserBilling(ctx, user)

	if err != nil {
		return nil, err
	}

	orderDao := r.OrderDao.
		Load(models.OrderRels.Currency)

	if pagination != nil {
		orderDao = orderDao.Paginate(pagination.Limit, pagination.Offset)
	}

	if ordering != nil {
		orderDao = orderDao.Order(ordering)
	}

	var orders models.OrderSlice
	var ordersWithoutPagination models.OrderSlice

	if filter != nil {
		ordersWithoutPagination, _ = r.OrderDao.FindAllByBillingId(ctx, billing.ID, filter.DateFrom, filter.DateTo, filter.Number)
		orders, err = orderDao.FindAllByBillingId(ctx, billing.ID, filter.DateFrom, filter.DateTo, filter.Number)
	} else {
		ordersWithoutPagination, _ = r.OrderDao.FindAllByBillingId(ctx, billing.ID, nil, nil, nil)
		orders, err = orderDao.FindAllByBillingId(ctx, billing.ID, nil, nil, nil)
	}

	totalCount := len(ordersWithoutPagination)

	if err != nil {
		return nil, err
	}

	graphOrders, err := aggregators.AggregateOrderModels(orders)

	if err != nil {
		return nil, err
	}

	return &model.OrdersPaginator{
		Data:       graphOrders,
		Pagination: converters.PaginationToGraphql(pagination, totalCount),
	}, nil
}

// PaymentMethods is the resolver for the paymentMethods field.
func (r *queryResolver) PaymentMethods(ctx context.Context) ([]*model.PaymentMethod, error) {
	methods, err := r.PaymentDao.FindAllPaymentMethods(ctx)

	if err != nil {
		return nil, err
	}

	var graphModels []*model.PaymentMethod

	for _, method := range methods {
		graphModels = append(graphModels, converters.PaymentMethodToGraphQL(method))
	}

	return graphModels, err
}

// Stats is the resolver for the stats field.
func (r *queryResolver) Stats(ctx context.Context) (*model.StatResponse, error) {
	totalUsers, err := r.UserDao.TotalUsers(ctx)

	if err != nil {
		return nil, err
	}

	defCur, err := r.CurrencyDao.FindDefault(ctx)

	if err != nil {
		return nil, err
	}

	to, err := r.OrderDao.TotalOrders(ctx, time.Now().Add(-time.Hour*24*30), time.Now())

	if err != nil {
		return nil, err
	}

	tof, err := currency.ToFloat(to.TotalPrice, defCur.IsoCode)

	if err != nil {
		return nil, err
	}

	brand, err := r.BrandDao.BestBrand(ctx, time.Now().Add(-time.Hour*24*30), time.Now())

	if err != nil {
		return nil, err
	}

	return &model.StatResponse{
		TotalUsers:  int(totalUsers),
		TotalOrders: tof,
		BestBrand:   converters.BrandToGraphQL(brand),
	}, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
