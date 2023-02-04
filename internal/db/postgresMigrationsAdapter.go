package db

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type PostgresMigrationAdapter struct {
}

func (p PostgresMigrationAdapter) GetConnector() (*pq.Connector, error) {
	dsn := fmt.Sprintf("postgres://%s@%s/%s?port=%s&password=%s",
		viper.Get("MIGRATION_DB_USER"),
		viper.Get("MIGRATION_DB_HOST"),
		viper.Get("MIGRATION_DB_NAME"),
		viper.Get("MIGRATION_DB_PORT"),
		viper.Get("MIGRATION_DB_PASSWORD"))

	return pq.NewConnector(dsn)
}
