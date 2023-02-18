package jobs

import (
	"context"
	"database/sql"
	"fmt"
	ftp3 "github.com/jlaffaye/ftp"
	"github.com/volatiletech/null/v8"
	"io"
	"os"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/import_job"
	"pillowww/titw/internal/domain/supplier"
	"pillowww/titw/internal/domain/supplier/supplier_factory"
	"pillowww/titw/models"
	ftp2 "pillowww/titw/pkg/ftp"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ImportProductFromFile() {
	ctx := context.Background()
	sDao := supplier.NewDao(db.DB)

	sup, err := sDao.GetLastImported(ctx)
	dirName := strings.ToLower(sup.Code)

	check(err)

	sup.ImportedAt = null.TimeFrom(time.Now())
	err = sDao.Update(ctx, sup)

	check(err)

	var factory supplier_factory.Importer

	ftp, err := ftp2.Init(ctx)
	defer ftp.Quit()

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

		if exists {
			fmt.Println("job already exists")
			continue
		}

		err = db.WithTx(ctx, func(tx *sql.Tx) error {
			ijService := import_job.NewImportJobService(import_job.NewDao(tx))
			jobModel, err := ijService.CreateJob(ctx, *sup, entry.Name)
			if err != nil {
				return err
			}

			fileName := dirName + "/" + entry.Name
			tmpDir := "import/" + dirName
			tmpFile := tmpDir + "/" + entry.Name

			res, err := ftp.Connection.Retr(fileName)
			if err != nil {
				return err
			}

			buf, err := io.ReadAll(res)
			if err != nil {
				return err
			}

			err = os.MkdirAll(tmpDir, 0777)
			if err != nil {
				return err
			}

			err = os.WriteFile(tmpFile, buf, 0644)
			if err != nil {
				return err
			}

			_ = ijService.StartNow(ctx, jobModel)

			records, err := factory.ReadProductsFromFile(ctx, tmpFile)

			if err != nil {
				_ = ijService.EndNowWithError(ctx, jobModel, err.Error())
				return nil
			}

			storeRecords(sup, records)

			_ = ijService.EndNow(ctx, jobModel)

			err = os.Remove(tmpFile)

			if err != nil {
				return err
			}

			return nil
		})

		check(err)

		break
	}
}

func storeRecords(sup *models.Supplier, records []*supplier_factory.ProductRecord) {

}
