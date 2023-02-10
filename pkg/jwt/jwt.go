package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/models"
	"time"
)

type UserJwtClaims struct {
	jwt.RegisteredClaims
	UserID   int64
	Email    string
	Name     null.String
	Username null.String
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	uuid uuid.UUID
}

func CreateAccessTokenFromUser(user models.User) (string, error) {
	expirationMin := viper.GetInt64("security.jwt.access_token_expiration")
	expiration := jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMin) * time.Minute))

	userClaims := UserJwtClaims{
		UserID:   user.ID,
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

	if err != nil {
		return nil, err
	}

	return t.Claims.(*UserJwtClaims), nil
}
