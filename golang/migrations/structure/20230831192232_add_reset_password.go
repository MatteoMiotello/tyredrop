package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddResetPassword, downAddResetPassword)
}

func upAddResetPassword(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.reset_passwords").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		Column("token", types.Varchar.Options("500"), false).
		Column("issued_at", types.Timestamptz, false).
		Column("expiry_at", types.Timestamptz, false).
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

func downAddResetPassword(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS public.reset_passwords")
	if err != nil {
		return err
	}

	return nil
}
