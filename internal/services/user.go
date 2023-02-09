package services

import (
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/controllers"
	"pillowww/titw/internal/repositories"
	"pillowww/titw/models"
	"pillowww/titw/pkg/security"
)

type UserService struct {
	User models.User
}

func (s *UserService) CreateUserFromPayload(payload controllers.SignUpPayload) (*models.User, error) {
	userRoleRepo := new(repositories.UserRoleRepo)
	userRepo := new(repositories.UserRepo)
	adminRole, err := userRoleRepo.FindByRoleCode(repositories.USER_ROLE)

	if err != nil {
		return nil, err
	}

	password, err := security.HashPassword(payload.Password)

	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Email:      payload.Email,
		Username:   null.StringFrom(payload.Username),
		Password:   string(password),
		Name:       payload.Name,
		Surname:    payload.Surname,
		UserRoleID: adminRole.ID,
	}

	err = userRepo.Insert(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
