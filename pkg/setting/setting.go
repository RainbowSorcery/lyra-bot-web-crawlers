package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg      *ini.File
	RunMode  string
	HTTPPort int
)

func Init() {
	var err error
	Cfg, err = ini.Load("conf/application.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/application.ini':%v", err)
	}

	LoadBase()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("app_mode").MustString("develop")
}

func LoadServer() {
	HTTPPort = Cfg.Section("server").Key("http_port").MustInt()
}
