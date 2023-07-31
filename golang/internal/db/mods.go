package db

import (
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DaoMod interface {
	Clone() DaoMod
	AddMods(mods ...qm.QueryMod)
}

type OrderMods struct {
	Column string
	Desc   *bool
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

func Order[T DaoMod](d T, ordersMods []OrderMods) *T {
	var mods []qm.QueryMod

	for _, order := range ordersMods {
		direction := " ASC"

		if order.Desc != nil && *order.Desc {
			direction = " DESC"
		}

		mods = append(mods, qm.OrderBy(order.Column+direction))
	}

	return newInstanceWithMods(
		d,
		mods...,
	)
}

func WithDeletes[T DaoMod](d T) *T {
	return newInstanceWithMods(
		d,
		qm.WithDeleted(),
	)
}

func ForUpdate[T DaoMod](d T) *T {
	return newInstanceWithMods(
		d,
		qm.For("NO KEY UPDATE"),
	)
}
