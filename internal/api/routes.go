package api

import (
	"github.com/gin-gonic/gin"
	"pillowww/titw/internal/controllers"
)

func registerRoutes(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/register", authController.SignUp)

	graphController := new(controllers.GraphqlController)
	router.POST("/query", graphController.Query)
	router.GET("/playground", graphController.Playground)
}
