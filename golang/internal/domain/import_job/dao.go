package import_job

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"pillowww/titw/internal/db"
)

type Dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		Dao: db.DaoFromExecutor(executor),
	}
}
