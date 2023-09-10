package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"pillowww/titw/pkg/security"
)

func init() {
	goose.AddMigration(upAdminUser, downAdminUser)
}

func upAdminUser(tx *sql.Tx) error {
	defPass := viper.GetString("DEFAULT_PASSWORD")
	hashed, err := security.HashPassword(defPass)

	if err != nil {
		return err
	}

	query, vals := sqlbuilder.InsertInto("public.users").
		Cols("user_role_id", "default_language_id", "email", "username", "password", "name", "surname", "confirmed").
		Values(1, 1, viper.GetString("DEFAULT_EMAIL"), "administrator", string(hashed), "Admin", nil, true).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = tx.Exec(query, vals...)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func downAdminUser(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
