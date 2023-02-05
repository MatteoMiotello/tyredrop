package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upSupplier, downSupplier)
}

func upSupplier(tx *sql.Tx) error {
	supplierQuery := sqlbuilder.CreateTable("public.suppliers").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		FKColumn("public.user_billings", "user_billing_id", false).
		FKColumn("public.currencies", "currency_id", false).
		Column("code", types.Varchar.Options("45"), false).
		Column("ftp_username", types.Varchar.Options("255"), true).
		Column("ftp_password", types.Varchar.Options("255"), true).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(supplierQuery)
	if err != nil {
		return err
	}

	return nil
}

func downSupplier(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.suppliers ")
	if err != nil {
		return err
	}
	return nil
}
