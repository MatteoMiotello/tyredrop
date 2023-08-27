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

type ProductPriceMarkup struct {
	ID                int64  `json:"id"`
	MarkupPercentage  int    `json:"markupPercentage"`
	ProductCategoryID *int64 `json:"productCategoryID,omitempty"`
	BrandID           *int64 `json:"brandID,omitempty"`
	ProductID         *int64 `json:"productID,omitempty"`
}
