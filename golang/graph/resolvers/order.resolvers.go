package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"pillowww/titw/graph"
	"pillowww/titw/graph/aggregators"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/graphErrors"
	"pillowww/titw/graph/model"
	"pillowww/titw/graph/policies"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Currency is the resolver for the currency field.
func (r *orderResolver) Currency(ctx context.Context, obj *model.Order) (*model.Currency, error) {
	langDef := auth2.CurrentLanguage(ctx)

	currencyModel, err := r.CurrencyDao.
		Load(
			models.CurrencyRels.CurrencyLanguages,
			models.CurrencyLanguageWhere.LanguageID.EQ(langDef.L.ID),
		).
		FindById(ctx, obj.CurrencyID)

	if err != nil {
		return nil, err
	}

	return converters.CurrencyToGraphQL(currencyModel), nil
}

// UserBilling is the resolver for the userBilling field.
func (r *orderResolver) UserBilling(ctx context.Context, obj *model.Order) (*model.UserBilling, error) {
	billingModel, err := r.UserDao.FindUserBillingById(ctx, obj.UserBillingID)

	if err != nil {
		return nil, err
	}

	return converters.UserBillingToGraphQL(billingModel), nil
}

// Payment is the resolver for the payment field.
func (r *orderResolver) Payment(ctx context.Context, obj *model.Order) (*model.Payment, error) {
	if obj.PaymentID == nil {
		return nil, nil
	}

	p, err := r.PaymentDao.
		Load(
			models.PaymentRels.Currency,
		).
		FindOneById(ctx, *obj.PaymentID)

	if err != nil {
		return nil, err
	}

	return converters.PaymentToGraphQL(p)
}

// OrderRows is the resolver for the orderRows field.
func (r *orderResolver) OrderRows(ctx context.Context, obj *model.Order) ([]*model.OrderRow, error) {
	orderModel, err := r.OrderDao.
		Load(models.OrderRels.OrderRows).
		Load(models.OrderRels.Currency).
		FindOneById(ctx, obj.ID)

	if err != nil {
		return nil, err
	}

	return aggregators.RowsFromOrderModel(orderModel), err
}

// Order is the resolver for the order field.
func (r *orderRowResolver) Order(ctx context.Context, obj *model.OrderRow) (*model.Order, error) {
	orderModel, err := r.OrderDao.
		Load(models.OrderRels.Currency).
		FindOneById(ctx, obj.OrderID)

	if err != nil {
		return nil, err
	}

	policy := policies.NewOrderPolicy(orderModel, r.OrderDao)

	if !policy.CanRead(ctx) {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	return converters.OrderToGraphQL(orderModel)
}

// ProductItemPrice is the resolver for the productItemPrice field.
func (r *orderRowResolver) ProductItemPrice(ctx context.Context, obj *model.OrderRow) (*model.ProductItemPrice, error) {
	lang := auth2.CurrentLanguage(ctx)

	dbModel, err := r.ProductItemPriceDao.
		WithDeletes().
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

// Order returns graph.OrderResolver implementation.
func (r *Resolver) Order() graph.OrderResolver { return &orderResolver{r} }

// OrderRow returns graph.OrderRowResolver implementation.
func (r *Resolver) OrderRow() graph.OrderRowResolver { return &orderRowResolver{r} }

type orderResolver struct{ *Resolver }
type orderRowResolver struct{ *Resolver }
