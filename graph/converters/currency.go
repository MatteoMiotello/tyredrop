package converters

import (
	"github.com/friendsofgo/errors"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/domain/language"
	"pillowww/titw/models"
)

func CurrencyToGraphQL(currency *models.Currency, language *language.Language) (*model.Currency, error) {
	langs := currency.R.CurrencyLanguages
	var selected *models.CurrencyLanguage
	for _, lang := range langs {
		if lang.LanguageID == language.L.ID {
			selected = lang
			break
		}
	}

	if selected == nil {
		return nil, errors.New("language not found for currency: " + currency.IsoCode)
	}

	return &model.Currency{
		ID:      currency.ID,
		IsoCode: currency.IsoCode,
		Symbol:  currency.Symbol,
		Tag:     currency.Tag,
		Name:    selected.Name,
	}, nil
}
