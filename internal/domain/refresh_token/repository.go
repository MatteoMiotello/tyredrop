package refresh_token

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/net/context"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
	"time"
)

type RefreshTokenRepo db.Repo

func NewRefreshTokenRepo(db *sql.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{
		Db: db,
	}
}

func (r RefreshTokenRepo) FindAllByUser(ctx context.Context, user models.User) (models.RefreshTokenSlice, error) {
	return models.RefreshTokens(models.RefreshTokenWhere.UserID.EQ(user.ID)).All(ctx, r.Db)
}

func (r RefreshTokenRepo) Delete(ctx context.Context, refreshToken *models.RefreshToken) error {
	_, err := refreshToken.Delete(ctx, r.Db, false)

	if err != nil {
		return err
	}

	return nil
}

func (r RefreshTokenRepo) Insert(ctx context.Context, token *models.RefreshToken) error {
	err := token.Insert(ctx, r.Db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r RefreshTokenRepo) FindValidOneFromRefreshToken(ctx context.Context, refreshToken string) (*models.RefreshToken, error) {
	return models.RefreshTokens(
		models.RefreshTokenWhere.RefreshToken.EQ(refreshToken),
		models.RefreshTokenWhere.ExpiresAt.GTE(time.Now()),
	).One(ctx, r.Db)
}

func (r RefreshTokenRepo) GetUser(ctx context.Context, token models.RefreshToken) (*models.User, error) {
	return token.User().One(ctx, r.Db)
}

func (r RefreshTokenRepo) Update(ctx context.Context, token *models.RefreshToken) error {
	_, err := token.Update(ctx, r.Db, boil.Infer())
	return err
}
