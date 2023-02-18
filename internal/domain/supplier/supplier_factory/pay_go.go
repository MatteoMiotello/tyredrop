package supplier_factory

import "context"

type PayGo Factory

func (g PayGo) ReadProductsFromFile(ctx context.Context, filePath string) ([]*ProductRecord, error) {
	return nil, nil
}

func (g PayGo) NeedsImportFromFile() bool {
	return true
}
