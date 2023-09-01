package seeds

import (
	"database/sql"
	"fmt"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
	"pillowww/titw/pkg/constants"
)

func init() {
	goose.AddMigration(upAddRunflat, downAddRunflat)
}

func upAddRunflat(tx *sql.Tx) error {
	var catId int64

	query, _ := sqlbuilder.Select("id").
		From("public.product_categories").
		Where("category_code = 'TYRES'").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := tx.QueryRow(query)
	err := row.Scan(&catId)
	if err != nil {
		return err
	}

	query, values := sqlbuilder.InsertInto("public.product_specifications").
		Cols("product_category_id", "specification_code", "type", "mandatory").
		Values(catId, "RUNFLAT", constants.SPEC_TYPE_STRING, true).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	var id int64
	rows, err := tx.Query("SELECT id FROM product_specifications ORDER BY id DESC LIMIT 1 ")

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&id)
	}

	fmt.Println(id)

	query, values = sqlbuilder.InsertInto("public.product_specification_languages").
		Cols("language_id", "product_specification_id", "name").
		Values(1, id, "Runflat").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE products SET completed = false")
	if err != nil {
		return err
	}

	return nil
}

func downAddRunflat(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
