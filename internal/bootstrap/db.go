package bootstrap

import "pillowww/Totw/internal/db"

func initDb() {
	pqConnector := new(db.PostgresAdapter)

	err := db.Initialize(pqConnector)
	if err != nil {
		panic("error initializing Db with postgres connector")
	}
}
