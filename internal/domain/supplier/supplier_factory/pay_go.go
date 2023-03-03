package supplier_factory

import (
	"context"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
	"strconv"
	"strings"
)

type PayGo Factory

func (g PayGo) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
	records, err := utils.CsvReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var recordSlice pdtos.ProductDtoSlice

	for _, record := range records {
		slices := strings.Split(record[0], ";")

		var pRecord = &pdtos.Tyre{}

		for i, slice := range slices {
			err := g.matchRecords(pRecord, i, slice)

			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if pRecord.Construction == "" {
			pRecord.Construction = "R"
		}

		recordSlice = append(recordSlice, pRecord)
	}

	return recordSlice, nil
}

func (g PayGo) matchRecords(pRecord *pdtos.Tyre, index int, slice string) error {
	switch index {
	case 0:
		pRecord.EANCode = slice
		break
	case 3:
		q, err := strconv.Atoi(slice)

		if err != nil {
			return err
		}

		pRecord.Quantity = q

		break
	case 4:
		pRecord.Price = slice
		break
	case 5:
		s := cases.Lower(language.Und).String(slice)
		pRecord.Brand = cases.Title(language.Und).String(s)

		break
	case 7:
		pRecord.Season = g.getSeason(slice)
		break
	case 8:
		pRecord.ProductName = slice
	case 10:
		f, err := strconv.ParseFloat(slice, 32)

		if err != nil {
			return err
		}
		pRecord.Width = int(f)
	case 12:
		f, err := strconv.ParseFloat(slice, 32)

		if err == nil {
			pRecord.AspectRatio = int(f)
		}
		break
	case 14:
		f, err := strconv.ParseFloat(slice, 32)

		if err != nil {
			return err
		}
		pRecord.Rim = int(f)
		break
	case 16:
		f, err := strconv.ParseFloat(slice, 32)

		if err != nil {
			return err
		}
		pRecord.Load = int(f)
		break
	case 17:
		pRecord.Speed = slice
		break
	}
	return nil
}

func (g PayGo) NeedsImportFromFile() bool {
	return true
}

func (g PayGo) getSeason(slice string) string {
	slice = cases.Upper(language.Und).String(slice)

	switch slice {
	case "SUMMER":
		return constants.TYPE_SUMMER
	case "WIN":
		return constants.TYPE_WINTER
	case "4SEASON":
		return constants.TYPE_ALL_SEASON
	case "ALL SEASONS":
		return constants.TYPE_ALL_SEASON
	}
	return ""
}
