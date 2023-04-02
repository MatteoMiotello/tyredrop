package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
	"pillowww/titw/pkg/constants"
)

func init() {
	goose.AddMigration(upVehicleTypes, downVehicleTypes)
}

func upVehicleTypes(tx *sql.Tx) error {
	query, params := sqlbuilder.InsertInto("public.vehicle_types").
		Cols("code").
		Values(constants.VEHICLE_CAR).
		Values(constants.VEHICLE_MOTO).
		Values(constants.VEHICLE_TRUCK).
		Values(constants.VEHICLE_QUAD).
		Values(constants.VEHICLE_AGRICULTURAL).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, params...)
	if err != nil {
		return err
	}

	codes := map[constants.VehicleType]string{
		constants.VEHICLE_CAR:          "Automobili",
		constants.VEHICLE_MOTO:         "Moto",
		constants.VEHICLE_TRUCK:        "Camion e Furgoni",
		constants.VEHICLE_QUAD:         "Quad",
		constants.VEHICLE_AGRICULTURAL: "Agricoli",
	}

	for code, name := range codes {
		querySelect := sqlbuilder.Select("id").
			From("vehicle_types").
			Where("code = '" + string(code) + "'").
			String()

		var id int64

		row := tx.QueryRow(querySelect)
		err := row.Scan(&id)
		if err != nil {
			return err
		}

		query, params = sqlbuilder.InsertInto("public.vehicle_type_languages").
			Cols("language_id", "vehicle_type_id", "name").
			Values(1, id, name).
			BuildWithFlavor(sqlbuilder.PostgreSQL)

		_, err = tx.Exec(query, params...)
		if err != nil {
			return err
		}
	}

	return nil
}

func downVehicleTypes(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
