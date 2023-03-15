package converter

import (
	"context"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/domain/supplier"
	"pillowww/titw/models"
)

type SupplierConverter struct {
	SupplierDao *supplier.Dao
}

func (c SupplierConverter) Supplier(ctx context.Context, sup *models.Supplier) *model.Supplier {
	return &model.Supplier{
		ID:   sup.ID,
		Name: sup.Name,
		Code: sup.Code,
	}
}
