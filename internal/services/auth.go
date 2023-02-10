package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"pillowww/titw/internal/language"
	"pillowww/titw/models"
)

const ctxKey string = "auth"

type AuthService struct {
	User     *models.User
	Language *language.Language
}

func AuthServiceFromCtx(ctx context.Context) AuthService {
	value := ctx.Value(ctxKey)

	if value == nil {
		return AuthService{
			Language: language.FallbackLanguage,
			User:     nil,
		}
	} else {
		return value.(AuthService)
	}
}

func (a AuthService) InsertToCtx(ctx *gin.Context) {
	ctx.Set(ctxKey, a)
}
