package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"pillowww/titw/internal/language"
	"pillowww/titw/internal/repositories"
	"pillowww/titw/internal/services"
	"pillowww/titw/pkg/api/responses"
	"pillowww/titw/pkg/jwt"
)

func UserAuthenticated(ctx *gin.Context) {
	aToken, err := ctx.Cookie(viper.GetString("security.jwt.cookie-key"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "cookie not found",
		})
		return
	}

	userJwt, err := jwt.ParseUserJwt(aToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "error parsing jwt",
		})
		return
	}

	uRepo := repositories.NewUserRepoWithCtx(ctx)
	uModel, err := uRepo.FindOneById(userJwt.UserID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "user not found"})
		return
	}
	lModel, err := uRepo.GetDefaultLanguage(*uModel)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	auth := services.AuthService{
		User: uModel,
		Language: &language.Language{
			L: lModel,
		},
	}

	auth.InsertToCtx(ctx)
}
