package import_job

import (
	"context"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/models"
	"time"
)

type ImportJobService struct {
	IJDao *Dao
}

func NewImportJobService(ijDao *Dao) *ImportJobService {
	return &ImportJobService{
		IJDao: ijDao,
	}
}

func (i ImportJobService) CreateJob(ctx context.Context, sup models.Supplier, fileName string) (*models.ImportJob, error) {
	job := &models.ImportJob{
		SupplierID: sup.ID,
		Filename:   fileName,
	}

	err := i.IJDao.Insert(ctx, job)

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (i ImportJobService) StartNow(ctx context.Context, job *models.ImportJob) error {
	job.StartedAt = null.TimeFrom(time.Now())

	return i.IJDao.Update(ctx, job)
}

func (i ImportJobService) EndNow(ctx context.Context, job *models.ImportJob) error {
	job.EndedAt = null.TimeFrom(time.Now())

	return i.IJDao.Update(ctx, job)
}

func (i ImportJobService) EndNowWithError(ctx context.Context, job *models.ImportJob, err string) error {
	job.ErrorMessage = null.StringFrom(err)

	return i.EndNow(ctx, job)
}
