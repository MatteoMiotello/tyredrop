package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/graphErrors"
	"pillowww/titw/graph/model"
	"pillowww/titw/graph/policies"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/order"
	"pillowww/titw/internal/domain/payment"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/models"
)

// OrderSupport is the resolver for the orderSupport field.
func (r *mutationResolver) OrderSupport(ctx context.Context, orderID int64, message string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: OrderSupport - orderSupport"))
}

// ConfirmOrder is the resolver for the confirmOrder field.
func (r *mutationResolver) ConfirmOrder(ctx context.Context, orderID int64) (*model.Order, error) {
	o, err := r.OrderDao.
		Load(
			models.OrderRels.Currency,
		).
		FindOneById(ctx, orderID)

	if err != nil {
		return nil, err
	}

	policy := policies.NewOrderPolicy(o, r.OrderDao)

	if !policy.CanConfirm(ctx) {
		return nil, graphErrors.NewGraphError(ctx, errors.New("Order can't be confirmed"), "UNABLE_TO_PAY_ORDER")
	}

	s := order.NewService(r.OrderDao, r.CurrencyDao, r.ProductItemDao, r.ProductItemPriceDao)

	err = s.ConfirmOrder(ctx, o)

	if err != nil {
		return nil, err
	}

	return converters.OrderToGraphQL(o)
}

// PayOrder is the resolver for the payOrder field.
func (r *mutationResolver) PayOrder(ctx context.Context, orderID int64, paymentMethodCode string) (*model.Order, error) {
	o, err := r.OrderDao.
		Load(
			models.OrderRels.Currency,
		).
		FindOneById(ctx, orderID)

	if err != nil {
		return nil, err
	}

	policy := policies.NewOrderPolicy(o, r.OrderDao)

	if !policy.CanPay(ctx) {
		return nil, graphErrors.NewGraphError(ctx, errors.New("Order can't be payed"), "UNABLE_TO_PAY_ORDER")
	}

	method, err := r.PaymentDao.FindPaymentMethodByCode(ctx, paymentMethodCode)

	err = db.WithTx(ctx, func(tx *sql.Tx) error {
		oDao := order.NewDao(tx)
		currencyDao := currency.NewDao(tx)
		itemDao := product.NewItemDao(tx)
		itemPriceDao := product.NewItemPriceDao(tx)

		oService := order.NewService(
			oDao,
			currencyDao,
			itemDao,
			itemPriceDao,
		)

		service := payment.NewService(
			payment.NewDao(tx),
			oDao,
		)

		p, err := service.CreatePayment(ctx, o, method)

		if err != nil {
			return err
		}

		err = oService.PayOrder(ctx, o, p)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return converters.OrderToGraphQL(o)
}

// UpdateOrderStatus is the resolver for the updateOrderStatus field.
func (r *mutationResolver) UpdateOrderStatus(ctx context.Context, orderID int64, newStatus model.OrderStatus) (*model.Order, error) {
	oService := order.NewService(
		r.OrderDao,
		r.CurrencyDao,
		r.ProductItemDao,
		r.ProductItemPriceDao,
	)

	o, err := r.OrderDao.
		Load(
			models.OrderRels.Currency,
		).
		FindOneById(ctx, orderID)

	if err != nil {
		return nil, err
	}

	err = oService.UpdateOrderStatus(ctx, o, newStatus)

	if err != nil {
		return nil, err
	}

	return converters.OrderToGraphQL(o)
}
