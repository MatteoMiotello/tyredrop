package jobs

import (
	"context"
	"io"
	"os"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/supplier"
	"pillowww/titw/internal/domain/supplier/supplier_factory"
	ftp2 "pillowww/titw/pkg/ftp"
	"pillowww/titw/pkg/log"
)

func CopySupplierFiles() {
	ctx := context.Background()
	sDao := supplier.NewDao(db.DB)

	suppliers, err := sDao.GetAll(ctx)
	check(err)

	ftp, err := ftp2.Init(ctx)
	check(err)

	if err != nil {
		return
	}

	for _, sup := range suppliers {
		if sup.BaseFolder.IsZero() {
			continue
		}

		dirName := sup.BaseFolder.String

		var factory supplier_factory.Importer

		factory = supplier.GetFactory(sup)

		if factory == nil {
			log.Error("Factory not found for supplier with code" + sup.Code)
			return
		}

		if !factory.NeedsImportFromFile() {
			continue
		}

		list, err := ftp.Connection.List(dirName)
		check(err)

		tmpDir := "import/" + dirName
		err = os.MkdirAll(tmpDir, 0777)
		check(err)

		for _, entry := range list {
			fileName := dirName + "/" + entry.Name
			tmpFile := tmpDir + "/" + entry.Name

			if _, err := os.Stat(tmpFile); !os.IsNotExist(err) {
				log.Info("File already exists: ", tmpFile)
				continue
			}

			res, err := ftp.Connection.Retr(fileName)
			check(err)

			buf, err := io.ReadAll(res)
			check(err)

			err = os.WriteFile(tmpFile, buf, 0644)
			check(err)

			err = ftp.Connection.Delete(fileName)
			check(err)
		}
	}

	err = ftp.Quit()
	check(err)
}
