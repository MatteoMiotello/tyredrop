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
	_, err := c.AddFunc("@every 15s", jobs.ImportProductsFromFile)
	if err != nil {
		panic(err.Error())
	}
	_, err = c.AddFunc("@every 1m", jobs.CopySupplierFiles)
	if err != nil {
		panic(err.Error())
	}
	//_, err = c.AddJob("@every 1s", jobs.UpdateTyresSpecificationJob{MaxChildren: 20})
	//if err != nil {
	//	panic(err.Error())
	//}
}

func runCron() {
	c := cron.New(
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
		cron.WithLogger(cron.VerbosePrintfLogger(log.Log)))
	addJobs(c)

	c.Start()
}

func main() {
	defer db.Close()

	log.Info("Cron started")
	fmt.Println("Cron started")
	runCron()

	select {}
}
