package user

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
	"time"
)

type RPDao struct {
	*db.Dao
}

func NewResetPasswordDao(executor boil.ContextExecutor) *RPDao {
	return &RPDao{
		db.DaoFromExecutor(executor),
	}
}

func (u RPDao) Load(relationship string, mods ...qm.QueryMod) *RPDao {
	return db.Load(u, relationship, mods...)
}

func (u RPDao) Paginate(limit int, offset int) *RPDao {
	return db.Paginate(u, limit, offset)
}

func (u RPDao) ForUpdate() *RPDao {
	return db.ForUpdate(u)
}

func (u RPDao) Clone() db.DaoMod {
	return RPDao{
		u.Dao.Clone(),
	}
}

func (u RPDao) FindValidByToken(ctx context.Context, token string) (*models.ResetPassword, error) {
	return models.ResetPasswords(
		u.GetMods(
			models.ResetPasswordWhere.Token.EQ(token),
			models.ResetPasswordWhere.ExpiryAt.GT(time.Now()),
		)...,
	).One(ctx, db.DB)
}
