package db

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type PostgresAdapter struct {
}

func (p PostgresAdapter) GetConnector() (*pq.Connector, error) {
	dsn := fmt.Sprintf("postgres://%s@%s/%s?port=%s&password=%s&sslmode=disable&lock_timeout=%dms",
		viper.Get("DB_USER"),
		viper.Get("DB_HOST"),
		viper.Get("DB_NAME"),
		viper.Get("DB_PORT"),
		viper.Get("DB_PASSWORD"),
		1000,
	)

	return pq.NewConnector(dsn)
}
