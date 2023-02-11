package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/domain/language"
	"pillowww/titw/models"
	"time"
)

const ctxKey string = "auth"

type Auth struct {
	User       *models.User
	Language   *language.Language
	Expiration time.Time
	Username   null.String
	Email      string
}

func FromCtx(ctx context.Context) Auth {
	value := ctx.Value(ctxKey)

	fmt.Println(value)

	if value == nil {
		return Auth{
			Language: language.FallbackLanguage(),
			User:     nil,
		}
	} else {
		return value.(Auth)
	}
}

func (a Auth) InsertToCtx(ctx *gin.Context) {
	ctx.Set(ctxKey, a)
}
