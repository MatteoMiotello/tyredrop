package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddSdiBilling, downAddSdiBilling)
}

func upAddSdiBilling(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE public.user_billings ADD COLUMN sdi_code varchar default null")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.user_billings ADD COLUMN sdi_pec varchar default null")

	if err != nil {
		return err
	}

	return nil
}

func downAddSdiBilling(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
