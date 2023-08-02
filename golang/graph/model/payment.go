package model

type Payment struct {
	ID                  int64   `json:"id"`
	UserPaymentMethodID int64   `json:"userPaymentMethodID"`
	CurrencyID          int64   `json:"currencyID"`
	Amount              float64 `json:"amount"`
}

type UserPaymentMethod struct {
	ID              int64   `json:"id"`
	PaymentMethodID int64   `json:"paymentMethodId"`
	Name            string  `json:"name"`
	Value           *string `json:"value,omitempty"`
	TypePrimary     bool    `json:"typePrimary"`
}
