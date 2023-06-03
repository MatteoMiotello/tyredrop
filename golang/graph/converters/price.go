package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
)

func ProductItemPriceToGraphQL(price *models.ProductItemPrice) (*model.ProductPrice, error) {
	cur := price.R.Currency

	floatValue, err := currency.ToFloat(price.Price, cur.IsoCode)

	if err != nil {
		return nil, err
	}

	curGraph := CurrencyToGraphQL(cur)

	return &model.ProductPrice{
		ID:       price.ID,
		Value:    floatValue,
		Currency: curGraph,
	}, nil
}
