package supplier_factory

import (
	"context"
	"fmt"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/pkg/utils"
	"strings"
)

type Tyre24 Factory

func (f Tyre24) ReadProductsFromFile(ctx context.Context, filePath string) (pdtos.ProductDtoSlice, error) {
	records, err := utils.CsvReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	var recordSlice pdtos.ProductDtoSlice

	for _, record := range records {
		slices := strings.Split(record[0], ";")

		var pRecord = &pdtos.Tyre{}

		for i, slice := range slices {
			err := f.matchRecords(pRecord, i, slice)

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

func (f Tyre24) NeedsImportFromFile() bool {

}

func (f Tyre24) matchRecords(record *pdtos.Tyre, index int, slice string) error {
	switch index {
		
	}
}
