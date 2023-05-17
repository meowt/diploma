package server

import (
	"Diploma/app/modules"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Start(handlers modules.HandlerModule) (err error) {
	router := SetupGinRouter(handlers)
	if err = router.Run(viper.GetString("server.address")); err != nil {
		return err
	}
	return
}

func SetupGinRouter(handlers modules.HandlerModule) *gin.Engine {
	router := gin.Default()
	router.Use(
		CORSMiddleware(),
		gin.Recovery(),
		gin.Logger(),
	)
	handlers.AuthHandler.InitAuthRoutes(router)
	handlers.UserHandler.InitUserRoutes(router)
	handlers.ThemeHandler.InitThemeRoutes(router)
	handlers.MediaHandler.InitMediaRoutes(router)
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
