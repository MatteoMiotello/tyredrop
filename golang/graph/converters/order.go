package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
)

func OrderToGraphQL(order *models.Order) (*model.Order, error) {
	curr := order.R.Currency

	if curr == nil {
		panic("Currency not loaded on order")
	}

	amountFloatValue, err := currency.ToFloat(order.PriceAmount, curr.IsoCode)

	if err != nil {
		return nil, err
	}

	amountTotalFloatValue, err := currency.ToFloat(order.PriceAmountTotal, curr.IsoCode)

	if err != nil {
		return nil, err
	}

	taxesFloatValue, err := currency.ToFloat(order.TaxesAmount, curr.IsoCode)

	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:               order.ID,
		CurrencyID:       order.CurrencyID,
		UserBillingID:    order.UserBillingID,
		TaxID:            order.TaxID,
		PriceAmount:      amountFloatValue,
		PriceAmountTotal: amountTotalFloatValue,
		TaxesAmount:      taxesFloatValue,
		PaymentID:        order.PaymentID.Ptr(),
		AddressName:      order.AddressName,
		AddressLine1:     order.AddressLine1,
		AddressLine2:     &order.AddressLine2.String,
		Country:          order.Country,
		PostalCode:       order.PostalCode,
		City:             order.City,
		Province:         order.Province,
		Status:           model.OrderStatus(order.Status),
		CreatedAt:        order.CreatedAt,
	}, nil
}

func OrderRowToGraphQL(row *models.OrderRow, cur *models.Currency) (*model.OrderRow, error) {
	floatValue, err := currency.ToFloat(row.Amount, cur.IsoCode)

	if err != nil {
		return nil, err
	}

	additionsFloatValue, err := currency.ToFloat(row.AdditionsAmount, cur.IsoCode)

	if err != nil {
		return nil, err
	}

	return &model.OrderRow{
		ID:                 row.ID,
		OrderID:            row.OrderID,
		ProductItemPriceID: row.ProductItemPriceID,
		Quantity:           row.Quantity,
		Amount:             floatValue,
		AdditionsAmount:    additionsFloatValue,
		TrackingNumber:     &row.TrackingNumber.String,
		DeliveredAt:        &row.DeliveredAt.Time,
	}, nil
}
