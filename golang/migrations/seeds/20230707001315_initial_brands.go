package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upInitialBrands, downInitialBrands)
}

func upInitialBrands(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("brands").
		Cols("brand_code", "name", "quality").
		Values("BRIDGESTONE", "Bridgestone", 5).
		Values("CONTINENTAL", "Continental", 5).
		Values("DUNLOP", "Dunlop", 5).
		Values("GOODYEAR", "Goodyear", 5).
		Values("HANKOOK", "Hankook", 5).
		Values("MICHELIN", "Michelin", 5).
		Values("PIRELLI", "Pirelli", 5).
		Values("AVON", "AVON", 4).
		Values("BFGOODRICH", "BFGoodrich", 4).
		Values("COOPER", "Cooper", 4).
		Values("FALKEN", "Falken", 4).
		Values("FEDERAL", "Federal", 4).
		Values("FIRESTONE", "Firestone", 4).
		Values("FULDA", "Fulda", 4).
		Values("GENERAL TIRE", "General Tire", 4).
		Values("GT-RADIAL", "GT-Radial", 4).
		Values("KLEBER", "Kleber", 4).
		Values("KUMHO", "Kumho", 4).
		Values("LAUFENN", "Laufenn", 4).
		Values("NOKIAN", "Nokian", 4).
		Values("TOYO", "Toyo", 4).
		Values("UNIROYAL", "Uniroyal", 4).
		Values("VREDESTEIN", "Vredestein", 4).
		Values("YOKOHAMA", "Yokohama", 4).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func downInitialBrands(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
