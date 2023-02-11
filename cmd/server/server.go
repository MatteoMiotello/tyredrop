package main

import (
	"github.com/getsentry/sentry-go"
	"pillowww/titw/internal/api"
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/internal/db"
	"time"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
	bootstrap.InitLanguage()
	bootstrap.InitLog()
}

func main() {
	defer db.Close()
	defer sentry.Flush(2 * time.Second) //to check

	api.Serve()
}
