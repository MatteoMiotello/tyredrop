package supplier_factory

import (
	"context"
	"encoding/csv"
	"os"
	"pillowww/titw/internal/domain/product/pdtos"
	"strconv"
)

type Gun Factory

func (g Gun) readCsv(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true
	reader.Comma = ';'

	return reader.ReadAll()
}

func (g Gun) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
	records, err := g.readCsv(filePath)

	if err != nil {
		return nil, err
	}

	var recordSlice pdtos.ProductDtoSlice

	for _, record := range records {
		var pRecord = &pdtos.Tyre{}
		var err error = nil

		for i, slice := range record {
			err = g.matchRecords(pRecord, i, slice)

			if err != nil {
				continue
			}
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
