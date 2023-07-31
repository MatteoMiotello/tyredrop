package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upProductEprelCode, downProductEprelCode)
}

func upProductEprelCode(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE public.products ADD COLUMN eprel_product_code varchar default null")
	if err != nil {
		return err
	}

	return nil
}

func downProductEprelCode(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
