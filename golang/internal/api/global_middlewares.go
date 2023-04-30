package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func registerGlobalMiddlewares(router *gin.Engine) {
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = false
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))
}
