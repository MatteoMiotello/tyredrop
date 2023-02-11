package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"google.golang.org/appengine/log"
	"net/http"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/cookie"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/refresh_token"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
	"pillowww/titw/pkg/api/responses"
	"pillowww/titw/pkg/jwt"
	"pillowww/titw/pkg/security"
	"pillowww/titw/pkg/utils"
	"time"
)

type AuthController Controller

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpPayload user.CreateUserPayload

type role struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type SignUpResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Role     role   `json:"role"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token" binding:"required" validate:"jwt"`
}

func (a AuthController) Login(ctx *gin.Context) {
	loginPayload := new(LoginPayload)

	err := ctx.BindJSON(loginPayload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userRepo := user.NewUserRepo(db.DB)

	var uModel *models.User

	if utils.IsEmail(loginPayload.Username) {
		uModel, err = userRepo.FindOneByEmail(ctx, loginPayload.Username)
	} else {
		uModel, err = userRepo.FindOneByUsername(ctx, loginPayload.Username)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "incorrect username or password"})
		return
	}

	samePass := security.CheckPassword(uModel.Password, loginPayload.Password)

	if !samePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "incorrect username or password"})
		return
	}

	accessToken, err := jwt.CreateAccessTokenFromUser(*uModel)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating access token"})
		return
	}

	refreshToken, err := jwt.CreateUniqueRefreshToken()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating refresh token"})
		return
	}

	err = refresh_token.StoreNew(ctx, *uModel, refreshToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error storing refresh token"})
		return
	}

	cookie.StoreAccessToken(ctx, accessToken)

	ctx.JSON(http.StatusOK, LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (a AuthController) SignUp(ctx *gin.Context) {
	signupPayload := new(SignUpPayload)

	err := ctx.BindJSON(signupPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	userRepo := user.NewUserRepo(db.DB)

	uModel, _ := userRepo.FindOneByEmail(ctx, signupPayload.Email)

	if uModel != nil && uModel.DeletedAt.IsZero() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "user with the same email already exists",
		})
		return
	}

	if signupPayload.Username != "" {
		uModel, _ = userRepo.FindOneByUsername(ctx, signupPayload.Username)

		if uModel != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
				Error: "user with the same username already exists",
			})
			return
		}
	}

	uModel, err = user.CreateUserFromPayload(ctx, user.CreateUserPayload(*signupPayload))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating uModel: " + err.Error()})
		return
	}

	r, err := userRepo.GetUserRole(ctx, uModel)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "uModel role not found for uModel"})
		return
	}

	language := *auth.FromCtx(ctx).Language.L
	rLang, err := user.NewUserRepo(db.DB).GetUserRoleLanguage(ctx, r, language)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := SignUpResponse{
		Email:    uModel.Email,
		Username: uModel.Username.String,
		Name:     uModel.Name,
		Surname:  uModel.Surname,
		Role: role{
			Name: rLang.Name,
			Code: r.RoleCode,
		},
	}

	ctx.JSON(http.StatusOK, response)
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (a AuthController) RefreshToken(ctx *gin.Context) {
	payload := new(RefreshTokenPayload)
	rtRepo := refresh_token.NewRefreshTokenRepo(db.DB)
	err := ctx.BindJSON(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	rToken, err := rtRepo.FindValidOneFromRefreshToken(ctx, payload.RefreshToken)
	if err != nil || rToken == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "user not found for this refresh token"})
		return
	}

	rToken.TimeLastUse = null.TimeFrom(time.Now())
	_ = rtRepo.Update(ctx, rToken)

	uModel, err := rtRepo.GetUser(ctx, *rToken)
	token, err := jwt.CreateAccessTokenFromUser(*uModel)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "unable to create jwt token: " + err.Error()})
		return
	}

	response := RefreshTokenResponse{
		AccessToken: token,
	}

	if rToken.ExpiresAt.After(time.Now()) && rToken.ExpiresAt.Before(time.Now().Add(time.Duration(viper.GetInt("security.refresh_token.refresh_threshold"))*time.Minute)) {
		newRToken, err := jwt.CreateUniqueRefreshToken()

		if err != nil {
			log.Warningf(ctx, "%s: %s", "error generating new refresh token", err.Error())
		}

		err = refresh_token.StoreNew(ctx, *uModel, newRToken)

		if err != nil {
			log.Warningf(ctx, "%s: %s", "error generating new refresh token", err.Error())
		}

		response.RefreshToken = newRToken
	}

	cookie.StoreAccessToken(ctx, token)
	ctx.JSON(http.StatusOK, response)
}
