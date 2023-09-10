package supplier_factory

import (
	"context"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
	"strconv"
	"strings"
)

type Pneusdata Factory

func (g Pneusdata) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
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

func (g Pneusdata) NeedsImportFromFile() bool {
	return true
}

func (g Pneusdata) matchRecords(pRecord *pdtos.Tyre, index int, slice string) error {
	switch index {
	case 1:
		pRecord.VehicleType = constants.VEHICLE_CAR
	case 3:
		pRecord.EANCode = slice
	case 4:
		pRecord.Reference = slice
	case 5:
		pRecord.ProductName = slice
	case 6:
		pRecord.Brand = slice
	case 7:
		pRecord.EprelID = slice
	case 8:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.Width = int(f)
	case 9:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.AspectRatio = int(f)
	case 10:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.Rim = f
	case 11:
		pRecord.Season = matchSeason(slice)
	case 12:
		pRecord.Speed = slice
	case 13:
		pRecord.FuelEfficiency = slice
	case 14:
		pRecord.WetGripClass = slice
	case 15:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.Load = int(f)
	case 16:
		f, err := strconv.ParseFloat(slice, 64)

		if err != nil {
			return err
		}

		pRecord.ExternalRollingNoiseLevel = int(f)
	case 17:
		pRecord.ExternalRollingNoiseClass = slice
	case 18:
		if slice == "0" {
			pRecord.RunFlat = false
		} else {
			pRecord.RunFlat = true
		}
	case 25:
		pRecord.ImageUrl = &slice
	}

	pRecord.Construction = "R"
	return nil
}

func matchSeason(slice string) string {
	switch slice {
	case "E":
		return "SUMMER"
	case "I":
		return "WINTER"
	case "4":
		return "ALL_SEASON"
	}

	return "SUMMER"
}
