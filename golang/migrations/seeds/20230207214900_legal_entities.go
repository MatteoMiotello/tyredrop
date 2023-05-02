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
		Cols("name", "is_person").
		Values("Persona fisica", true).
		Values("Ditta individuale", true).
		Values("Società a responsabilità limitata (S.r.l.)", false).
		Values("Società a responsabilità illimitata (S.p.A.)", false).
		Values("Società per azioni semplificata (S.p.A.S.)", false).
		Values("Società in accomandita semplice (S.a.s.)", false).
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
