package jwt

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/language"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
	"time"
)

type UserStatus int

const (
	USER_COMPLETED   UserStatus = iota
	USER_REGISTERING UserStatus = iota
)

type RoleClaims struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
type UserJwtClaims struct {
	jwt.RegisteredClaims
	Email    string      `json:"email"`
	UserID   int64       `json:"userID"`
	Name     null.String `json:"name"`
	Username null.String `json:"username"`
	Language string      `json:"language"`
	Role     RoleClaims  `json:"role"`
	Status   UserStatus  `json:"status"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	uuid uuid.UUID
}

func CreateAccessTokenFromUser(ctx context.Context, userModel models.User) (string, error) {
	expirationMin := viper.GetInt64("security.jwt.access_token_expiration")
	expiration := jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMin) * time.Minute))
	uRepo := user.NewDao(db.DB)
	uRole, err := uRepo.GetUserRole(ctx, &userModel)

	if err != nil {
		return "", err
	}

	uLanguage, err := uRepo.GetDefaultLanguage(ctx, userModel)

	if err != nil {
		return "", err
	}

	uLanguage = language.FallbackLanguage().L

	rLang, err := user.NewDao(db.DB).GetUserRoleLanguage(ctx, uRole, *uLanguage)

	if err != nil {
		return "", err
	}

	status := USER_COMPLETED

	uBilling, _ := uRepo.GetUserBilling(ctx, &userModel)

	if uBilling == nil {
		status = USER_REGISTERING
	}

	userClaims := UserJwtClaims{
		UserID:   userModel.ID,
		Email:    userModel.Email,
		Username: userModel.Username,
		Language: uLanguage.IsoCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiration,
			Issuer:    viper.GetString("security.jwt.issuer"),
		},
		Role: RoleClaims{
			Name: rLang.Name,
			Code: uRole.RoleCode,
		},
		Status: status,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signedString, err := t.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func CreateUniqueRefreshToken() (string, error) {
	expirationMin := viper.GetInt64("security.jwt.refresh_token_expiration")
	expiration := jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMin) * time.Minute))

	claims := RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiration,
			Issuer:    viper.GetString("security.jwt.issuer"),
		},
		uuid: uuid.New(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := t.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ParseUserJwt(token string) (*UserJwtClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	return t.Claims.(*UserJwtClaims), err
}
