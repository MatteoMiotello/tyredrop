package api

import (
	"github.com/gin-gonic/gin"
	"pillowww/titw/internal/api/middlewares"
	"pillowww/titw/internal/controllers"
)

func registerRoutes(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/register", authController.SignUp)
	router.POST("/refresh_token", authController.RefreshToken)

	graphController := new(controllers.GraphqlController)
	router.POST("/query", middlewares.InjectAuth, middlewares.IsAuthenticated, graphController.Query)
	router.GET("/playground", middlewares.TestEnv, graphController.Playground)
}
