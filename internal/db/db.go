package db

import (
	"context"
	"database/sql"
	"github.com/spf13/viper"
	"log"
	"time"
)

var DB *sql.DB

func Initialize(adapter dbAdapter) error {
	connector, err := adapter.GetConnector()

	if err != nil {
		return err
	}

	db := sql.OpenDB(connector)
	db.SetMaxOpenConns(viper.GetInt("db.max-open-connection"))
	db.SetMaxIdleConns(viper.GetInt("db.max-idle-connection"))
	db.SetConnMaxLifetime(time.Duration(viper.GetInt("db.max-connection-lifetime")) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(viper.GetInt("db.max-connection-idle-lifetime")) * time.Minute)

	DB = db
	return nil
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatalf("goose: failed to close DB: %v\n", err)
	}
}

func WithTx(ctx context.Context, handle func(tx *sql.Tx) error) error {
	tx, err := DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})

	if err != nil {
		return err
	}

	defer tx.Rollback()

	handleErr := handle(tx)

	if handleErr != nil {
		if err != nil {
			return err
		}

		return handleErr
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
