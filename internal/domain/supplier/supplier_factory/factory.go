package supplier_factory

import (
	"context"
	"pillowww/titw/models"
)

type Importer interface {
	ImportProductsFromFile(ctx context.Context, filePath string) (err error)
	NeedsImportFromFile() bool
}

type Factory struct {
	Importer
	S *models.Supplier
}
