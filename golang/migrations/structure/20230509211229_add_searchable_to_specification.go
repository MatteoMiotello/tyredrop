package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddSearchableToSpecification, downDirMigrationsStructure)
}

func upAddSearchableToSpecification(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE public.product_specifications ADD COLUMN searchable bool not null default false")
	if err != nil {
		return err
	}

	return nil
}

func downDirMigrationsStructure(tx *sql.Tx) error {
	return nil
}
