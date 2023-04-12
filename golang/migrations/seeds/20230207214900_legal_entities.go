package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upLegalEntities, downLegalEntities)
}

func upLegalEntities(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.legal_entity_types").
		Cols("name").
		Values("Persona fisica").
		Values("Ditta individuale").
		Values("Società a responsabilità limitata (S.r.l.)").
		Values("Società a responsabilità illimitata (S.p.A.)").
		Values("Società per azioni semplificata (S.p.A.S.)").
		Values("Società in accomandita semplice (S.a.s.)").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func downLegalEntities(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
