package supplier_factory

import (
	"context"
	"errors"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
	"regexp"
	"strings"
)

type Importer interface {
	ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error)
	NeedsImportFromFile() bool
}

type Factory struct {
	Importer
	S *models.Supplier
}

func getSeasonFromGerman(slice string) string {
	switch slice {
	case "SOMMER":
		return constants.TYPE_SUMMER
	case "WINTER":
		return constants.TYPE_WINTER
	case "ANHÄNGER":
		return constants.TYPE_TRAILER
	case "GANZJAHR":
		return constants.TYPE_ALL_SEASON
	}
	return ""
}

func extractEprelIDFromLink(slice string) string {
	if slice == "" {
		return ""
	}

	splitted := strings.Split(slice, "/")

	if len(splitted) != 5 {
		return ""
	}

	return splitted[4]
}

func extractDimensionsFromName(slice string) (*pdtos.TyreDimension, error) {
	r := regexp.MustCompile("/([0-9]){2,3}/([0-9]{2,3})[A-Z]([0-9]){2,3} ([0-9]){2,3}([A-Z])(.*)/")

	match := r.Match([]byte(slice))

	if !match {
		return nil, errors.New("slice not match tyre pattern")
	}

	return nil, nil
}
