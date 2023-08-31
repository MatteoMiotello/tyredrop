package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upInvoiceTable, downInvoiceTable)
}

func upInvoiceTable(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.invoices").
		PKColumn().
		FKColumn("public.user_billings", "user_billing_id", false).
		Column("number", types.Varchar.Options("255"), false).
		Column("file_path", types.Varchar.Options("255"), false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downInvoiceTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS public.invoices ")
	if err != nil {
		return err
	}
	return nil
}
