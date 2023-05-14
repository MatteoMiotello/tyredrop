package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func UserToGraphQL(user *models.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  &user.Username.String,
		Confirmed: user.Confirmed,
		Name:      &user.Name,
		Surname:   &user.Surname,
	}
}

func UserBillingToGraphQL(billing *models.UserBilling) *model.UserBilling {
	return &model.UserBilling{
		ID:                billing.ID,
		LegalEntityTypeID: billing.LegalEntityTypeID,
		Name:              billing.Name,
		Surname:           &billing.Surname.String,
		FiscalCode:        billing.FiscalCode,
		VatNumber:         billing.VatNumber,
		AddressLine1:      billing.AddressLine1,
		AddressLine2:      &billing.AddressLine2.String,
		City:              billing.City,
		Province:          billing.Province,
		Cap:               billing.Cap,
		Country:           billing.Country,
	}
}
