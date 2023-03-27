package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func registerGlobalMiddlewares(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(cors.Default())
}
