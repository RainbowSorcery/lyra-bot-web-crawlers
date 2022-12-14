package routers

import (
	"github.com/gin-gonic/gin"
	"lyra-bot-web-crawlers/pkg/setting"
	"lyra-bot-web-crawlers/routers/api"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	PonyAPI := engine.Group("/randomImage/pony")
	{
		PonyAPI.GET("/randomLyraImg", api.RandomLyraImage)
		PonyAPI.GET("/randomBackgroundPonyImage", api.RandomBackgroundPonyImage)
	}

	return engine
}
