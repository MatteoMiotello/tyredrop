package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func ProductSpecificationValueToGraphQL(value *models.ProductSpecificationValue) *model.ProductSpecificationValue {
	return &model.ProductSpecificationValue{
		ID:              value.ID,
		Value:           value.SpecificationValue,
		SpecificationID: value.ProductSpecificationID,
	}
}

func ProductSpecificationToGraphQL(spec *models.ProductSpecification) *model.ProductSpecification {
	lang := spec.R.ProductSpecificationLanguages[0]

	return &model.ProductSpecification{
		ID:                spec.ID,
		Code:              spec.SpecificationCode,
		Type:              spec.Type,
		Name:              lang.Name,
		ProductCategoryID: spec.ProductCategoryID,
	}
}

func ProductToGraphQL(product *models.Product) *model.Product {
	return &model.Product{
		ID:                product.ID,
		Code:              product.ProductCode.String,
		ProductCategoryID: product.ProductCategoryID,
		BrandID:           product.BrandID,
	}
}

func ProductItemToGraphQL(productItem *models.ProductItem) *model.ProductItem {
	return &model.ProductItem{
		ID:               productItem.ID,
		SupplierQuantity: productItem.SupplierQuantity,
		ProductID:        productItem.ProductID,
		SupplierID:       productItem.SupplierID,
	}
}

func ProductCategoryToGraphQL(productCategory *models.ProductCategory) *model.ProductCategory {
	categoryLangs := productCategory.R.ProductCategoryLanguages

	return &model.ProductCategory{
		ID:   productCategory.ID,
		Code: productCategory.CategoryCode,
		Name: categoryLangs[0].Name,
	}
}
