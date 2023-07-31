package supplier_factory

import (
	"context"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
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
			err = g.matchRecords(pRecord, i, slice)

			if err != nil {
				continue
			}
		}

		if pRecord.Construction == "" {
			pRecord.Construction = "R"
		}

		pRecord.VehicleType = constants.VEHICLE_CAR

		recordSlice = append(recordSlice, pRecord)
	}

	return recordSlice, nil
}

func (g Gun) NeedsImportFromFile() bool {
	return true
}

func (g Gun) matchRecords(pRecord *pdtos.Tyre, index int, slice string) error {
	var err error

	switch index {
	case 1:
		pRecord.EANCode = slice
		break
	case 3:
		pRecord.Brand = slice
		break
	case 4:
		pRecord.Reference = slice
		break
	case 8:
		pRecord.Season = getSeasonFromGerman(slice)
		break
	case 10:
		w, err := strconv.ParseFloat(slice, 32)

		if err == nil {
			pRecord.Width = int(w)
		}

		break
	case 11:
		a, err := strconv.ParseFloat(slice, 32)

		if err == nil {
			pRecord.AspectRatio = int(a)
			return err
		}

		break
	case 12:
		r, err := strconv.ParseFloat(slice, 32)

		if err == nil {
			pRecord.Rim = r
		}

		break
	case 13:
		if err == nil {
			pRecord.Load, err = strconv.Atoi(slice)
		}
		break
	case 15:
		pRecord.Speed = slice
		break
	case 17:
		pRecord.ProductName = slice
		break
	case 19:
		pRecord.Price = slice
		break
	case 20:
		f, err := strconv.ParseFloat(slice, 32)

		if err == nil {
			pRecord.Quantity = int(f)
		}

	case 42:
		pRecord.EprelID = extractEprelIDFromLink(slice)
		break
	}

	return nil
}

// llkv -> furgoni
