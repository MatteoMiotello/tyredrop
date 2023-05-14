package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddUserConfirmedColumn, downAddUserConfirmedColumn)
}

func upAddUserConfirmedColumn(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE public.users ADD COLUMN confirmed bool NOT NULL default false")
	if err != nil {
		return err
	}

	return nil
}

func downAddUserConfirmedColumn(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
