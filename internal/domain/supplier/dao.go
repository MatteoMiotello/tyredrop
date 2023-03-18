package supplier

import (
	"context"
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
		Dao: db.DaoFromExecutor(executor),
	}
}

func (r Dao) GetAll(ctx context.Context) (models.SupplierSlice, error) {
	return models.Suppliers().All(ctx, r.GetConnection(ctx))
}

func (r Dao) GetLastImported(ctx context.Context) (*models.Supplier, error) {
	return models.Suppliers(qm.OrderBy(models.SupplierColumns.ImportedAt+" ASC NULLS FIRST")).One(ctx, r.GetConnection(ctx))
}

func (r Dao) ExistsJobForFilename(ctx context.Context, supplier models.Supplier, fileName string) (bool, error) {
	return models.ImportJobs(
		models.ImportJobWhere.SupplierID.EQ(supplier.ID),
		models.ImportJobWhere.Filename.EQ(fileName),
	).Exists(ctx, r.GetConnection(ctx))
}

func (r Dao) ExistRunningJob(ctx context.Context) (bool, error) {
	return models.ImportJobs(
		qm.Where(models.ImportJobColumns.StartedAt+" IS NOT NULL"),
		qm.And(models.ImportJobColumns.EndedAt+" IS NULL"),
	).Exists(ctx, r.Db)
}

func (r Dao) FindOneById(ctx context.Context, id int64) (*models.Supplier, error) {
	return models.Suppliers(
		models.SupplierWhere.ID.EQ(id),
	).One(ctx, db.DB)
}
