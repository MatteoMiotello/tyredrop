package supplier_factory

import (
	"context"
	"errors"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
	"regexp"
	"strconv"
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
	r := regexp.MustCompile(`^(\d{3})\/(\d{2})\s(\w{1,3})(\d{2})\s(?:\w{2}\s)?(\d{2,3})([A-Z])\s([A-Z])(?:\s([\w-]+)\s)?`)

	match := r.FindStringSubmatch(slice)

	if len(match) > 0 {
		width, _ := strconv.Atoi(match[1])
		aspectRatio, _ := strconv.Atoi(match[2])
		construction := match[3]
		diameter, _ := strconv.Atoi(match[4])
		loadIndex, _ := strconv.Atoi(match[5])
		speedIndex := match[6]

		return &pdtos.TyreDimension{
			Width:        width,
			AspectRatio:  aspectRatio,
			Construction: construction,
			Rim:          diameter,
			Load:         loadIndex,
			Speed:        speedIndex,
		}, nil
	}

	return nil, errors.New("String not matching")
}
