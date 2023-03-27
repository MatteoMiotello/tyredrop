package brand

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

func (d Dao) SetDao(dao *db.Dao) db.DaoMod {
	d.Dao = dao
	return d
}

func (d Dao) GetDao() *db.Dao {
	return d.Dao
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(first int, offset int) *Dao {
	return db.Paginate(d, first, offset)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d *Dao) FindOneById(ctx context.Context, id int64) (*models.Brand, error) {
	return models.Brands(
		d.GetMods(
			models.BrandWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d *Dao) FindOneByCode(ctx context.Context, code string) (*models.Brand, error) {
	return models.Brands(
		d.GetMods(
			models.BrandWhere.BrandCode.EQ(code),
		)...,
	).One(ctx, d.Db)
}
