package jobs

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
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
	"pillowww/titw/internal/domain/vehicle"
	"pillowww/titw/models"
	"pillowww/titw/pkg/log"
	"pillowww/titw/pkg/task"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Error("Error importing file", err)
	}
}

func ImportProductsFromFile() {
	ctx := context.Background()

	_, err := product.NewPriceMarkupDao(db.DB).FindPriceMarkupDefault(ctx)
	if err != nil {
		log.Error("Default Price markup not found")
		return
	}

	sDao := supplier.NewDao(db.DB)
	jobExists, _ := sDao.ExistRunningJob(ctx)

	if jobExists {
		fmt.Println("one job is running")
		return
	}

	sup, err := sDao.GetLastImported(ctx)
	check(err)

	if sup == nil {
		log.Error(errors.New("supplier not found in import"))
		return
	}

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

		err = ijService.StartNow(ctx, jobModel)
		check(err)

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

		check(err)

		err = ijService.EndNow(ctx, jobModel)

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
	chanWorker := task.NewChannelWorker[pdtos.ProductDto](50, rChan)

	chanWorker.Run(func(record pdtos.ProductDto) {
		importNextRecord(ctx, sup, record)
	})

	for _, record := range records {
		chanWorker.InsertToChannel(record)
	}

	//
	//var codes []string
	//
	//for _, r := range records {
	//	codes = append(codes, r.GetProductCode())
	//}
	//
	//dao := product.NewItemDao(db.DB)
	//err := dao.RemoveOldItems(ctx, sup, codes)
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

func importNextRecord(ctx context.Context, sup *models.Supplier, record pdtos.ProductDto) {

	if !record.Validate() {
		return
	}

	err := db.WithTx(ctx, func(tx *sql.Tx) error {
		dao := product.NewDao(tx)
		itemDao := product.NewItemDao(tx)
		pService := product.NewService(
			dao,
			brand.NewDao(tx),
			product.NewCategoryDao(tx),
			itemDao,
			product.NewSpecificationDao(tx),
			product.NewSpecificationValueDao(tx),
			vehicle.NewDao(tx),
		)

		pPriceService := product.NewPriceService(
			dao,
			itemDao,
			product.NewPriceMarkupDao(tx),
			currency2.NewDao(tx),
			product.NewItemPriceDao(tx),
		)

		p, err := pService.FindOrCreateProduct(ctx, record)
		if err != nil {
			return err
		}

		err = pService.UpdateSpecifications(ctx, p, record)
		if err != nil {
			return err
		}

		pi, err := pService.CreateOrUpdateProductItem(ctx, p, sup, record.GetSupplierProductPrice(), record.GetSupplierProductQuantity())
		if err != nil {
			return err
		}

		err = pPriceService.CalculateAndStoreProductPrices(ctx, pi)
		if err != nil {
			return err
		}

		err = pPriceService.CalculatePriceAdditions(ctx, pi, record)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("end of input", err.Error())
	}
}
