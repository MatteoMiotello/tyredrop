package user

import (
	"context"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/language"
	"pillowww/titw/models"
	"pillowww/titw/pkg/security"
)

type CreateUserPayload struct {
	Email        string `json:"email" binding:"required" validate:"email"`
	Username     string `json:"username"`
	Password     string `json:"password" binding:"required"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	LanguageCode string `json:"language_code"`
}

func CreateUserFromPayload(ctx context.Context, payload CreateUserPayload) (*models.User, error) {
	userRepo := NewUserRepo(db.DB)
	lRepo := language.NewLanguageRepo(db.DB)
	lModel, err := lRepo.FindOneFromIsoCode(ctx, payload.LanguageCode)

	if err != nil {
		lModel = language.FallbackLanguage().L
	}

	adminRole, err := userRepo.FindUserRoleByCode(ctx, USER_ROLE)

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
		DefaultLanguageID: lModel.ID,
	}

	err = userRepo.Insert(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
