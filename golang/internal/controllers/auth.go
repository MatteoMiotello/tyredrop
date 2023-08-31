package controllers

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"net/http"
	"pillowww/titw/internal/cookie"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/rt"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/internal/email/mailer"
	"pillowww/titw/models"
	"pillowww/titw/pkg/api/responses"
	"pillowww/titw/pkg/jwt"
	"pillowww/titw/pkg/log"
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

type ResetPasswordPayload struct {
	Email string `json:"email" binding:"required" validate:"email"`
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
		AccessToken:  token,
		RefreshToken: rToken.RefreshToken,
	}

	cookie.StoreAccessToken(ctx, token)
	ctx.JSON(http.StatusOK, response)
}

func (a AuthController) IssueResetPassword(ctx *gin.Context) {
	var payload ResetPasswordPayload

	uDao := user.NewDao(db.DB)
	err := ctx.BindJSON(payload)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.ErrorResponse{Error: err.Error()})
		return
	}

	u, err := uDao.FindOneByEmail(ctx, payload.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.ErrorResponse{Error: "user not found" + err.Error()})
		return
	}

	key := make([]byte, 64)
	_, err = rand.Read(key)
	if err != nil {
		log.Error("Error generating key for reset password: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	rp := &models.ResetPassword{
		UserID:   u.ID,
		Token:    string(key),
		IssuedAt: time.Now(),
		ExpiryAt: time.Now().Add(time.Minute * 5),
	}

	err = uDao.Save(ctx, rp)
	if err != nil {
		log.Error("Error saving reset password", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	m := mailer.NewResetPasswordMailer(rp)
	err = m.SendResetEmail(u)

	if err != nil {
		log.Error("Error sending reset password email", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
}
