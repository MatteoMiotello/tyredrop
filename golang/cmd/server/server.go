package main

import (
	"pillowww/titw/internal/api"
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/internal/db"
	"pillowww/titw/pkg/log"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
	bootstrap.InitLanguage()
	bootstrap.InitLog("server")
}

func main() {
	defer db.Close()

	log.Info("Application started")
	api.Serve()
}
