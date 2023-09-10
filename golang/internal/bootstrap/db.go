package bootstrap

import (
	"pillowww/titw/internal/db"
	"pillowww/titw/models/hooks"
)

func InitDb() {
	pqConnector := new(db.PostgresAdapter)

	err := db.Initialize(pqConnector)
	if err != nil {
		panic("error initializing Db with postgres connector")
	}

	hooks.RegisterHooks()
}
