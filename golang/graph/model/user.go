package model

import "time"

type User struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	Username   *string   `json:"username"`
	UserCode   *string   `json:"userCode"`
	Confirmed  bool      `json:"confirmed"`
	Rejected   bool      `json:"rejected"`
	Name       *string   `json:"name"`
	Surname    *string   `json:"surname"`
	UserRoleID int64     `json:"userRoleId"`
	AvatarPath *string   `json:"avatarPath"`
	CreatedAt  time.Time `json:"createdAt"`
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
	SdiCode           *string `json:"sdiCode"`
	SdiPec            *string `json:"sdiEmail"`
}

type UserAddress struct {
	ID           int64   `json:"ID"`
	UserID       int64   `json:"userID"`
	AddressName  string  `json:"addressName"`
	AddressLine1 string  `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	City         string  `json:"city"`
	Province     string  `json:"province"`
	PostalCode   string  `json:"postalCode"`
	Country      string  `json:"country"`
	IsDefault    bool    `json:"isDefault"`
}
