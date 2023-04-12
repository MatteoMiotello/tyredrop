package user

import (
	"context"
	"github.com/volatiletech/null/v8"
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

type Service struct {
	UDao *Dao
}

func NewUserService(dao *Dao) *Service {
	return &Service{
		UDao: dao,
	}
}

func (s Service) CreateUserFromPayload(ctx context.Context, payload CreateUserPayload) (*models.User, error) {
	lRepo := language.NewDao(s.UDao.Db)
	lModel, err := lRepo.FindOneFromIsoCode(ctx, payload.LanguageCode)

	if err != nil {
		lModel = language.FallbackLanguage().L
	}

	adminRole, err := s.UDao.FindUserRoleByCode(ctx, USER_ROLE)

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

	err = s.UDao.Insert(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
