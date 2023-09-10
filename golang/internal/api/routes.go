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
	router.POST("/reset_password", authController.IssueResetPassword)
	router.POST("/change_password", authController.ChangePassword)

	supportController := controllers.NewSupportController()
	router.POST("/support_request", supportController.SendSupportEmail)

	//graphql
	graphController := new(controllers.GraphqlController)
	router.POST("/query", middlewares.InjectAuth, middlewares.IsAuthenticated, graphController.Query)
	router.GET("/playground", middlewares.TestEnv, graphController.Playground)

	//assets
	assets := router.Group("/assets")
	assets.Static("/img", "./assets/images")

	pAssets := router.Group("/private", middlewares.InjectAuth, middlewares.IsAuthenticated)
	pAssets.Static("/invoices", "./storage/invoices")
	pAssets.Static("/avatar", "./storage/avatar")
}
