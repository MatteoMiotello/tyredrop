package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upUserAddresses, downUserAddresses)
}

func upUserAddresses(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.user_address").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		Column("address_line_1", types.Varchar.Options("255"), false).
		Column("address_line_2", types.Varchar.Options("255"), true).
		Column("city", types.Varchar.Options("45"), false).
		Column("province", types.Varchar.Options("45"), false).
		Column("postal_code", types.Varchar.Options("5"), false).
		Column("country", types.Varchar.Options("45"), false).
		Column("is_default", types.Bool, false).
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

func downUserAddresses(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS public.user_address ")
	if err != nil {
		return err
	}
	return nil
}
