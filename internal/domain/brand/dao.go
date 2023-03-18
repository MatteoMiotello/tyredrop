package brand

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

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Brand, error) {
	return models.FindBrand(ctx, d.Db, id)
}

func (d Dao) FindOneByCode(ctx context.Context, code string) (*models.Brand, error) {
	return models.Brands(models.BrandWhere.BrandCode.EQ(code)).One(ctx, d.Db)
}
