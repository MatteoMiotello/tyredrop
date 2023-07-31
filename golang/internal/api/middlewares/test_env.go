package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"pillowww/titw/pkg/api/responses"
)

func TestEnv(ctx *gin.Context) {
	env := viper.GetString("APPLICATION_ENV")

	if env != "test" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "not in test env"})
		return
	}
}
