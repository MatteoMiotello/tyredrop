package repositories

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type RefreshTokenRepo DbRepo

func NewRefreshTokenRepoFromCtx(ctx context.Context) *RefreshTokenRepo {
	return &RefreshTokenRepo{
		context: ctx,
	}
}

func (r RefreshTokenRepo) FindAllByUser(user models.User) (models.RefreshTokenSlice, error) {
	return models.RefreshTokens(models.RefreshTokenWhere.UserID.EQ(user.ID)).All(r.context, db.DB)
}

func (r RefreshTokenRepo) Delete(refreshToken *models.RefreshToken) error {
	_, err := refreshToken.Delete(r.context, db.DB, false)

	if err != nil {
		return err
	}

	return nil
}

func (r RefreshTokenRepo) Insert(token *models.RefreshToken) error {
	err := token.Insert(r.context, db.DB, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
