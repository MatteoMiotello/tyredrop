package supplier_factory

import "context"

type PayGo Factory

func (g PayGo) ImportProductsFromFile(ctx context.Context, filePath string) error {
	return nil
}

func (g PayGo) NeedsImportFromFile() bool {
	return true
}
