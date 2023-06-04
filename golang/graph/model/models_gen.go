// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Brand struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	ImageLogo string `json:"image_logo"`
}

type CartResponse struct {
	Items      []*Cart     `json:"items"`
	TotalPrice *TotalPrice `json:"totalPrice,omitempty"`
}

type CreateAdminUserInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserBilling struct {
	LegalEntityTypeID int64   `json:"legalEntityTypeId"`
	Name              string  `json:"name"`
	Surname           *string `json:"surname,omitempty"`
	FiscalCode        *string `json:"fiscalCode,omitempty"`
	VatNumber         string  `json:"vatNumber"`
	AddressLine1      string  `json:"addressLine1"`
	AddressLine2      *string `json:"addressLine2,omitempty"`
	City              string  `json:"city"`
	Province          string  `json:"province"`
	Cap               string  `json:"cap"`
	Country           string  `json:"country"`
	Iban              string  `json:"iban"`
}

type Currency struct {
	ID      int64  `json:"id"`
	IsoCode string `json:"iso_code"`
	Symbol  string `json:"symbol"`
	Tag     string `json:"tag"`
	Name    string `json:"name"`
}

type LegalEntityType struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsPerson bool   `json:"isPerson"`
}

type Pagination struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
	Totals *int `json:"totals,omitempty"`
}

type PaginationInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ProductItemPaginate struct {
	Pagination   *Pagination    `json:"pagination,omitempty"`
	ProductItems []*ProductItem `json:"productItems,omitempty"`
}

type ProductPaginate struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Products   []*Product  `json:"products,omitempty"`
}

type ProductPrice struct {
	ID       int64     `json:"id"`
	Value    float64   `json:"value"`
	Currency *Currency `json:"currency"`
}

type ProductSearchInput struct {
	Brand          *string                      `json:"brand,omitempty"`
	Name           *string                      `json:"name,omitempty"`
	Code           *string                      `json:"code,omitempty"`
	Specifications []*ProductSpecificationInput `json:"specifications,omitempty"`
}

type ProductSpecificationInput struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

type Supplier struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Tax struct {
	ID               int64   `json:"id"`
	MarkupPercentage float64 `json:"markupPercentage"`
	Name             string  `json:"name"`
}

type TotalPrice struct {
	Value    float64   `json:"value"`
	Currency *Currency `json:"currency,omitempty"`
}

type UserAddressInput struct {
	AddressLine1 string  `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	City         string  `json:"city"`
	Province     string  `json:"province"`
	PostalCode   string  `json:"postalCode"`
	Country      string  `json:"country"`
	IsDefault    bool    `json:"IsDefault"`
}

type UserRole struct {
	ID       int64  `json:"id"`
	RoleCode string `json:"roleCode"`
	Name     string `json:"name"`
	IsAdmin  *bool  `json:"isAdmin,omitempty"`
}

type VehicleType struct {
	ID   int64  `json:"ID"`
	Code string `json:"code"`
	Name string `json:"name"`
}
