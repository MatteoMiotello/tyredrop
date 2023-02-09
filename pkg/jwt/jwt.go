package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/models"
	"time"
)

type UserJwtClaims struct {
	jwt.RegisteredClaims
	UserId   int64
	Email    string
	Name     null.String
	Username null.String
}

func CreateAccessTokenFromUser(user models.User) (string, error) {
	expirationMin := viper.GetInt64("security.jwt.access_token_expiration")
	expiration := jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMin) * time.Minute))

	userClaims := UserJwtClaims{
		UserId:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiration,
			Issuer:    viper.GetString("security.jwt.issuer"),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signedString, err := t.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func CreateRefreshToken() (string, error) {
	expirationMin := viper.GetInt64("security.jwt.refresh_token_expiration")
	expiration := jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMin) * time.Minute))

	claims := jwt.RegisteredClaims{
		ExpiresAt: expiration,
		Issuer:    viper.GetString("security.jwt.issuer"),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := t.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
