package model

type Product struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

type ProductCategory struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProductItem struct {
	ID               int64 `json:"id"`
	SupplierQuantity int   `json:"supplier_quantity"`
}

type ProductSpecification struct {
	ID             int64         `json:"id"`
	Code           string        `json:"code"`
	Name           string        `json:"name"`
	Type           string        `json:"type"`
	Values         []UnionValues `json:"values"`
	PossibleValues []*string     `json:"possible_values"`
}
