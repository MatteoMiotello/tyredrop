package db

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

type DaoMod interface {
	addMod(mod qm.QueryMod)
}

func Load(dao DaoMod, relation string, mods ...qm.QueryMod) DaoMod {
	dao.addMod(qm.Load(relation, mods...))

	return dao
}
