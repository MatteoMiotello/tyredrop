package main

import (
	"pillowww/titw/internal/api"
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/internal/db"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
	bootstrap.InitLanguage()
}

func main() {
	defer db.Close()

	api.Serve()
}
