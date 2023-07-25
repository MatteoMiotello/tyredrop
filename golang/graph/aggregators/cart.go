package aggregators

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/cart"
	"pillowww/titw/models"
)

func GetAllCartsByUserId(ctx context.Context, cartDao *cart.Dao, userId int64) (*model.CartResponse, error) {
	defaultCur, err := auth2.CurrentCurrency(ctx)

	if err != nil {
		return nil, err
	}

	cartModels, _ := cartDao.
		Load(
			qm.Rels(
				models.CartRels.ProductItemPrice,
				models.ProductItemPriceRels.Currency,
			),
		).
		FindAllByUserId(ctx, userId)

	amountTotal := new(float64)

	if cartModels == nil {
		return &model.CartResponse{
			Items:      []*model.Cart{},
			TotalPrice: nil,
		}, nil
	}

	var graphModels []*model.Cart
	for _, c := range cartModels {
		price := c.R.ProductItemPrice

		amount, err := currency.ToFloat(price.Price, price.R.Currency.IsoCode)

		if err != nil {
			return nil, err
		}

		amount = *amountTotal + (amount * float64(c.Quantity))
		amountTotal = &amount
		graphModels = append(graphModels, converters.CartToGraphQL(c))
	}

	return &model.CartResponse{
		Items: graphModels,
		TotalPrice: &model.TotalPrice{
			Value:    *amountTotal,
			Currency: converters.CurrencyToGraphQL(defaultCur),
		},
	}, nil
}
