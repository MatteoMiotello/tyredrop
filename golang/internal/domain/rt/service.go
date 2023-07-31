package rt

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"pillowww/titw/models"
	"time"
)

type RefreshTokenService struct {
	RTDao *Dao
}

func NewRefreshTokenService(dao *Dao) *RefreshTokenService {
	return &RefreshTokenService{
		RTDao: dao,
	}
}

func (r RefreshTokenService) StoreNew(ctx context.Context, user models.User, refreshToken string) error {
	olds, _ := r.RTDao.FindAllByUser(ctx, user)

	if len(olds) > 0 {
		for _, old := range olds {
			err := r.RTDao.Delete(ctx, old)
			if err != nil {
				return err
			}
		}
	}

	expirationMin := viper.GetInt("security.refresh_token.expiration")

	newRt := &models.RefreshToken{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(expirationMin) * time.Minute),
	}

	err := r.RTDao.Insert(ctx, newRt)
	if err != nil {
		return err
	}

	return nil
}
