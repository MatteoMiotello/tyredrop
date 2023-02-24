package supplier_factory

import (
	"context"
	"pillowww/titw/internal/domain/product/pdtos"
)

type Seng Factory

func (s Seng) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {

	return nil, nil
}

func (s Seng) NeedsImportFromFile() bool {
	return true
}
