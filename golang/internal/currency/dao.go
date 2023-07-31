package currency

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type Dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		db.DaoFromExecutor(executor),
	}
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(first int, offset int) *Dao {
	return db.Paginate(d, first, offset)
}

func (d Dao) ForUpdate() *Dao {
	return db.ForUpdate(d)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d Dao) FindDefault(ctx context.Context) (*models.Currency, error) {
	return models.Currencies(
		d.GetMods()...,
	).One(ctx, d.Db)
}

func (d Dao) FindById(ctx context.Context, id int64) (*models.Currency, error) {
	return models.Currencies(
		d.GetMods(
			models.CurrencyWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindAll(ctx context.Context) (models.CurrencySlice, error) {
	return models.Currencies(
		d.GetMods()...,
	).All(ctx, d.Db)
}
