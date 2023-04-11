package language

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/net/context"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *dao {
	return &dao{
		db.DaoFromExecutor(executor),
	}
}

func (d dao) Clone() db.DaoMod {
	return dao{
		d.Dao.Clone(),
	}
}

func (d dao) Load(relationship string, mods ...qm.QueryMod) *dao {
	return db.Load(d, relationship, mods...)
}

func (d dao) Paginate(first int, offset int) *dao {
	return db.Paginate(d, first, offset)
}

func (d *dao) FindOneFromIsoCode(ctx context.Context, isoCode string) (*models.Language, error) {
	return models.Languages(
		d.GetMods(
			qm.Where(models.LanguageColumns.IsoCode+"= ?", isoCode),
		)...,
	).One(ctx, d.Db)
}
