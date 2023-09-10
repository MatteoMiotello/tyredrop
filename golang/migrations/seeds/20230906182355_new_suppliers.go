package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upNewSuppliers, downNewSuppliers)
}

func upNewSuppliers(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.suppliers").
		Cols("name", "code", "force_update").
		Values("Inreifen", "INREIFEN", false).
		Values("Pneusdata", "PNEUSDATA", true).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func downNewSuppliers(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
