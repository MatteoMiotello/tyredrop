package supplier_factory

import "context"

type Seng Factory

func (s Seng) ImportProductsFromFile(ctx context.Context, filePath string) error {

	return nil
}

func (s Seng) NeedsImportFromFile() bool {
	return true
}
