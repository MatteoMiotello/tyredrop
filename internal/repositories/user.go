package repositories

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type UserRepo DbRepo

func NewUserRepoWithCtx(ctx context.Context) *UserRepo {
	return &UserRepo{context: ctx}
}

func (u UserRepo) FindOneByUsername(username string) (*models.User, error) {
	return models.Users(qm.Where("username = ?", username)).One(u.context, db.DB)
}

func (u UserRepo) FindOneByEmail(email string) (*models.User, error) {
	return models.Users(qm.Where("email = ?", email)).One(u.context, db.DB)
}

func (u UserRepo) FindOneById(id int64) (*models.User, error) {
	return models.FindUser(u.context, db.DB, id)
}

func (u UserRepo) Insert(user *models.User) error {
	return user.Insert(u.context, db.DB, boil.Infer())
}
