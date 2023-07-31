package supplier

import (
	"pillowww/titw/internal/domain/supplier/supplier_factory"
	"pillowww/titw/models"
)

func GetFactory(supplier *models.Supplier) supplier_factory.Importer {
	switch supplier.Code {
	case PAY_GO:
		return &supplier_factory.PayGo{S: supplier}
	case GUN:
		return &supplier_factory.Gun{S: supplier}
	case SENG:
		return &supplier_factory.Seng{S: supplier}
	case TYRE_WORLD:
		return &supplier_factory.TyreWorld{S: supplier}
	}
	return nil
}
