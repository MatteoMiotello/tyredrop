package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"pillowww/titw/pkg/log"
)

func Serve() {
	router := gin.Default()
	gin.DefaultErrorWriter = log.Log.Writer()

	router.ContextWithFallback = true
	registerGlobalMiddlewares(router)
	registerRoutes(router)

	err := router.Run(":" + viper.GetString("api.port"))

	if err != nil {
		panic("error starting api")
	}
}
