package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upLanguage, downLanguage)
}

func upLanguage(tx *sql.Tx) error {
	row := tx.QueryRow("SELECT id FROM currencies where iso_code='EUR'")
	var id int64

	err := row.Scan(&id)
	if err != nil {
		return err
	}

	query, values := sqlbuilder.InsertInto("public.languages").
		Cols("currency_id", "name", "iso_code", "tag", "std_timezone", "created_at").
		Values(id, "Italiano", "it", "IT", "CET", sqlbuilder.Raw("NOW()")).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	row = tx.QueryRow("SELECT id FROM languages LIMIT 1")
	var languageId int64
	err = row.Scan(&languageId)
	if err != nil {
		return err
	}

	query, values = sqlbuilder.InsertInto("public.currency_languages").
		Cols("currency_id", "language_id", "name").
		Values(id, languageId, "Euro").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)

	if err != nil {
		return err
	}

	return nil
}

func downLanguage(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
