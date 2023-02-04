package main

import (
	"flag"
	"fmt"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"log"
	"os"
	"pillowww/titw/internal/bootstrap"

	_ "github.com/lib/pq"
	_ "pillowww/titw/migrations"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func init() {
	bootstrap.InitConfig()
}

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	command := args[0]

	dir := viper.GetString("migrations.dir")

	if dir == "" {
		dir = "./migrations"
	}

	dbString := fmt.Sprintf("postgres://%s@%s/%s?port=%s&password=%s&sslmode=disable",
		viper.Get("MIGRATION_DB_USER"),
		viper.Get("MIGRATION_DB_HOST"),
		viper.Get("MIGRATION_DB_NAME"),
		viper.Get("MIGRATION_DB_PORT"),
		viper.Get("MIGRATION_DB_PASSWORD"))

	db, err := goose.OpenDBWithDriver(viper.GetString("MIGRATION_DB_DRIVER"), dbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
