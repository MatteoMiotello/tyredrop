package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCurrency, downCurrency)
}

func upCurrency(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.currencies").
		Cols("iso_code", "symbol", "tag", "precision", "created_at").
		Values("EUR", "€", "EURO", 1000, sqlbuilder.Raw("NOW()")).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func downCurrency(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
