package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
	"pillowww/titw/pkg/log"
)

func registerGlobalMiddlewares(router *gin.Engine) {
	router.Use(ginlogrus.Logger(log.Log), gin.Recovery())
	router.Use(cors.Default())
}
