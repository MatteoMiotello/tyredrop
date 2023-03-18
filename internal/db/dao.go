package db

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

type Deleter interface {
	Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error)
}

type Dao struct {
	Db   boil.ContextExecutor
	Mods []qm.QueryMod
}

func DaoFromExecutor(executor boil.ContextExecutor) *Dao {
	return &Dao{
		Db: executor,
	}
}

func (d *Dao) addMods(mods ...qm.QueryMod) {
	d.Mods = append(d.Mods, mods...)
}

func (d *Dao) GetMods(mods ...qm.QueryMod) []qm.QueryMod {
	defer func() {
		d.Mods = nil
	}()

	return append(mods, d.Mods...)
}

func (d *Dao) GetConnection(ctx context.Context) boil.ContextExecutor {
	return DB
}

func (d *Dao) Upsert(ctx context.Context, model Upserter, updateOnConflict bool, cols []string) error {
	return model.Upsert(ctx, d.Db, updateOnConflict, cols, boil.Infer(), boil.Infer())
}

func (d *Dao) Update(ctx context.Context, model Updater) error {
	_, err := model.Update(ctx, d.Db, boil.Infer())

	return err
}

func (d *Dao) Insert(ctx context.Context, model Inserter) error {
	return model.Insert(ctx, d.Db, boil.Infer())
}

func (d *Dao) Delete(ctx context.Context, model Deleter) error {
	_, err := model.Delete(ctx, d.Db, false)

	return err
}
