package cookie

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StoreAccessToken(ctx *gin.Context, accessToken string) {
	ctx.SetCookie(
		viper.GetString("security.jwt.cookie-key"),
		accessToken,
		viper.GetInt("security.jwt.expiration"),
		"/",
		"."+viper.GetString("APPLICATION_DOMAIN"),
		true,
		true,
	)
}

func RetrieveAccessToken(ctx *gin.Context) (string, error) {
	return ctx.Cookie(viper.GetString("security.jwt.cookie-key"))
}
