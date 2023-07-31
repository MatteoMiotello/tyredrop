package model

type Cart struct {
	ID                 int64 `json:"id"`
	UserID             int64 `json:"userId"`
	ProductItemPriceID int64 `json:"productItemPriceId"`
	Quantity           int   `json:"quantity"`
}
