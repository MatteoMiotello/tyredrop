package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upBrandQuality, downBrandQuality)
}

func upBrandQuality(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE public.brands ADD COLUMN quality int default null")
	if err != nil {
		return err
	}

	return nil
}

func downBrandQuality(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
