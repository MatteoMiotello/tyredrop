package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"pillowww/titw/internal/bootstrap"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/jobs"
	"pillowww/titw/pkg/log"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
	bootstrap.InitLog("cron")
}

func addJobs(c *cron.Cron) {
	_, err := c.AddFunc("@every 1s", jobs.ImportProductFromFile)
	if err != nil {
		panic(err.Error())
	}
}

func runCron() {
	c := cron.New()

	addJobs(c)

	c.Start()
}

func main() {
	defer db.Close()

	log.Info("Cron started")
	runCron()
	fmt.Scanln()
}
