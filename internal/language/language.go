package language

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type Language struct {
	*models.Language
}

const languageContextKey string = "language"

var fallbackLanguage *Language

func SetFallbackLanguage(isoCode string) error {
	language, err := models.Languages(qm.Where("iso_code = ?", isoCode)).One(context.Background(), db.DB)

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
