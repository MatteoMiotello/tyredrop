package services

import (
	"context"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/language"
	"pillowww/titw/internal/repositories"
	"pillowww/titw/models"
	"pillowww/titw/pkg/security"
)

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}

type CreateUserPayload struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func (s *UserService) CreateUserFromPayload(ctx context.Context, payload CreateUserPayload) (*models.User, error) {
	userRoleRepo := repositories.NewUserRoleRepoFromCtx(ctx)
	userRepo := repositories.NewUserRepoWithCtx(ctx)
	defLanguage := language.FromContext(ctx).Language
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
