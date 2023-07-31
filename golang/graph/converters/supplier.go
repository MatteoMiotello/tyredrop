package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func SupplierToGraphQL(sup *models.Supplier) *model.Supplier {
	return &model.Supplier{
		ID:   sup.ID,
		Name: sup.Name,
		Code: sup.Code,
	}
}
