package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() {
	router := gin.Default()

	registerGlobalMiddlewares(router)
	registerRoutes(router)
	
	err := router.Run(":" + viper.GetString("api.port"))

	if err != nil {
		panic("error starting api")
	}
}
