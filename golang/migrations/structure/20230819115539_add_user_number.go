package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddUserNumber, downAddUserNumber)
}

func upAddUserNumber(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users ADD COLUMN user_code varchar NULL")
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET user_code = ( SELECT concat( 'USR', upper( CAST( to_hex( user1.id + 120000 ) AS varchar( 100 ) ) ) ) FROM users as user1 WHERE user1.id = users.id )")
	if err != nil {
		return err
	}
	return nil
}

func downAddUserNumber(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users DROP COLUMN user_code")

	if err != nil {
		return err
	}
	return nil
}
