package supplier_factory

import "context"

type TyreWorld Factory

func (t TyreWorld) ReadProductsFromFile(ctx context.Context, filePath string) ([]*ProductRecord, error) {

	return nil, nil
}

func (t TyreWorld) NeedsImportFromFile() bool {
	return true
}
