package language

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/net/context"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type repo db.Repo

func NewLanguageRepo(db *sql.DB) *repo {
	return &repo{
		Db: db,
	}
}

func (l repo) FindOneFromIsoCode(ctx context.Context, isoCode string) (*models.Language, error) {
	return models.Languages(qm.Where(models.LanguageColumns.IsoCode+"= ?", isoCode)).One(ctx, l.Db)
}
