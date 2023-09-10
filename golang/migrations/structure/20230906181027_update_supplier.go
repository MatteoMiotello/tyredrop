package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upUpdateSupplier, downUpdateSupplier)
}

func upUpdateSupplier(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE suppliers ADD COLUMN base_folder varchar NULL")

	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE suppliers ADD COLUMN force_update bool DEFAULT false NOT NULL")

	if err != nil {
		return err
	}

	return nil
}

func downUpdateSupplier(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE suppliers DROP COLUMN base_folder")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE suppliers DROP COLUMN force_update")
	if err != nil {
		return err
	}
	return nil
}
