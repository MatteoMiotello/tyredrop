package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upTax, downTax)
}

func upTax(tx *sql.Tx) error {
	query, params := sqlbuilder.InsertInto("taxes").
		Cols("markup_percentage", "name").
		Values(22, "IVA").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func downTax(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
