package api

import (
	"github.com/gin-gonic/gin"
	"pillowww/titw/internal/api/middlewares"
	"pillowww/titw/internal/controllers"
)

type SpecificationValues struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Response struct {
	Name                string                `json:"name"`
	Code                string                `json:"code"`
	Price               string                `json:"price"`
	Brand               string                `json:"brand"`
	Supplier            string                `json:"supplier"`
	SpecificationValues []SpecificationValues `json:"specification_values"`
}

func registerRoutes(router *gin.Engine) {
	//auth
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/register", authController.SignUp)
	router.POST("/refresh_token", authController.RefreshToken)

	//graphql
	graphController := new(controllers.GraphqlController)
	router.POST("/query", graphController.Query)
	router.GET("/playground", middlewares.TestEnv, graphController.Playground)

	//assets
	group := router.Group("/assets")
	group.Static("/img", "./assets/images")
}
