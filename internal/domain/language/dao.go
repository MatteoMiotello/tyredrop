package language

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/net/context"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *dao {
	return &dao{
		db.DaoFromExecutor(executor),
	}
}

func (l dao) FindOneFromIsoCode(ctx context.Context, isoCode string) (*models.Language, error) {
	return models.Languages(qm.Where(models.LanguageColumns.IsoCode+"= ?", isoCode)).One(ctx, l.Db)
}
