package supplier

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type dao struct {
	db.Dao
}

func NewDao(executor boil.ContextExecutor) *dao {
	return &dao{
		Dao: db.DaoFromExecutor(executor),
	}
}

func (r dao) GetAll(ctx context.Context) (models.SupplierSlice, error) {
	return models.Suppliers().All(ctx, r.GetConnection(ctx))
}

func (r dao) GetLastImported(ctx context.Context) (*models.Supplier, error) {
	return models.Suppliers(qm.OrderBy(models.SupplierColumns.ImportedAt+" ASC NULLS FIRST")).One(ctx, r.GetConnection(ctx))
}

func (r dao) ExistsJobForFilename(ctx context.Context, supplier models.Supplier, fileName string) (bool, error) {
	return models.ImportJobs(
		models.ImportJobWhere.SupplierID.EQ(supplier.ID),
		models.ImportJobWhere.Filename.EQ(fileName),
	).Exists(ctx, r.GetConnection(ctx))
}

func (r dao) ExistRunningJob(ctx context.Context) (bool, error) {
	return models.ImportJobs(
		qm.Where(models.ImportJobColumns.StartedAt+" IS NOT NULL"),
		qm.And(models.ImportJobColumns.EndedAt+" IS NULL"),
	).Exists(ctx, r.Db)
}
