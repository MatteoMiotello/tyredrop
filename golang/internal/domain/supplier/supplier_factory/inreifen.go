package supplier_factory

import (
	"context"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
	"strconv"
	"strings"
)

type Inreifen Factory

func (g Inreifen) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
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

		recordSlice = append(recordSlice, pRecord)
	}

	return recordSlice, nil
}

func (g Inreifen) NeedsImportFromFile() bool {
	return true
}

func (g Inreifen) matchRecords(pRecord *pdtos.Tyre, index int, slice string) error {
	switch index {
	case 1:
		pRecord.EANCode = slice
	case 3:
		pRecord.Reference = slice
	case 4:
		pRecord.Brand = strings.ToUpper(slice)
	case 5:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.Width = int(f)
	case 6:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.AspectRatio = int(f)
	case 7:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.Rim = f
	case 8:
		pRecord.Speed = slice
	case 9:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.Load = int(f)
	case 12:
		pRecord.WetGripClass = slice
	case 13:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.ExternalRollingNoiseLevel = int(f)
	case 16:
		pRecord.VehicleType = constants.VEHICLE_CAR //todo
	case 17:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.Quantity = int(f)
	case 18:
		pRecord.Price = slice
	}

	return nil
}
