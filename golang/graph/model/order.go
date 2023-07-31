package model

import (
	"pillowww/titw/models"
	"time"
)

type Order struct {
	ID               int64       `json:"id"`
	CurrencyID       int64       `json:"currencyID"`
	Tax              *Tax        `json:"tax"`
	TaxID            int64       `json:"taxID"`
	UserBillingID    int64       `json:"userBillingID"`
	Status           OrderStatus `json:"status,omitempty"`
	PriceAmount      float64     `json:"priceAmount"`
	PriceAmountTotal float64     `json:"priceAmountTotal"`
	TaxesAmount      float64     `json:"taxesAmount"`
	AddressName      string      `json:"addressName"`
	AddressLine1     string      `json:"addressLine1"`
	AddressLine2     *string     `json:"addressLine2,omitempty"`
	City             string      `json:"city"`
	Province         string      `json:"province,omitempty"`
	PostalCode       string      `json:"cap"`
	Country          string      `json:"country"`
	CreatedAt        time.Time   `json:"createdAt"`
}

type OrderRow struct {
	ID                 int64      `json:"id"`
	OrderID            int64      `json:"orderID"`
	ProductItemPriceID int64      `json:"productItemPriceID"`
	Quantity           int        `json:"quantity"`
	Amount             float64    `json:"amount"`
	AdditionsAmount    float64    `json:"additionsAmount"`
	TrackingNumber     *string    `json:"trackingNumber,omitempty"`
	DeliveredAt        *time.Time `json:"deliveredAt,omitempty"`
}

func (o OrderStatus) IsValidForOrder(order *models.Order) bool {
	switch o {
	case OrderStatusNew:
		return false
	case OrderStatusConfirmed, OrderStatusRejected, OrderStatusCanceled:
		return order.Status == OrderStatusNew.String()
	case OrderStatusDelivered:
		return order.Status == OrderStatusConfirmed.String()
	case OrderStatusReturned:
		return order.Status == OrderStatusConfirmed.String() || order.Status == OrderStatusNew.String()
	}
	return false
}
