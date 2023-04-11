package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upProductName, downProductName)
}

func upProductName(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE public.products ADD \"name\" varchar(255) NULL;")
	if err != nil {
		return err
	}
	return nil
}

func downProductName(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
