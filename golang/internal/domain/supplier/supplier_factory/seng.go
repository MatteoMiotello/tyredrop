package supplier_factory

import (
	"context"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
	"strconv"
	"strings"
)

type Seng Factory

func (s Seng) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
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
			slice = s.purgeSlice(slice)
			err = s.matchRecords(pRecord, i, slice)

			if err != nil {
				continue
			}
		}

		pRecord.VehicleType = constants.VEHICLE_CAR

		recordSlice = append(recordSlice, pRecord)
	}

	return recordSlice, nil
}

func (s Seng) purgeSlice(slice string) string {
	return strings.Trim(slice, "\"")
}

func (s Seng) matchRecords(pRecord *pdtos.Tyre, i int, slice string) error {
	switch i {
	case 1:
		pRecord.EANCode = slice
		break
	case 2:
		w, err := strconv.Atoi(slice)

		if err == nil {
			pRecord.Width = w
		}

		break
	case 3:
		a, err := strconv.Atoi(slice)

		if err == nil {
			pRecord.AspectRatio = a
		}

		break
	case 4:
		pRecord.Construction = slice
		break
	case 5:
		a, err := strconv.ParseFloat(slice, 64)

		if err == nil {
			pRecord.Rim = a
		}

		break

	case 6:
		a, err := strconv.Atoi(slice)

		if err == nil {
			pRecord.Load = a
		}

		break
	case 7:
		pRecord.Speed = slice
		break
	case 9:
		pRecord.Reference = pRecord.BuildName() + " " + slice
		pRecord.ProductName = slice
		break
	case 10:
		s := cases.Lower(language.Und).String(slice)
		pRecord.Brand = cases.Title(language.Und).String(s)
		break
	case 12:
		q, err := strconv.Atoi(slice)

		if err == nil {
			pRecord.Quantity = q
		}

		break
	case 15:
		pRecord.Season = getSeasonFromGerman(strings.ToUpper(slice))
		break
	case 16:
		pRecord.Price = slice
		break
	case 32:
		pRecord.EprelID = extractEprelIDFromLink(slice)
		break
	}

	return nil
}

func (s Seng) NeedsImportFromFile() bool {
	return true
}
