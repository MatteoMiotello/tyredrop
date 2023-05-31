package model

type Cart struct {
	ID            int64 `json:"id"`
	UserID        int64 `json:"userId"`
	ProductItemID int64 `json:"productItemId"`
	Quantity      int   `json:"quantity"`
}
