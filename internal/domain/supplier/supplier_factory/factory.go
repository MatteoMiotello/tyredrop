package supplier_factory

import (
	"context"
	"pillowww/titw/models"
)

type Importer interface {
	ReadProductsFromFile(ctx context.Context, filePath string) ([]*ProductRecord, error)
	NeedsImportFromFile() bool
}

type Factory struct {
	Importer
	S *models.Supplier
}

type ProductRecord struct {
	EANNumber   int
	ProductName string
	Brand       string
	Season      string
	Price       string
	EprelID     string
	ProductDimension
}

type ProductDimension struct {
	Width        int
	AspectRatio  int
	Construction string
	Rim          int
	Load         int
	Speed        string
}
