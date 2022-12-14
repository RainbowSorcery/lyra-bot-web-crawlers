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
		PonyAPI.GET("/RandomTSImage", api.RandomTSImage)
		PonyAPI.GET("/RandomRDImage", api.RandomRDImage)
		PonyAPI.GET("/RandomRRImage", api.RandomRRImage)
		PonyAPI.GET("/RandomSnowdropImage", api.RandomSnowdropImage)
		PonyAPI.GET("/RandomNyxImage", api.RandomNyxImage)
		PonyAPI.GET("/RandomPinkiePieImage", api.RandomPinkiePieImage)
		PonyAPI.GET("/RandomFlutterShyPieImage", api.RandomFlutterShyPieImage)
		PonyAPI.GET("/RandomAppleJackPieImage", api.RandomAppleJackPieImage)
		PonyAPI.GET("/RandomLunaPieImage", api.RandomLunaImage)
		PonyAPI.GET("/GetRandomPonyImageToTag", api.GetRandomPonyImageToTag)
	}

	return engine
}
