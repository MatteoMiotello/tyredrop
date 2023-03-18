package db

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

type DaoMod interface {
	addMods(mod ...qm.QueryMod)
}

func Load(dao DaoMod, relation string, mods ...qm.QueryMod) DaoMod {
	dao.addMods(qm.Load(relation, mods...))

	return dao
}

func Paginate(dao DaoMod, first int, offset int) {
	dao.addMods(
		qm.Limit(first),
		qm.Offset(offset),
	)
}
