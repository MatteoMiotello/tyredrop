package vehicle

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
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

func (d Dao) Paginate(limit int, offset int) *Dao {
	return db.Paginate(d, limit, offset)
}

func (d Dao) ForUpdate() *Dao {
	return db.ForUpdate(d)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d Dao) FindById(ctx context.Context, id int64) (*models.VehicleType, error) {
	return models.VehicleTypes(d.GetMods()...).One(ctx, d.Db)
}

func (d Dao) FindByCode(ctx context.Context, code constants.VehicleType) (*models.VehicleType, error) {
	return models.VehicleTypes(
		d.GetMods(
			models.VehicleTypeWhere.Code.EQ(string(code)),
		)...,
	).One(ctx, d.Db)
}
