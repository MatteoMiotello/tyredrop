package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/internal/domain/vehicle"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/utils"
	"strconv"
	"strings"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
	bootstrap.InitLog("cron")
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

func matchRecord(pRecord *pdtos.Tyre, index int, slice string) error {
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

func main() {
	ctx := context.Background()
	filePath := flag.String("file", "", "file to import")
	flag.Parse()
	fmt.Println(*filePath)

	if len(*filePath) < 1 {
		panic("File path is required")
	}

	records, err := utils.CsvReadFile(*filePath)

	if err != nil {
		panic(err)
	}

	totals := len(records)

	for current, record := range records {
		if current == 0 {
			continue
		}
		fmt.Println(fmt.Sprintf("%d/%d", current, totals))

		slices := strings.Split(record[0], ";")

		var pRecord = &pdtos.Tyre{}

		for i, slice := range slices {
			slice = strings.Trim(slice, "\"")

			err := matchRecord(pRecord, i, slice)

			if err != nil {
				fmt.Println(err)

				continue
			}
		}

		err := db.WithTx(ctx, func(tx *sql.Tx) error {
			brandDao := brand.NewDao(tx)

			service := product.NewService(
				product.NewDao(tx),
				brandDao,
				product.NewCategoryDao(tx),
				product.NewItemDao(tx),
				product.NewSpecificationDao(tx),
				product.NewSpecificationValueDao(tx),
				vehicle.NewDao(tx),
			)

			bService := brand.NewBrandService(
				brandDao,
			)

			_, err := bService.FindOrCreateBrand(ctx, pRecord.Brand)
			if err != nil {
				return err
			}

			p, err := service.FindOrCreateProduct(ctx, pRecord)
			if err != nil {
				return err
			}

			err = service.UpdateSpecifications(ctx, p, pRecord)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}
