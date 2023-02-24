package supplier_factory

import (
	"context"
	"pillowww/titw/internal/domain/product/pdtos"
)

type PayGo Factory

func (g PayGo) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
	return nil, nil
}

func (g PayGo) NeedsImportFromFile() bool {
	return true
}
