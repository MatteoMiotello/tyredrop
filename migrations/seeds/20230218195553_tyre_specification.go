package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
	"pillowww/titw/pkg/constants"
)

func init() {
	goose.AddMigration(upTyreSpecification, downTyreSpecification)
}

func upTyreSpecification(tx *sql.Tx) error {
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
		Cols("product_category_id", "specification_code", "type").
		Values(catId, "NAME", constants.SPEC_TYPE_STRING).
		Values(catId, "REFERENCE", constants.SPEC_TYPE_STRING).
		Values(catId, "WIDTH", constants.SPEC_TYPE_INT).
		Values(catId, "ASPECT_RATIO", constants.SPEC_TYPE_INT).
		Values(catId, "CONSTRUCTION", constants.SPEC_TYPE_STRING).
		Values(catId, "RIM", constants.SPEC_TYPE_INT).
		Values(catId, "LOAD", constants.SPEC_TYPE_INT).
		Values(catId, "SPEED", constants.SPEC_TYPE_STRING).
		Values(catId, "SEASON", constants.SPEC_TYPE_STRING).
		Values(catId, "EPREL_ID", constants.SPEC_TYPE_INT).
		Values(catId, "FUEL_EFFICIENCY", constants.SPEC_TYPE_STRING).
		Values(catId, "WET_GRIP_CLASS", constants.SPEC_TYPE_STRING).
		Values(catId, "EXTERNAL_ROLLING_NOISE_CLASS", constants.SPEC_TYPE_STRING).
		Values(catId, "EXTERNAL_ROLLING_NOISE_LEVEL", constants.SPEC_TYPE_INT).
		Values(catId, "LOAD_VERSION", constants.SPEC_TYPE_STRING).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	var ids []int64
	sQuery, _ := sqlbuilder.Select("id").From("public.product_specifications").BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, _ := tx.Query(sQuery)
	defer rows.Close()
	for rows.Next() {
		var id int64
		rows.Scan(&id)

		ids = append(ids, id)
	}
	if err != nil {
		return err
	}

	query, values = sqlbuilder.InsertInto("public.product_specification_languages").
		Cols("language_id", "product_specification_id", "name").
		Values(1, ids[0], "Profilo").
		Values(1, ids[1], "Referenza").
		Values(1, ids[2], "Larghezza").
		Values(1, ids[3], "Altezza").
		Values(1, ids[4], "Costruzione").
		Values(1, ids[5], "Diametro").
		Values(1, ids[6], "Carico").
		Values(1, ids[7], "Velocità").
		Values(1, ids[8], "Stagione").
		Values(1, ids[9], "Id Eprel").
		Values(1, ids[10], "Categoria di consumo carburante").
		Values(1, ids[11], "Categoria di aderenza sul bagnato").
		Values(1, ids[12], "Categoria di rumore esterno di rotolamento").
		Values(1, ids[13], "Livello di rumore esterno di rotolamento").
		Values(1, ids[14], "Versione rispetto al carico").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func downTyreSpecification(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
