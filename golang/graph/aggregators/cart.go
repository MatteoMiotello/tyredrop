package aggregators

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/cart"
	"pillowww/titw/internal/domain/order"
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
				models.ProductItemPriceRels.ProductItemPriceAdditions,
				models.ProductItemPriceAdditionRels.PriceAdditionType,
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
	var additions map[string]int

	for _, c := range cartModels {
		price := c.R.ProductItemPrice

		amount, err := currency.ToFloat(price.Price, price.R.Currency.IsoCode)

		if err != nil {
			return nil, err
		}

		amount = *amountTotal + (amount * float64(c.Quantity))
		amountTotal = &amount

		for _, add := range price.R.ProductItemPriceAdditions {
			if _, ok := additions[add.R.PriceAdditionType.AdditionName]; ok == false {
				additions[add.R.PriceAdditionType.AdditionName] = add.AdditionValue
				continue
			}

			additions[add.R.PriceAdditionType.AdditionName] = additions[add.R.PriceAdditionType.AdditionName] + add.AdditionValue
		}

		graphModels = append(graphModels, converters.CartToGraphQL(c))
	}

	tax, err := order.NewDao(cartDao.Db).FindDefaultTax(ctx)

	if err != nil {
		return nil, err
	}

	taxValue := float64(tax.MarkupPercentage/100) * (*amountTotal)

	var additionValues []*model.AdditionValue

	for name, value := range additions {
		additionValues = append(additionValues, &model.AdditionValue{
			AdditionName: name,
			Value:        currency.ToFloat(value),
		})
	}

	return &model.CartResponse{
		Items: graphModels,
		TotalPrice: &model.TotalPrice{
			Value:      *amountTotal,
			TaxesValue: taxValue,
			Currency:   converters.CurrencyToGraphQL(defaultCur),
		},
	}, nil
}
