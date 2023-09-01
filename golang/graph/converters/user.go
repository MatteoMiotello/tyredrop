package converters

import (
	"github.com/volatiletech/null/v8"
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func UserToGraphQL(user *models.User) *model.User {
	return &model.User{
		ID:         user.ID,
		Email:      user.Email,
		UserCode:   user.UserCode.Ptr(),
		Username:   &user.Username.String,
		Confirmed:  user.Confirmed,
		Rejected:   user.Rejected,
		Name:       &user.Name,
		Surname:    user.Surname.Ptr(),
		UserRoleID: user.UserRoleID,
		AvatarPath: user.AvatarPath.Ptr(),
		CreatedAt:  user.CreatedAt,
	}
}

func UserRoleToGraphQL(role *models.UserRole) *model.UserRole {
	if len(role.R.UserRoleLanguages) == 0 {
		panic("role languages not loaded")
	}

	lang := role.R.UserRoleLanguages[0]

	return &model.UserRole{
		ID:       role.ID,
		RoleCode: role.RoleCode,
		Name:     lang.Name,
		IsAdmin:  role.Admin,
	}
}

func UserBillingToGraphQL(billing *models.UserBilling) *model.UserBilling {
	return &model.UserBilling{
		ID:                billing.ID,
		LegalEntityTypeID: billing.LegalEntityTypeID,
		UserID:            billing.UserID,
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
		SdiCode:           billing.SdiCode.Ptr(),
		SdiPec:            billing.SdiPec.Ptr(),
	}
}

func UserAddressToGraphQL(address *models.UserAddress) *model.UserAddress {
	return &model.UserAddress{
		ID:           address.ID,
		UserID:       address.UserID,
		AddressName:  address.AddressName,
		AddressLine1: address.AddressLine1,
		AddressLine2: &address.AddressLine2.String,
		City:         address.City,
		Province:     address.Province,
		PostalCode:   address.PostalCode,
		Country:      address.Country,
		IsDefault:    address.IsDefault,
	}
}

func GraphQLToUserAddress(address model.UserAddressInput, dbModel *models.UserAddress) {
	dbModel.AddressLine1 = address.AddressLine1
	dbModel.AddressLine2 = null.StringFrom(*address.AddressLine2)
	dbModel.City = address.City
	dbModel.Province = address.Province
	dbModel.PostalCode = address.PostalCode
	dbModel.Country = address.Country
	dbModel.IsDefault = address.IsDefault
	dbModel.AddressName = address.AddressName
}
