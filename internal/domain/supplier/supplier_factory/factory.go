package supplier_factory

import (
	"context"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/models"
)

type Importer interface {
	ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error)
	NeedsImportFromFile() bool
}

type Factory struct {
	Importer
	S *models.Supplier
}
