package bootstrap

import (
	"github.com/spf13/viper"
	"pillowww/titw/internal/domain/language"
)

func InitLanguage() {
	code := viper.GetString("language.default")

	err := language.SetFallbackLanguage(code)
	if err != nil {
		panic("error initializing default language with error: " + err.Error())
	}
}
