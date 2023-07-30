package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
)

func PriceAdditionTypeToGraphQL(additionType *models.PriceAdditionType) (*model.PriceAdditionType, error) {
	curr := additionType.R.Currency

	if curr == nil {
		panic("currency not loaded on addition type")
	}

	floatVal, err := currency.ToFloat(additionType.AdditionValue, curr.IsoCode)

	if err != nil {
		return nil, err
	}

	return &model.PriceAdditionType{
		ID:            additionType.ID,
		CurrencyID:    additionType.CurrencyID,
		AdditionType:  additionType.AdditionType,
		AdditionName:  additionType.AdditionName,
		AdditionCode:  additionType.AdditionCode,
		AdditionValue: floatVal,
	}, nil
}

func ProductItemPriceAdditionToGraphQL(addition *models.ProductItemPriceAddition) (*model.ProductItemPriceAddition, error) {
	t := addition.R.ProductItemPrice

	if t == nil {
		panic("price not loaded in price addition")
	}

	curr := t.R.Currency

	if curr == nil {
		panic("currency not loaded in price addition")
	}

	floatVal, err := currency.ToFloat(addition.AdditionValue, curr.IsoCode)

	if err != nil {
		return nil, err
	}

	return &model.ProductItemPriceAddition{
		ID:                  addition.ID,
		ProductItemPriceID:  addition.ProductItemPriceID,
		PriceAdditionTypeID: addition.PriceAdditionTypeID,
		AdditionValue:       floatVal,
	}, nil
}
