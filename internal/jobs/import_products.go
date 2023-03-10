package jobs

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/null/v8"
	"os"
	currency2 "pillowww/titw/internal/currency"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/import_job"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/internal/domain/supplier"
	"pillowww/titw/internal/domain/supplier/supplier_factory"
	"pillowww/titw/models"
	"pillowww/titw/pkg/log"
	"pillowww/titw/pkg/task"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Error("Error importing file", err.Error())
	}
}

func ImportProductFromFile() {
	ctx := context.Background()
	sDao := supplier.NewDao(db.DB)

	jobExists, _ := sDao.ExistRunningJob(ctx)

	if jobExists {
		return
	}

	sup, err := sDao.GetLastImported(ctx)
	check(err)

	dirName := strings.ToLower(sup.Code)
	tmpDir := "import/" + dirName

	sup.ImportedAt = null.TimeFrom(time.Now())
	err = sDao.Update(ctx, sup)

	check(err)

	var factory supplier_factory.Importer

	factory = supplier.GetFactory(sup)
	if !factory.NeedsImportFromFile() {
		return
	}

	list, err := os.ReadDir(tmpDir)
	if err != nil {
		return
	}

	for _, entry := range list {
		if entry.IsDir() {
			continue
		}

		exists, err := sDao.ExistsJobForFilename(ctx, *sup, entry.Name())

		if exists {
			check(err)
			continue
		}

		ijService := import_job.NewImportJobService(import_job.NewDao(db.DB))
		jobModel, err := ijService.CreateJob(ctx, *sup, entry.Name())
		check(err)

		fileName := tmpDir + "/" + entry.Name()

		_ = ijService.StartNow(ctx, jobModel)

		records, err := factory.ReadProductsFromFile(ctx, fileName)

		if err != nil {
			log.Error("error reading from file: " + entry.Name())
			_ = ijService.EndNowWithError(ctx, jobModel, err.Error())
			break
		}

		if len(records) == 0 {
			log.Warn("no records found for file: " + entry.Name())
			_ = ijService.EndNowWithError(ctx, jobModel, err.Error())
			break
		}

		storeBrands(ctx, records)

		err = storeRecords(ctx, sup, records)
		_ = ijService.EndNow(ctx, jobModel)

		check(err)

		err = os.Remove(fileName)

		check(err)

		break
	}
}

func storeBrands(ctx context.Context, records []pdtos.ProductDto) {
	var uniqueB = make(map[string]string)
	bDao := brand.NewDao(db.DB)
	bService := brand.NewBrandService(bDao)

	for _, record := range records {
		if !record.Validate() {
			continue
		}

		if _, found := uniqueB[record.GetBrandName()]; found {
			continue
		}

		uniqueB[record.GetBrandName()] = record.GetBrandName()
	}

	for _, brandName := range uniqueB {
		_, err := bService.FindOrCreateBrand(ctx, brandName)
		if err != nil {
			log.Error("Error creating brand: ", err)
		}
	}
}

func storeRecords(ctx context.Context, sup *models.Supplier, records []pdtos.ProductDto) error {
	rChan := make(chan pdtos.ProductDto)
	chanWorker := task.NewChannelWorker[pdtos.ProductDto](24, rChan)
	chanWorker.Run(func(record pdtos.ProductDto) {
		importNextRecord(ctx, sup, record)
	})

	for _, record := range records {
		chanWorker.InsertToChannel(record)
	}

	close(rChan)

	return nil
}

func importNextRecord(ctx context.Context, sup *models.Supplier, record pdtos.ProductDto) {
	if !record.Validate() {
		return
	}

	err := db.WithTx(ctx, func(tx *sql.Tx) error {
		dao := product.NewDao(tx)
		bDao := brand.NewDao(tx)
		cDao := currency2.NewDao(tx)
		pService := product.NewService(dao, bDao, cDao)

		p, err := pService.FindOrCreateProduct(ctx, record)
		if err != nil {
			return err
		}

		err = pService.UpdateSpecifications(ctx, p, record)
		if err != nil {
			return err
		}

		pi, err := pService.CreateProductItem(ctx, p, sup, record.GetSupplierProductPrice(), record.GetSupplierProductQuantity())

		if err != nil {
			return err
		}

		err = pService.CalculateAndStoreProductPrices(ctx, pi)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
