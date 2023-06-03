package cart

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

func (d Dao) FindAllByUserId(ctx context.Context, userId int64) (models.CartSlice, error) {
	return models.Carts(
		d.GetMods(
			models.CartWhere.UserID.EQ(userId),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindOneByUserAndProductItemId(ctx context.Context, userId int64, itemId int64) (*models.Cart, error) {
	return models.Carts(
		d.GetMods(
			models.CartWhere.UserID.EQ(userId),
			models.CartWhere.ProductItemID.EQ(itemId),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindOneById(ctx context.Context, cartId int64) (*models.Cart, error) {
	return models.Carts(
		d.GetMods(
			models.CartWhere.ID.EQ(cartId),
		)...,
	).One(ctx, d.Db)
}
