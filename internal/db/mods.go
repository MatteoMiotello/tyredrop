package db

import (
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DaoMod interface {
	SetDao(dao *Dao)
	GetDao() *Dao
}

type DaoModded[T DaoMod] struct {
	dao T
}

func Load(d DaoMod, relation string, mods ...qm.QueryMod) {
	n := &Dao{
		Db:   d.GetDao().Db,
		Mods: d.GetDao().Mods,
	}

	n.addMods(qm.Load(relation, mods...))
	d.SetDao(n)
}

func Paginate(d DaoMod, first int, offset int) {
	n := &Dao{
		Db:   d.GetDao().Db,
		Mods: d.GetDao().Mods,
	}

	n.addMods(
		qm.Limit(first),
		qm.Offset(offset),
	)
	d.SetDao(n)
}
