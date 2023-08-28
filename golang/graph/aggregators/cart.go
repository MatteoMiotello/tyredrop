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
	"pillowww/titw/internal/domain/product"
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
			),
			models.ProductItemPriceWhere.CurrencyID.EQ(defaultCur.ID),
			models.ProductItemPriceWhere.DeletedAt.IsNull(),
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
	var addMap map[string]int = make(map[string]int)

	for _, c := range cartModels {
		price := c.R.ProductItemPrice

		amount, err := currency.ToFloat(price.Price, defaultCur.IsoCode)

		if err != nil {
			return nil, err
		}

		amount = *amountTotal + (amount * float64(c.Quantity))
		amountTotal = &amount

		pAdds, err := product.NewItemPriceDao(cartDao.Db).
			Load(
				models.ProductItemPriceAdditionRels.PriceAdditionType,
			).
			FindPriceAdditionsByProductItemPriceID(ctx, price.ID)

		if err != nil {
			return nil, err
		}

		for _, add := range pAdds {
			if _, ok := addMap[add.R.PriceAdditionType.AdditionName]; ok == false {
				addMap[add.R.PriceAdditionType.AdditionName] = add.AdditionValue * c.Quantity
				continue
			}

			addMap[add.R.PriceAdditionType.AdditionName] = addMap[add.R.PriceAdditionType.AdditionName] + (add.AdditionValue * c.Quantity)
		}

		graphModels = append(graphModels, converters.CartToGraphQL(c))
	}

	tax, err := order.NewDao(cartDao.Db).FindDefaultTax(ctx)

	if err != nil {
		return nil, err
	}

	var additionValues []*model.AdditionValue
	var totalAdditions float64

	for name, value := range addMap {
		floatVal, err := currency.ToFloat(value, defaultCur.IsoCode)

		if err != nil {
			return nil, err
		}

		totalAdditions = totalAdditions + floatVal
		additionValues = append(additionValues, &model.AdditionValue{
			AdditionName: name,
			Value:        floatVal,
		})
	}

	taxValue := float64(tax.MarkupPercentage) / 100 * (*amountTotal + totalAdditions)

	return &model.CartResponse{
		Items: graphModels,
		TotalPrice: &model.TotalPrice{
			Value:           *amountTotal,
			TotalValue:      *amountTotal + taxValue + totalAdditions,
			TaxesValue:      taxValue,
			AdditionsValues: additionValues,
			Currency:        converters.CurrencyToGraphQL(defaultCur),
		},
	}, nil
}
