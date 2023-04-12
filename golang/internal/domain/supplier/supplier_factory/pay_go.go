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
	case 6:
		pRecord.VehicleType = g.getVehicleType(slice)
	case 7:
		pRecord.Season = g.getSeason(slice)
		break
	case 8:
		pRecord.Reference = slice
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
	case 23:
		if slice != "" {
			pRecord.ProductName = slice
		} else {
			pRecord.ProductName = pRecord.Reference
		}
		break
	}
	return nil
}

func (g PayGo) NeedsImportFromFile() bool {
	return true
}

func (g PayGo) getVehicleType(slice string) constants.VehicleType {
	switch slice {
	case "4 X 4 M&S":
		return constants.VEHICLE_CAR
	case "4X4":
		return constants.VEHICLE_CAR
	case "4x4":
		return constants.VEHICLE_CAR
	case "ALL-SEASON LIGHT-TRUCK/VAN":
		return constants.VEHICLE_TRUCK
	case "ALL-SEASON PW":
		return constants.VEHICLE_CAR
	case "ATV":
		return constants.VEHICLE_QUAD
	case "FURGONETA":
		return constants.VEHICLE_TRUCK
	case "LITRUCK":
		return constants.VEHICLE_TRUCK
	case "LUXE BANDEN":
		return constants.VEHICLE_CAR
	case "LUXE BANDEN M&S":
		return constants.VEHICLE_CAR
	case "MOT":
		return constants.VEHICLE_MOTO
	case "MOTO":
		return constants.VEHICLE_MOTO
	case "N4X4":
		return constants.VEHICLE_CAR
	case "NCTA":
		return constants.VEHICLE_TRUCK
	case "NMOTO":
		return constants.VEHICLE_MOTO
	case "NTURISMO":
		return constants.VEHICLE_CAR
	case "PCR":
		return constants.VEHICLE_CAR
	case "QUAD":
		return constants.VEHICLE_QUAD
	case "quad":
		return constants.VEHICLE_QUAD
	case "RMOTO":
		return constants.VEHICLE_MOTO
	case "RUNFLAT BANDEN":
		return constants.VEHICLE_CAR
	case "RUNFLAT BANDEN M&S":
		return constants.VEHICLE_CAR
	case "TODO TERRENO 4X4":
		return constants.VEHICLE_CAR
	case "TURISMO":
		return constants.VEHICLE_CAR
	}

	return constants.VEHICLE_CAR
}

func (g PayGo) getSeason(slice string) string {
	slice = cases.Upper(language.Und).String(slice)
	slice = strings.ToValidUTF8(slice, "")
	if strings.ContainsAny(slice, "SUM") {
		return constants.TYPE_ALL_SEASON
	}

	if strings.ContainsAny(slice, "WIN") {
		return constants.TYPE_WINTER
	}

	if strings.ContainsAny(slice, "SEAS") {
		return constants.TYPE_ALL_SEASON
	}

	return constants.TYPE_SUMMER
}
