package supplier

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
		Dao: db.DaoFromExecutor(executor),
	}
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(first int, offset int) *Dao {
	return db.Paginate(d, first, offset)
}

func (d Dao) GetAll(ctx context.Context) (models.SupplierSlice, error) {
	return models.Suppliers(d.GetMods()...).All(ctx, d.Db)
}

func (d Dao) GetLastImported(ctx context.Context) (*models.Supplier, error) {
	return models.Suppliers(
		d.GetMods(
			qm.OrderBy(models.SupplierColumns.ImportedAt+" ASC NULLS FIRST"),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) ExistsJobForFilename(ctx context.Context, supplier models.Supplier, fileName string) (bool, error) {
	return models.ImportJobs(
		d.GetMods(
			models.ImportJobWhere.SupplierID.EQ(supplier.ID),
			models.ImportJobWhere.Filename.EQ(fileName),
		)...,
	).Exists(ctx, d.Db)
}

func (d Dao) ExistRunningJob(ctx context.Context) (bool, error) {
	return models.ImportJobs(
		d.GetMods(
			qm.Where(models.ImportJobColumns.StartedAt+" IS NOT NULL"),
			qm.And(models.ImportJobColumns.EndedAt+" IS NULL"),
		)...,
	).Exists(ctx, d.Db)
}

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Supplier, error) {
	return models.Suppliers(
		d.GetMods(
			models.SupplierWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}
