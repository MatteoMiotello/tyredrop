package repositories

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type LanguageRepo DbRepo

func NewLanguageRepoFromCtx(ctx context.Context) *LanguageRepo {
	return &LanguageRepo{context: ctx}
}

func (l LanguageRepo) FindOneFromIsoCode(isoCode string) (*models.Language, error) {
	return models.Languages(qm.Where(models.LanguageColumns.IsoCode+"= ?", isoCode)).One(l.context, db.DB)
}
