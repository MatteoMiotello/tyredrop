package model

type User struct {
	ID       int64   `json:"id"`
	Email    string  `json:"email"`
	Username *string `json:"username"`
}

type UserBilling struct {
	ID                int64   `json:"id"`
	UserID            int64   `json:"userID"`
	LegalEntityTypeID int64   `json:"legalEntityTypeID"`
	Name              string  `json:"name"`
	Surname           *string `json:"surname"`
	FiscalCode        string  `json:"fiscalCode"`
	VatNumber         string  `json:"vatNumber"`
	AddressLine1      string  `json:"addressLine1"`
	AddressLine2      *string `json:"addressLine2"`
	City              string  `json:"city"`
	Province          string  `json:"province"`
	Cap               string  `json:"cap"`
	Country           string  `json:"country"`
}
