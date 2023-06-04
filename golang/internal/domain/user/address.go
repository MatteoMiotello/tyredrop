package user

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type AddressDao struct {
	*db.Dao
}

func NewAddressDao(executor boil.ContextExecutor) *AddressDao {
	return &AddressDao{
		db.DaoFromExecutor(executor),
	}
}

func (d AddressDao) Load(relationship string, mods ...qm.QueryMod) *AddressDao {
	return db.Load(d, relationship, mods...)
}

func (d AddressDao) Paginate(limit int, offset int) *AddressDao {
	return db.Paginate(d, limit, offset)
}

func (d AddressDao) ForUpdate() *AddressDao {
	return db.ForUpdate(d)
}

func (d AddressDao) Clone() db.DaoMod {
	return AddressDao{
		d.Dao.Clone(),
	}
}

func (d AddressDao) FindOneById(ctx context.Context, id int64) (*models.UserAddress, error) {
	return models.UserAddresses(
		d.GetMods(
			models.UserAddressWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d AddressDao) FindAllByUserId(ctx context.Context, userId int64) (models.UserAddressSlice, error) {
	return models.UserAddresses(
		d.GetMods(
			models.UserAddressWhere.UserID.EQ(userId),
		)...,
	).All(ctx, d.Db)
}
