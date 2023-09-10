package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddAvatarToUser, downAddAvatarToUser)
}

func upAddAvatarToUser(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users ADD COLUMN avatar_path varchar NULL")
	if err != nil {
		return err
	}

	return nil
}

func downAddAvatarToUser(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users DROP COLUMN avatar_path")
	if err != nil {
		return err
	}
	return nil
}
