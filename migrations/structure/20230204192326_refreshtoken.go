package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upRefreshtoken, downRefreshtoken)
}

func upRefreshtoken(tx *sql.Tx) error {
	refreshTokenQuery := sqlbuilder.CreateTable("public.refresh_tokens").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		Column("refresh_token", types.Varchar.Options("500"), false).
		Column("expires_at", types.Timestamptz, false).
		Column("time_last_use", types.Timestamptz, true).
		DeletedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(refreshTokenQuery)
	if err != nil {
		return err
	}
	return nil
}

func downRefreshtoken(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.refresh_tokens")
	if err != nil {
		return err
	}
	return nil
}
