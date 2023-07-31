package model

type ProductItemPriceAddition struct {
	ID                  int64   `json:"ID"`
	ProductItemPriceID  int64   `json:"productItemPriceId"`
	PriceAdditionTypeID int64   `json:"priceAdditionTypeId"`
	AdditionValue       float64 `json:"additionValue"`
}

type PriceAdditionType struct {
	ID            int64   `json:"ID"`
	CurrencyID    int64   `json:"currencyId"`
	AdditionType  string  `json:"additionType"`
	AdditionName  string  `json:"additionName"`
	AdditionCode  string  `json:"additionCode"`
	AdditionValue float64 `json:"additionValue"`
}
