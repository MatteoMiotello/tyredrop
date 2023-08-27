package converters

import (
	"github.com/friendsofgo/errors"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
)

func ProductItemPriceToGraphQL(price *models.ProductItemPrice) (*model.ProductItemPrice, error) {
	cur := price.R.Currency

	if cur == nil {
		return nil, errors.New("Currency not loaded for item price")
	}

	floatValue, err := currency.ToFloat(price.Price, cur.IsoCode)

	if err != nil {
		return nil, err
	}

	curGraph := CurrencyToGraphQL(cur)

	return &model.ProductItemPrice{
		ID:            price.ID,
		Value:         floatValue,
		Currency:      curGraph,
		ProductItemID: price.ProductItemID,
	}, nil
}

func PriceMarkupToGraphQL(markup *models.ProductPriceMarkup) *model.ProductPriceMarkup {
	return &model.ProductPriceMarkup{
		ID:                markup.ID,
		MarkupPercentage:  markup.MarkupPercentage,
		ProductID:         markup.ProductID.Ptr(),
		ProductCategoryID: markup.ProductCategoryID.Ptr(),
		BrandID:           markup.BrandID.Ptr(),
	}
}
