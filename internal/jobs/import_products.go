package jobs

import (
	"context"
	ftp3 "github.com/jlaffaye/ftp"
	"github.com/volatiletech/null/v8"
	"google.golang.org/appengine/log"
	"io"
	"os"
	"pillowww/titw/internal/domain/supplier"
	"pillowww/titw/internal/domain/supplier/supplier_factory"
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
	sRepo := supplier.NewRepository()

	sup, err := sRepo.GetLastImported(ctx)
	dirName := strings.ToLower(sup.Code)

	check(err)

	sup.ImportedAt = null.TimeFrom(time.Now())
	err = sRepo.Update(ctx, sup)

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

		_, err := sRepo.GetJobsFromFilename(ctx, *sup, entry.Name)

		if err != nil {
			log.Infof(ctx, "job already present for file: %s", entry.Name)
			continue
		}

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

		err = factory.ImportProductsFromFile(ctx, tmpFile)
		check(err)

		err = os.Remove(tmpFile)
		check(err)

		break
	}
}
