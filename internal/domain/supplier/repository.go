package supplier

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type repository db.Repo

func NewRepository() *repository {
	return &repository{
		Db: db.DB,
	}
}

func (r repository) GetAll(ctx context.Context) (models.SupplierSlice, error) {
	return models.Suppliers().All(ctx, r.Db)
}

func (r repository) GetLastImported(ctx context.Context) (*models.Supplier, error) {
	return models.Suppliers(qm.OrderBy(models.SupplierColumns.ImportedAt+" ASC")).One(ctx, r.Db)
}

func (r repository) Update(ctx context.Context, supplier *models.Supplier) error {
	_, err := supplier.Update(ctx, r.Db, boil.Infer())

	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetJobsFromFilename(ctx context.Context, supplier models.Supplier, fileName string) (models.ImportJobSlice, error) {
	return models.ImportJobs(
		models.ImportJobWhere.SupplierID.EQ(supplier.ID),
		models.ImportJobWhere.Filename.EQ(fileName),
		models.ImportJobWhere.EndedAt.IsNotNull(),
		models.ImportJobWhere.StartedAt.IsNotNull(),
	).All(ctx, r.Db)
}
