package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upImportJobs, downImportJobs)
}

func upImportJobs(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("import_jobs").
		PKColumn().
		FKColumn("public.suppliers", "supplier_id", false).
		Column("filename", types.Varchar.Options("255"), false).
		Column("error_message", types.Varchar.Options("255"), true).
		Column("started_at", types.Timestamptz, true).
		Column("ended_at", types.Timestamptz, true).
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downImportJobs(tx *sql.Tx) error {
	tx.Exec("DROP TABLE public.import_jobs")

	return nil
}
