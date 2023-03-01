package jobs

import (
	"context"
	"database/sql"
	"errors"
	ftp3 "github.com/jlaffaye/ftp"
	"github.com/volatiletech/null/v8"
	"io"
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
	ftp2 "pillowww/titw/pkg/ftp"
	"pillowww/titw/pkg/log"
	"pillowww/titw/pkg/task"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Panic("Error importing file", err.Error())
	}
}

func ImportProductFromFile() {
	ctx := context.Background()
	sDao := supplier.NewDao(db.DB)

	sup, err := sDao.GetLastImported(ctx)
	check(err)

	dirName := strings.ToLower(sup.Code)

	sup.ImportedAt = null.TimeFrom(time.Now())
	err = sDao.Update(ctx, sup)

	check(err)

	var factory supplier_factory.Importer

	ftp, err := ftp2.Init(ctx)

	check(err)

	factory = supplier.GetFactory(sup)
	if !factory.NeedsImportFromFile() {
		return
	}

	list, err := ftp.Connection.List(dirName)
	check(err)

	for _, entry := range list {
		if entry.Type != ftp3.EntryTypeFile {
			continue
		}

		exists, err := sDao.ExistsJobForFilename(ctx, *sup, entry.Name)

		log.WithField("entry", entry).Info("Importing file")

		if exists {
			log.WithField("entry", entry).Warn("File already imported")
			continue
		}

		ijService := import_job.NewImportJobService(import_job.NewDao(db.DB))
		jobModel, err := ijService.CreateJob(ctx, *sup, entry.Name)
		check(err)

		fileName := dirName + "/" + entry.Name
		tmpDir := "import/" + dirName
		tmpFile := tmpDir + "/" + entry.Name

		res, err := ftp.Connection.Retr(fileName)
		check(err)

		buf, err := io.ReadAll(res)
		check(err)

		err = os.MkdirAll(tmpDir, 0777)
		check(err)

		err = os.WriteFile(tmpFile, buf, 0644)
		check(err)

		ftp.Quit()

		_ = ijService.StartNow(ctx, jobModel)

		records, err := factory.ReadProductsFromFile(ctx, tmpFile)

		if err != nil {
			log.Warn("error reading from file: " + tmpFile)
			_ = ijService.EndNowWithError(ctx, jobModel, err.Error())
			break
		}

		err = storeRecords(ctx, sup, records)
		_ = ijService.EndNow(ctx, jobModel)

		check(err)

		err = os.Remove(tmpFile)

		check(err)

		break
	}
}

func importNextRecord(ctx context.Context, sup *models.Supplier, record pdtos.ProductDto) {
	db.WithTx(ctx, func(tx *sql.Tx) error {
		dao := product.NewDao(tx)
		bDao := brand.NewDao(tx)
		cDao := currency2.NewDao(tx)
		pService := product.NewService(dao, bDao, cDao)

		if !record.Validate() {
			return errors.New("record is not valid")
		}

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
