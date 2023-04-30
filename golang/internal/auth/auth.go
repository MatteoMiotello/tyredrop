package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"golang.org/x/net/context"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/language"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
	"time"
)

const ctxKey string = "auth"

type Role struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Auth struct {
	user         *models.User
	language     *language.Language
	Expiration   time.Time
	UserID       int64       `json:"userID"`
	Username     null.String `json:"username"`
	Email        string      `json:"email"`
	Role         Role
	LanguageCode string
}

func FromCtx(ctx context.Context) (access *Auth) {
	value := ctx.Value(ctxKey)

	fmt.Println(value)

	if value == nil {
		return &Auth{
			language:     language.FallbackLanguage(),
			LanguageCode: language.FallbackLanguage().L.IsoCode,
			user:         nil,
		}
	} else {
		a := value.(*Auth)

		return &Auth{
			UserID:       a.UserID,
			Expiration:   a.Expiration,
			Username:     a.Username,
			Email:        a.Email,
			LanguageCode: a.LanguageCode,
			Role: Role{
				Name: a.Role.Name,
				Code: a.Role.Code,
			},
		}
	}
}

func CurrentLanguage(ctx context.Context) *language.Language {
	return FromCtx(ctx).GetLanguage(ctx)
}

func (a *Auth) InsertToCtx(ctx *gin.Context) {
	ctx.Set(ctxKey, a)
}

func (a *Auth) GetUser(ctx context.Context) (*models.User, error) {
	if a.user != nil {
		return a.user, nil
	}

	uRepo := user.NewDao(db.DB)
	uModel, err := uRepo.FindOneById(ctx, a.UserID)

	if err != nil {
		return nil, err
	}
	a.user = uModel

	return uModel, nil
}

func (a *Auth) GetLanguage(ctx context.Context) *language.Language {
	if a.language != nil {
		return a.language
	}

	lRepo := language.NewDao(db.DB)
	lModel, err := lRepo.FindOneFromIsoCode(ctx, a.LanguageCode)
	if err != nil {
		return language.FallbackLanguage()
	}

	l := language.Language{
		L: lModel,
	}

	a.language = &l
	return &l
}
