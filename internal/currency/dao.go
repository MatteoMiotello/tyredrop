package currency

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (d Dao) FindDefault(ctx context.Context) (*models.Currency, error) {
	return models.Currencies().One(ctx, d.Db)
}
