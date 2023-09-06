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
		w, err := strconv.ParseFloat(slice, 64)

		if err == nil {
			pRecord.Width = int(w)
		}

		break
	case 5:
		a, err := strconv.ParseFloat(slice, 64)

		if err == nil {
			pRecord.AspectRatio = int(a)
			return err
		}

		break
	case 6:
		r, err := strconv.ParseFloat(slice, 64)

		if err == nil {
			pRecord.Rim = r
		}

		break
	case 8:
		pRecord.Reference = slice

		name, err := extractNameFromReference(slice)

		if err == nil {
			pRecord.ProductName = name
		}

		break
	case 9:
		if err == nil {
			pRecord.Load, err = strconv.Atoi(slice)
		}
		break
	case 10:
		pRecord.VehicleType = getVehicleTypeFromItalian(slice)
		break
	case 11:
		pRecord.Season = getSeasonFromItalian(slice)
		break
	case 13:
		if len(slice) > 0 {
			pRecord.RunFlat = true
		} else {
			pRecord.RunFlat = false
		}
		break
	case 14:
		pRecord.Price = slice
		break
	case 16:
		f, err := strconv.ParseFloat(slice, 64)

		if err == nil {
			pRecord.Quantity = int(f)
		}
	}

	return nil
}
