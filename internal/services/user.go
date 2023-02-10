package services

import (
	"context"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/repositories"
	"pillowww/titw/models"
	"pillowww/titw/pkg/security"
	"time"
)

type UserService struct {
	User *models.User
}

func NewUserService(user *models.User) *UserService {
	return &UserService{
		User: user,
	}
}

type CreateUserPayload struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func CreateUserFromPayload(ctx context.Context, payload CreateUserPayload) (*models.User, error) {
	userRoleRepo := repositories.NewUserRoleRepoFromCtx(ctx)
	userRepo := repositories.NewUserRepoWithCtx(ctx)
	defLanguage := AuthServiceFromCtx(ctx).Language.L
	adminRole, err := userRoleRepo.FindByRoleCode(repositories.USER_ROLE)

	if err != nil {
		return nil, err
	}

	password, err := security.HashPassword(payload.Password)

	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Email:             payload.Email,
		Username:          null.StringFrom(payload.Username),
		Password:          string(password),
		Name:              payload.Name,
		Surname:           payload.Surname,
		UserRoleID:        adminRole.ID,
		DefaultLanguageID: defLanguage.ID,
	}

	err = userRepo.Insert(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (s UserService) StoreNewRefreshToken(ctx context.Context, refreshToken string) error {
	rtRepo := repositories.NewRefreshTokenRepoFromCtx(ctx)

	olds, _ := rtRepo.FindAllByUser(*s.User)

	if len(olds) > 0 {
		for _, old := range olds {
			err := rtRepo.Delete(old)
			if err != nil {
				return err
			}
		}
	}

	expirationMin := viper.GetInt("security.refresh_token.expiration")

	newRt := &models.RefreshToken{
		UserID:       s.User.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(expirationMin) * time.Minute),
	}

	err := rtRepo.Insert(newRt)
	if err != nil {
		return err
	}

	return nil
}
