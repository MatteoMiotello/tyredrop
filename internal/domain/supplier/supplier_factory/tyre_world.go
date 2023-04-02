package supplier_factory

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/log"
	"strconv"
)

type TyreWorld Factory

func (t TyreWorld) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Warn("failed closing xlsx file: ", err)
		}
	}()

	var dtos pdtos.ProductDtoSlice
	sheet := f.GetSheetName(0)

	rows, err := f.GetRows(sheet)

	if err != nil {
		return nil, err
	}

	rowsNum := len(rows)

	for i := 1; i <= rowsNum; i++ {
		dto, err := t.matchCells(f, sheet, i)

		if err != nil {
			fmt.Println(err)
			continue
		}

		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (t TyreWorld) matchCells(f *excelize.File, sheet string, rowNumber int) (*pdtos.Tyre, error) {
	pDto := &pdtos.Tyre{}

	brand, err := f.GetCellValue(sheet, "A"+strconv.Itoa(rowNumber))
	if err != nil {
		return nil, err
	}
	brand = cases.Lower(language.Und).String(brand)
	pDto.Brand = cases.Title(language.Und).String(brand)

	ean, err := f.GetCellValue(sheet, "B"+strconv.Itoa(rowNumber))
	if err != nil {
		return nil, err
	}
	pDto.EANCode = ean

	name, err := f.GetCellValue(sheet, "C"+strconv.Itoa(rowNumber))
	if err != nil {
		return nil, err
	}
	pDto.Reference = name
	pDto.ProductName, _ = extractNameFromReference(name)
	dimensions, err := extractDimensionsFromName(name)

	if err != nil {
		return nil, err
	}
	pDto.TyreDimension = *dimensions

	quantity, err := f.GetCellValue(sheet, "E"+strconv.Itoa(rowNumber))
	if err != nil {
		return nil, err
	}
	pDto.Quantity, _ = strconv.Atoi(quantity)

	price, err := f.GetCellValue(sheet, "F"+strconv.Itoa(rowNumber))
	if err != nil {
		return nil, err
	}
	pDto.Price = price

	pDto.VehicleType = constants.VEHICLE_CAR

	return pDto, nil
}

func (t TyreWorld) NeedsImportFromFile() bool {
	return true
}
