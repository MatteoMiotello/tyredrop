package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddSupplier, downAddSupplier)
}

func upAddSupplier(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.suppliers").
		Cols("name", "code").
		Values("Tyre24", "TYRE_24").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func downAddSupplier(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
