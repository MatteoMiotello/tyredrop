package import_job

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (r Dao) Insert(ctx context.Context, job *models.ImportJob) error {
	return job.Insert(ctx, r.GetConnection(ctx), boil.Infer())
}

func (r Dao) Update(ctx context.Context, job *models.ImportJob) error {
	_, err := job.Update(ctx, r.GetConnection(ctx), boil.Infer())

	if err != nil {
		return err
	}

	return nil
}
