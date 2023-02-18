package supplier_factory

import (
	"context"
	"fmt"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
	"strconv"
	"strings"
)

type Gun Factory

func (g Gun) ReadProductsFromFile(ctx context.Context, filePath string) ([]*ProductRecord, error) {
	records, err := utils.CsvReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var recordSlice []*ProductRecord

	for _, record := range records {
		slices := strings.Split(record[0], ";")

		var pRecord = &ProductRecord{}

		for i, slice := range slices {
			matchRecords(pRecord, i, slice)
		}

		if pRecord.EANNumber == 0 {
			continue
		}

		fmt.Println(*pRecord)
		recordSlice = append(recordSlice, pRecord)
	}

	return recordSlice, nil
}

func (g Gun) NeedsImportFromFile() bool {
	return true
}

func matchRecords(pRecord *ProductRecord, index int, slice string) {
	switch index {
	case 1:
		pRecord.EANNumber, _ = strconv.Atoi(slice)
		break
	case 3:
		pRecord.Manufacturer = slice
		break
	case 4:
		pRecord.ProductName = slice
		break
	case 8:
		pRecord.Season = getSeason(slice)
	case 19:
		pRecord.Price = slice
	case 42:
		pRecord.EprelID = extractEprelCode(slice)
	}
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
