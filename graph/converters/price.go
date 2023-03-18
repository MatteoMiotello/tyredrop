package converters

import (
	currency2 "github.com/bojanz/currency"
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
	"strconv"
)

func ProductItemPriceToGraphQL(price *models.ProductItemPrice) (*model.ProductPrice, error) {
	cur := price.R.Currency

	amount, err := currency2.NewAmountFromInt64(int64(price.Price), cur.IsoCode)
	if err != nil {
		return nil, err
	}

	floatValue, err := strconv.ParseFloat(amount.Number(), 64)

	if err != nil {
		return nil, err
	}

	return &model.ProductPrice{
		ID:    price.ID,
		Value: floatValue,
	}, nil
}
