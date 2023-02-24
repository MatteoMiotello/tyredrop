package pdtos

import "pillowww/titw/pkg/constants"

type ProductDtoSlice []ProductDto

type ProductDto interface {
	GetProductCode() string
	Validate() bool
	GetProductCategoryCode() constants.ProductCategoryType
	GetSpecifications() map[constants.ProductSpecification]string
	GetBrandName() string
	GetSupplierProductPrice() string
	GetSupplierProductQuantity() int
}
