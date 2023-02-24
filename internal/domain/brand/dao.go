package brand

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type Dao struct {
	db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		db.DaoFromExecutor(executor),
	}
}

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Brand, error) {
	return models.FindBrand(ctx, d.Db, id)
}

func (d Dao) FindOneByName(ctx context.Context, name string) (*models.Brand, error) {
	return models.Brands(qm.Where(models.BrandColumns.Name+" LIKE ? ", fmt.Sprintf("%%%s%%", name))).One(ctx, d.Db)
}

func (d Dao) Update(ctx context.Context, brand *models.Brand) error {
	_, err := brand.Update(ctx, d.Db, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

func (d Dao) Insert(ctx context.Context, brand *models.Brand) error {
	return brand.Insert(ctx, d.Db, boil.Infer())
}
