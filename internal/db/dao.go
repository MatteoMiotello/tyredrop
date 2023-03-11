package db

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Inserter interface {
	Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

type Upserter interface {
	Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error
}

type Updater interface {
	Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
}

type Dao struct {
	Db boil.ContextExecutor
}

func DaoFromExecutor(executor boil.ContextExecutor) Dao {
	return Dao{
		Db: executor,
	}
}

func (r Dao) GetConnection(ctx context.Context) boil.ContextExecutor {
	return DB
}

func (d Dao) Upsert(ctx context.Context, model Upserter, updateOnConflict bool, cols []string) error {
	return model.Upsert(ctx, d.Db, updateOnConflict, cols, boil.Infer(), boil.Infer())
}

func (d Dao) Update(ctx context.Context, model Updater) error {
	_, err := model.Update(ctx, d.Db, boil.Infer())

	return err
}

func (d Dao) Insert(ctx context.Context, product Inserter) error {
	return product.Insert(ctx, d.Db, boil.Infer())
}
