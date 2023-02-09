package language

import (
	"context"
	"github.com/gin-gonic/gin"
	"pillowww/titw/internal/repositories"
	"pillowww/titw/models"
)

type Language struct {
	*models.Language
}

const languageContextKey string = "language"

var fallbackLanguage *Language

func SetFallbackLanguage(isoCode string) error {
	language, err := repositories.NewLanguageRepoFromCtx(context.Background()).FindOneFromIsoCode(isoCode)

	if err != nil {
		return err
	}

	fLanguage := Language{
		Language: language,
	}

	fallbackLanguage = &fLanguage

	return nil
}

func (l *Language) SetToContext(ctx *gin.Context) {
	ctx.Set(languageContextKey, l)
}

func FromContext(ctx context.Context) *Language {
	ctxLang := ctx.Value(languageContextKey)
	var language *Language

	if ctxLang != nil {
		language = ctxLang.(*Language)
	} else {
		language = fallbackLanguage
	}

	return language
}
