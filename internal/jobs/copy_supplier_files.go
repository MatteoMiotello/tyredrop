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
	"strings"
)

func CopySupplierFiles() {
	ctx := context.Background()
	sDao := supplier.NewDao(db.DB)

	suppliers, err := sDao.GetAll(ctx)
	check(err)

	ftp, err := ftp2.Init(ctx)
	check(err)

	for _, sup := range suppliers {
		dirName := strings.ToLower(sup.Code)

		var factory supplier_factory.Importer

		factory = supplier.GetFactory(sup)
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
		}
	}

	err = ftp.Quit()
	check(err)
}
