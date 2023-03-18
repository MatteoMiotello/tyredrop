package model

type Product struct {
	ID                int64  `json:"id"`
	Code              string `json:"code"`
	ProductCategoryID int64  `json:"product_category_id"`
	BrandID           int64  `json:"brand_id"`
}

type ProductCategory struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProductItem struct {
	ID               int64   `json:"id"`
	SupplierQuantity int     `json:"supplier_quantity"`
	ProductID        int64   `json:"product_id"`
	SupplierID       int64   `json:"supplier_id"`
	SupplierPrice    float64 `json:"supplier_price"`
}

type ProductSpecification struct {
	ID                int64  `json:"id"`
	Code              string `json:"code"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	ProductCategoryID int64  `json:"product_category_id"`
}

type ProductSpecificationValue struct {
	ID              int64  `json:"id"`
	Value           string `json:"value"`
	SpecificationID int64  `json:"specificationId"`
}
