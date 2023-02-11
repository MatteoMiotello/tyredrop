package middlewares

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
	"net/http"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/cookie"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/language"
	"pillowww/titw/internal/domain/user"
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

	uRepo := user.NewUserRepo(db.DB)
	uModel, err := uRepo.FindOneById(ctx, userJwt.UserID)

	if err != nil {
		log.Warningf(ctx, "%s", "user not found for jwt: "+aToken)

		return
	}

	lModel, err := uRepo.GetDefaultLanguage(ctx, *uModel)

	if err != nil {
		log.Warningf(ctx, "%s", "default language not found for user with id: "+string(uModel.ID))
		return
	}

	auth := auth.Auth{
		User: uModel,
		Language: &language.Language{
			L: lModel,
		},
		Expiration: userJwt.ExpiresAt.Time,
		Username:   userJwt.Username,
		Email:      userJwt.Email,
	}

	auth.InsertToCtx(ctx)

	ctx.Next()
}
