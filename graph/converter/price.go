package converter

import (
	"context"
	currency2 "github.com/bojanz/currency"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/models"
	"strconv"
)

type PriceConverter struct {
	ProductDao  *product.Dao
	CurrencyDao *currency.Dao
}

func (p PriceConverter) ProductItemPrice(ctx context.Context, price *models.ProductItemPrice) (*model.Price, error) {
	cur, err := p.ProductDao.Currency(ctx, price)

	if err != nil {
		return nil, err
	}

	amount, err := currency2.NewAmountFromInt64(int64(price.Price), cur.IsoCode)
	if err != nil {
		return nil, err
	}

	floatValue, err := strconv.ParseFloat(amount.Number(), 64)

	if err != nil {
		return nil, err
	}

	return &model.Price{
		Value:        floatValue,
		CurrencyCode: cur.IsoCode,
	}, nil
}
