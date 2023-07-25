package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
)

func OrderToGraphQL(order *models.Order) *model.Order {
	return &model.Order{
		ID:            order.ID,
		CurrencyID:    order.CurrencyID,
		UserBillingID: order.UserBillingID,
		TaxID:         order.TaxID,
		AddressLine1:  order.AddressLine1,
		AddressLine2:  &order.AddressLine2.String,
		Country:       order.Country,
		PostalCode:    order.PostalCode,
		City:          order.City,
		Province:      order.Province,
		Status:        model.OrderStatus(order.Status),
		CreatedAt:     order.CreatedAt,
	}
}

func OrderRowToGraphQL(row *models.OrderRow, cur *models.Currency) (*model.OrderRow, error) {
	floatValue, err := currency.ToFloat(row.Amount, cur.IsoCode)

	if err != nil {
		return nil, err
	}

	return &model.OrderRow{
		ID:                 row.ID,
		OrderID:            row.OrderID,
		ProductItemPriceID: row.ProductItemPriceID,
		Quantity:           row.Quantity,
		Amount:             floatValue,
		TrackingNumber:     &row.TrackingNumber.String,
		DeliveredAt:        &row.DeliveredAt.Time,
	}, nil
}
