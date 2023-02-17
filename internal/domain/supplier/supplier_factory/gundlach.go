package supplier_factory

import (
	"context"
)

type Gun Factory

func (g Gun) ImportProductsFromFile(ctx context.Context, filePath string) error {

	return nil
}

func (g Gun) NeedsImportFromFile() bool {
	return true
}
