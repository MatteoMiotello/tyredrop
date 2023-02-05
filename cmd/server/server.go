package main

import (
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/internal/db"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
}

func main() {
	defer db.Close()
}
