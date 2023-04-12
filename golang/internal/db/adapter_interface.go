package db

import (
	"github.com/lib/pq"
)

type dbAdapter interface {
	GetConnector() (*pq.Connector, error)
}
