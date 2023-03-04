package jobs

import (
	"context"
	"database/sql"
	"errors"
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

		log.WithField("entry", entry).Info("Importing file")

		if exists {
			log.WithField("entry", entry).Warn("File already imported")
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

		storeBrands(ctx, records)

		storeProducts(ctx, records)

		err = storeRecords(ctx, sup, records)
		_ = ijService.EndNow(ctx, jobModel)

		check(err)

		err = os.Remove(fileName)

		check(err)

		break
	}
}

func storeProducts(ctx context.Context, records []pdtos.ProductDto) {
	var uniqueP = make(map[string]pdtos.ProductDto)
	dao := product.NewDao(db.DB)
	bDao := brand.NewDao(db.DB)
	cDao := currency2.NewDao(db.DB)
	pService := product.NewService(dao, bDao, cDao)

	for _, record := range records {
		if !record.Validate() {
			continue
		}

		if _, found := uniqueP[record.GetProductCode()]; found {
			continue
		}

		uniqueP[record.GetProductCode()] = record
	}

	for _, record := range uniqueP {
		_, err := pService.FindOrCreateProduct(ctx, record)
		if err != nil {
			log.Error("Error creating product", err)
		}
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
	worker := task.NewWorker(8)
	worker.Run()

	for _, record := range records {
		worker.AddTask(func() {
			importNextRecord(ctx, sup, record)
		})
	}
	return nil
}

func importNextRecord(ctx context.Context, sup *models.Supplier, record pdtos.ProductDto) {
	err := db.WithTx(ctx, func(tx *sql.Tx) error {
		if !record.Validate() {
			return errors.New("record is not valid")
		}

		dao := product.NewDao(tx)
		bDao := brand.NewDao(tx)
		cDao := currency2.NewDao(tx)
		pService := product.NewService(dao, bDao, cDao)

		p, err := dao.FindOneByProductCode(ctx, record.GetProductCode())
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
