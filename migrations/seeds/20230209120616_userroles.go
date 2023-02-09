package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upUserroles, downUserroles)
}

func upUserroles(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.user_roles").
		Cols("role_code", "admin").
		Values("ADMIN", true).
		Values("SUPPLIER", false).
		Values("USER", false).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}

	var languageId int64

	row := tx.QueryRow("SELECT id FROM languages LIMIT 1")
	err = row.Scan(&languageId)
	if err != nil {
		return err
	}

	var roleId int64

	row = tx.QueryRow("SELECT id FROM user_roles WHERE role_code = 'ADMIN'")

	err = row.Scan(&roleId)
	if err != nil {
		return err
	}

	query, values = sqlbuilder.InsertInto("user_role_languages").
		Cols("language_id", "user_role_id", "name").
		Values(languageId, roleId, "Amministratore").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	row = tx.QueryRow("SELECT id FROM user_roles WHERE role_code = 'SUPPLIER'")

	err = row.Scan(&roleId)
	if err != nil {
		return err
	}

	query, values = sqlbuilder.InsertInto("user_role_languages").
		Cols("language_id", "user_role_id", "name").
		Values(languageId, roleId, "Fornitore").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	row = tx.QueryRow("SELECT id FROM user_roles WHERE role_code = 'USER'")

	err = row.Scan(&roleId)
	if err != nil {
		return err
	}

	query, values = sqlbuilder.InsertInto("user_role_languages").
		Cols("language_id", "user_role_id", "name").
		Values(languageId, roleId, "Utente").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func downUserroles(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
