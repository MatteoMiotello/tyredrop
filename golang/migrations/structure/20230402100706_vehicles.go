package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upVehicles, downVehicles)
}

func upVehicles(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.vehicle_types").
		PKColumn().
		Column("code", types.Varchar.Options("255"), false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	query = sqlbuilder.CreateTable("public.vehicle_type_languages").
		PKColumn().
		FKColumn("public.vehicle_types", "vehicle_type_id", false).
		FKColumn("public.languages", "language_id", false).
		Column("name", types.Varchar.Options("255"), false).
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err = tx.Exec(query)

	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.products ADD vehicle_type_id int8 NOT NULL;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.products ADD CONSTRAINT vehicle_types_fk FOREIGN KEY (vehicle_type_id) REFERENCES public.vehicle_types(id) ON UPDATE CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX idx_products_vehicle_types ON products (vehicle_type_id)")
	if err != nil {
		return err
	}

	return nil
}

func downVehicles(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.vehicle_type_languages")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE public.vehicle_types")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP INDEX idx_products_vehicle_types;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.products DROP CONSTRAINT vehicle_types_fk;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.products DROP COLUMN vehicle_id;")
	if err != nil {
		return err
	}

	return nil
}
