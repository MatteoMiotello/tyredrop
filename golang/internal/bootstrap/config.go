package bootstrap

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func InitConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic("init config failed " + err.Error())
	}

	files, err := os.ReadDir("./config/")

	if err != nil {
		panic("init config failed reading config files " + err.Error())
	}

	viper.AddConfigPath("./config")

	for _, file := range files {
		name := file.Name()

		parts := strings.Split(name, ".")
		n := parts[0]
		e := parts[1]

		viper.SetConfigName(n)
		viper.SetConfigType(e)
		err := viper.MergeInConfig()

		if err != nil {
			panic("init config failed setting config variables " + err.Error())
		}
	}
}
