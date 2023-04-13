package resolvers

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/internal/domain/supplier"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/internal/domain/vehicle"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductDao                   *product.Dao
	BrandDao                     *brand.Dao
	SupplierDao                  *supplier.Dao
	ProductCategoryDao           *product.CategoryDao
	ProductSpecificationDao      *product.SpecificationDao
	ProductSpecificationValueDao *product.SpecificationValueDao
	ProductItemDao               *product.ItemDao
	CurrencyDao                  *currency.Dao
	VehicleDao                   *vehicle.Dao
	UserDao                      *user.Dao
}

func NewResolver(exec boil.ContextExecutor) *Resolver {
	return &Resolver{
		ProductDao:                   product.NewDao(exec),
		ProductCategoryDao:           product.NewCategoryDao(exec),
		ProductSpecificationValueDao: product.NewSpecificationValueDao(exec),
		ProductSpecificationDao:      product.NewSpecificationDao(exec),
		BrandDao:                     brand.NewDao(exec),
		SupplierDao:                  supplier.NewDao(exec),
		ProductItemDao:               product.NewItemDao(exec),
		CurrencyDao:                  currency.NewDao(exec),
		VehicleDao:                   vehicle.NewDao(exec),
		UserDao:                      user.NewDao(exec),
	}
}