package model

type Product struct {
	ID                int64  `json:"id"`
	Code              string `json:"code"`
	ProductCategoryID int64  `json:"product_category_id"`
	BrandID           int64  `json:"brand_id"`
	VehicleTypeID     int64  `json:"vehicle_type_id"`
	Name              string `json:"name"`
	EprelProductCode  string `json:"eprel_product_code"`
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
	Mandatory         bool   `json:"mandatory"`
	Searchable        bool   `json:"searchable"`
	ProductCategoryID int64  `json:"product_category_id"`
}

type ProductSpecificationValue struct {
	ID              int64  `json:"id"`
	Value           string `json:"value"`
	SpecificationID int64  `json:"specificationId"`
}
