package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"pillowww/titw/internal/repositories"
	"pillowww/titw/internal/services"
	"pillowww/titw/models"
	"pillowww/titw/pkg/api/responses"
	"pillowww/titw/pkg/jwt"
	"pillowww/titw/pkg/security"
	"pillowww/titw/pkg/utils"
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

type SignUpPayload services.CreateUserPayload

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

func (a *AuthController) Login(ctx *gin.Context) {
	loginPayload := new(LoginPayload)

	err := ctx.BindJSON(loginPayload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userRepo := repositories.NewUserRepoWithCtx(ctx)

	var user *models.User

	if utils.IsEmail(loginPayload.Username) {
		user, err = userRepo.FindOneByEmail(loginPayload.Username)
	} else {
		user, err = userRepo.FindOneByUsername(loginPayload.Username)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "incorrect username or password"})
		return
	}

	samePass := security.CheckPassword(user.Password, loginPayload.Password)

	if !samePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "incorrect username or password"})
		return
	}

	accessToken, err := jwt.CreateAccessTokenFromUser(*user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating access token"})
		return
	}

	refreshToken, err := jwt.CreateUniqueRefreshToken()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating refresh token"})
		return
	}

	err = services.NewUserService(user).StoreNewRefreshToken(ctx, refreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error storing refresh token"})
		return
	}

	ctx.SetCookie(
		viper.GetString("security.jwt.cookie-key"),
		accessToken,
		viper.GetInt("security.jwt.expiration"),
		"/",
		viper.GetString("APPLICATION_DOMAIN"),
		true,
		true,
	)

	ctx.JSON(http.StatusOK, LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (a *AuthController) SignUp(ctx *gin.Context) {
	signupPayload := new(SignUpPayload)

	err := ctx.BindJSON(signupPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	userRepo := repositories.NewUserRepoWithCtx(ctx)

	user, _ := userRepo.FindOneByEmail(signupPayload.Email)

	if user != nil && user.DeletedAt.IsZero() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "user with the same email already exists",
		})
		return
	}

	if signupPayload.Username != "" {
		user, _ = userRepo.FindOneByUsername(signupPayload.Username)

		if user != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorResponse{
				Error: "user with the same username already exists",
			})
			return
		}
	}

	user, err = services.CreateUserFromPayload(ctx, services.CreateUserPayload(*signupPayload))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "error creating user: " + err.Error()})
		return
	}

	r, err := userRepo.GetUserRole(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "user role not found for user"})
		return
	}

	rLang, err := repositories.NewUserRoleRepoFromCtx(ctx).GetLanguage(r)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := SignUpResponse{
		Email:    user.Email,
		Username: user.Username.String,
		Name:     user.Name,
		Surname:  user.Surname,
		Role: role{
			Name: rLang.Name,
			Code: r.RoleCode,
		},
	}

	ctx.JSON(http.StatusOK, response)
}
