package supplier_factory

import (
	"context"
	"github.com/friendsofgo/errors"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

type Gun Factory

func (g Gun) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
	records, err := utils.CsvReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var recordSlice pdtos.ProductDtoSlice

	for _, record := range records {
		slices := strings.Split(record[0], ";")

		var pRecord = &pdtos.Tyre{}
		var err error = nil

		for i, slice := range slices {
			err = matchRecords(pRecord, i, slice)
		}

		if err != nil {
			continue
		}

		if pRecord.Construction == "" {
			pRecord.Construction = "R"
		}

		recordSlice = append(recordSlice, pRecord)
	}

	return recordSlice, nil
}

func (g Gun) NeedsImportFromFile() bool {
	return true
}

func matchRecords(pRecord *pdtos.Tyre, index int, slice string) error {
	var err error

	switch index {
	case 1:
		pRecord.EANCode = slice
		break
	case 3:
		pRecord.Brand = slice
		break
	case 4:
		pRecord.ProductName = slice
		break
	case 8:
		pRecord.Season = getSeason(slice)
		break
	case 10:
		w, err := strconv.ParseFloat(slice, 32)

		pRecord.Width = int(w)

		if err != nil {
			return err
		}

		break
	case 11:
		a, err := strconv.ParseFloat(slice, 32)

		pRecord.AspectRatio = int(a)

		if err != nil {
			return err
		}

		break
	case 12:
		r, err := strconv.ParseFloat(slice, 32)

		pRecord.Rim = int(r)
		if err != nil {
			return err
		}

		break
	case 13:
		pRecord.Load, err = strconv.Atoi(slice)
		if err != nil {
			return err
		}
		break
	case 15:
		pRecord.Speed = slice
		break
	case 19:
		pRecord.Price = slice
		break
	case 20:
		f, err := strconv.ParseFloat(slice, 32)

		if err != nil {
			return err
		}

		pRecord.Quantity = int(f)
	case 42:
		pRecord.EprelID = extractEprelCode(slice)
		break
	}

	return nil
}

func getSeason(slice string) string {
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

func extractEprelCode(slice string) string {
	if slice == "" {
		return ""
	}

	splitted := strings.Split(slice, "/")

	if len(splitted) != 5 {
		return ""
	}

	return splitted[4]
}

func extractDimensions(slice string) (*pdtos.TyreDimension, error) {
	r := regexp.MustCompile("/([0-9]){2,3}/([0-9]{2,3})[A-Z]([0-9]){2,3} ([0-9]){2,3}([A-Z])(.*)/")

	match := r.Match([]byte(slice))

	if !match {
		return nil, errors.New("slice not match tyre pattern")
	}

	return nil, nil
}
