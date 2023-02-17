package supplier_factory

import "context"

type TyreWorld Factory

func (t TyreWorld) ImportProductsFromFile(ctx context.Context, filePath string) error {

	return nil
}

func (t TyreWorld) NeedsImportFromFile() bool {
	return true
}
