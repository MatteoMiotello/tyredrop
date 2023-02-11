package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pillowww/titw/internal/auth"
	"pillowww/titw/pkg/api/responses"
	"time"
)

func IsAuthenticated(ctx *gin.Context) {
	a := auth.FromCtx(ctx)

	if a.User == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "user not found in request"})
		return
	}

	if a.Expiration.Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "token is expired"})
		return
	}

	ctx.Next()
}
