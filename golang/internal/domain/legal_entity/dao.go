package legal_entity

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

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(first int, offset int) *Dao {
	return db.Paginate(d, first, offset)
}

func (d Dao) GetAllTypes(ctx context.Context) (models.LegalEntityTypeSlice, error) {
	return models.LegalEntityTypes(d.GetMods()...).All(ctx, d.Db)
}

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.LegalEntityType, error) {
	return models.LegalEntityTypes(
		d.GetMods(
			models.LegalEntityTypeWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}
