package db

import (
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DaoMod interface {
	Clone() DaoMod
	AddMods(mods ...qm.QueryMod)
}

func newInstanceWithMods[T DaoMod](d T, mods ...qm.QueryMod) *T {
	no := d.Clone()
	no.AddMods(mods...)
	n := no.(T)

	return &n
}

func Load[T DaoMod](d T, relation string, mods ...qm.QueryMod) *T {
	return newInstanceWithMods(
		d,
		qm.Load(relation, mods...),
	)
}

func Paginate[T DaoMod](d T, limit int, offset int) *T {
	return newInstanceWithMods(
		d,
		qm.Limit(limit),
		qm.Offset(offset),
	)
}

func ForUpdate[T DaoMod](d T) *T {
	return newInstanceWithMods(
		d,
		qm.For("UPDATE"),
	)
}
