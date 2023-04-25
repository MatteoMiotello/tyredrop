package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"google.golang.org/appengine/log"
	"net/http"
	"pillowww/titw/internal/cookie"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/rt"
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
	Email    string `json:"email" binding:"required"`
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

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (a AuthController) createTokens(ctx *gin.Context, uModel *models.User) {
	accessToken, err := jwt.CreateAccessTokenFromUser(ctx, *uModel)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating access token", StatusCode: 4003})
		return
	}

	refreshToken, err := jwt.CreateUniqueRefreshToken()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating refresh token", StatusCode: 4004})
		return
	}

	rtService := rt.NewRefreshTokenService(rt.NewDao(db.DB))
	err = rtService.StoreNew(ctx, *uModel, refreshToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error storing refresh token", StatusCode: 4005})
		return
	}

	cookie.StoreAccessToken(ctx, accessToken)
	ctx.JSON(http.StatusOK, LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (a AuthController) Login(ctx *gin.Context) {
	loginPayload := new(LoginPayload)

	err := ctx.BindJSON(loginPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid username or password", StatusCode: 4000})
		return
	}

	userRepo := user.NewDao(db.DB)

	var uModel *models.User

	if utils.IsEmail(loginPayload.Email) {
		uModel, err = userRepo.FindOneByEmail(ctx, loginPayload.Email)
	} else {
		uModel, err = userRepo.FindOneByUsername(ctx, loginPayload.Email)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "incorrect username or password", StatusCode: 4001})
		return
	}

	samePass := security.CheckPassword(uModel.Password, loginPayload.Password)

	if !samePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "incorrect username or password", StatusCode: 4002})
		return
	}

	a.createTokens(ctx, uModel)
}

func (a AuthController) SignUp(ctx *gin.Context) {
	signupPayload := new(SignUpPayload)

	err := ctx.BindJSON(signupPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	userDao := user.NewDao(db.DB)

	uModel, _ := userDao.FindOneByEmail(ctx, signupPayload.Email)

	if uModel != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:      "user with the same email already exists",
			StatusCode: 5000,
		})
		return
	}

	if signupPayload.Username != "" {
		uModel, _ = userDao.FindOneByUsername(ctx, signupPayload.Username)

		if uModel != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
				Error:      "user with the same username already exists",
				StatusCode: 5001,
			})
			return
		}
	}

	uService := user.NewUserService(userDao)
	uModel, err = uService.CreateUserFromPayload(ctx, user.CreateUserPayload(*signupPayload))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: "error creating uModel: " + err.Error(),
		})
		return
	}

	a.createTokens(ctx, uModel)
}

func (a AuthController) RefreshToken(ctx *gin.Context) {
	payload := new(RefreshTokenPayload)
	rtDao := rt.NewDao(db.DB)
	err := ctx.BindJSON(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	rToken, err := rtDao.FindValidOneFromRefreshToken(ctx, payload.RefreshToken)
	if err != nil || rToken == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "refresh token not found"})
		return
	}

	rToken.TimeLastUse = null.TimeFrom(time.Now())
	_ = rtDao.Update(ctx, rToken)

	uModel, err := rtDao.GetUser(ctx, *rToken)
	token, err := jwt.CreateAccessTokenFromUser(ctx, *uModel)

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

		rtService := rt.NewRefreshTokenService(rtDao)
		err = rtService.StoreNew(ctx, *uModel, newRToken)

		if err != nil {
			log.Warningf(ctx, "%s: %s", "error generating new refresh token", err.Error())
		}

		response.RefreshToken = newRToken
	}

	cookie.StoreAccessToken(ctx, token)
	ctx.JSON(http.StatusOK, response)
}
