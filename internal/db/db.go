package db

import (
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
