package main

import (
	"fmt"
	"lyra-bot-web-crawlers/pkg/setting"
	"lyra-bot-web-crawlers/routers"
	"net/http"
)

func Init() {
	setting.Init()
}

func main() {
	Init()

	router := routers.InitRouter()

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
