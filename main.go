package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lyra-bot-web-crawlers/pkg/setting"
	"net/http"
)

func Init() {
	setting.Init()
}

func main() {
	Init()

	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
