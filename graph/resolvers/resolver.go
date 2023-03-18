package resolvers

import (
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/internal/domain/supplier"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductDao  *product.Dao
	BrandDao    *brand.Dao
	SupplierDao *supplier.Dao
}
