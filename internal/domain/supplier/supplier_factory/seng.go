package supplier_factory

import "context"

type Seng Factory

func (s Seng) ReadProductsFromFile(ctx context.Context, filePath string) ([]*ProductRecord, error) {

	return nil, nil
}

func (s Seng) NeedsImportFromFile() bool {
	return true
}
