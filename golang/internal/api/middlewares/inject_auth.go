package middlewares

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
	"net/http"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/cookie"
	"pillowww/titw/pkg/api/responses"
	"pillowww/titw/pkg/jwt"
)

func InjectAuth(ctx *gin.Context) {
	aToken, err := cookie.RetrieveAccessToken(ctx)

	if err != nil {
		log.Warningf(ctx, "%s", "auth cookie not found")
		return
	}

	userJwt, jwtErr := jwt.ParseUserJwt(aToken)

	if jwtErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{Error: "token expired or not valid: " + jwtErr.Error()})

		return
	}

	a := auth.Auth{
		Expiration:   userJwt.ExpiresAt.Time,
		Username:     userJwt.Username,
		Email:        userJwt.Email,
		Role:         userJwt.Role,
		LanguageCode: userJwt.Language,
	}

	a.InsertToCtx(ctx)
	ctx.Next()
}
