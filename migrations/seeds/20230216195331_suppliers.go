package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upSuppliers, downSuppliers)
}

func upSuppliers(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.suppliers").
		Cols("name", "code").
		Values("Seng", "SENG").
		Values("Gundlach", "GUN").
		Values("Tyre World", "TYRE_WORLD").
		Values("PAY&GO", "PAY_GO").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func downSuppliers(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
