package db

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Dao struct {
	Db boil.ContextExecutor
}

func DaoFromExecutor(executor boil.ContextExecutor) Dao {
	return Dao{
		Db: executor,
	}
}

func (r Dao) GetConnection(ctx context.Context) boil.ContextExecutor {
	return DB
}
