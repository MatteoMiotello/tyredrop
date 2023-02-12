package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pillowww/titw/internal/auth"
	"pillowww/titw/pkg/api/responses"
	"time"
)

func IsAuthenticated(ctx *gin.Context) {
	access := auth.FromCtx(ctx)

	if access.Expiration.Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "token is expired"})
		return
	}

	ctx.Next()
}
