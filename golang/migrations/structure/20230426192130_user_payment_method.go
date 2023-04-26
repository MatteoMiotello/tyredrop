package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upUserPaymentMethod, downUserPaymentMethod)
}

func upUserPaymentMethod(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.user_payment_methods").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		Column("type", types.Varchar.Options("45"), false).
		Column("name", types.Varchar.Options("255"), false).
		Column("value", types.Varchar.Options("500"), false).
		Column("type_primary", types.Bool, false).
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

func downUserPaymentMethod(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
