package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upTyreCategory, downTyreCategory)
}

func upTyreCategory(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.product_categories").
		Cols("category_code").
		Values("TYRES").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}

	tyreCatQuery, _ := sqlbuilder.
		Select("id").
		From("public.product_categories").
		Where("category_code = 'TYRES'").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	var id int64
	row := tx.QueryRow(tyreCatQuery)

	err = row.Scan(&id)
	if err != nil {
		return err
	}

	query, values = sqlbuilder.InsertInto("public.product_category_languages").
		Cols("language_id", "product_category_id", "name").
		Values(1, id, "Pneumatici").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func downTyreCategory(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
