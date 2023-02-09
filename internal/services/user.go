package services

import (
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/repositories"
	"pillowww/titw/models"
	"pillowww/titw/pkg/security"
)

type UserService struct {
	User models.User
}

type CreateUserPayload struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func (s *UserService) CreateUserFromPayload(payload CreateUserPayload) (*models.User, error) {
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

	s.User = newUser

	return &newUser, nil
}
