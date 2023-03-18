package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func CurrencyToGraphQL(currency *models.Currency) (*model.Currency, error) {
	lang := currency.R.CurrencyLanguages[0]

	return &model.Currency{
		ID:      currency.ID,
		IsoCode: currency.IsoCode,
		Symbol:  currency.Symbol,
		Tag:     currency.Tag,
		Name:    lang.Name,
	}, nil
}
