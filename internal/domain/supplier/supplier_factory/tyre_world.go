package supplier_factory

import (
	"context"
	"pillowww/titw/internal/domain/product/pdtos"
)

type TyreWorld Factory

func (t TyreWorld) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {

	return nil, nil
}

func (t TyreWorld) NeedsImportFromFile() bool {
	return true
}
