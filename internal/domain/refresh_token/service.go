package refresh_token

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
	"time"
)

func StoreNew(ctx context.Context, user models.User, refreshToken string) error {
	rtRepo := NewRefreshTokenRepo(db.DB)

	olds, _ := rtRepo.FindAllByUser(ctx, user)

	if len(olds) > 0 {
		for _, old := range olds {
			err := rtRepo.Delete(ctx, old)
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

	err := rtRepo.Insert(ctx, newRt)
	if err != nil {
		return err
	}

	return nil
}
