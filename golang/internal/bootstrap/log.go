package bootstrap

import (
	"fmt"
	"github.com/evalphobia/logrus_sentry"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"pillowww/titw/pkg/log"
	"time"
)

func InitLog(applicationName string) {
	log.New(applicationName)

	log.Log.SetFormatter(&logrus.JSONFormatter{})

	writer, err := rotatelogs.New(
		fmt.Sprintf("%s/%s", viper.GetString("log.dir"), "%Y-%m-%d.%H:%M:%S"),
		rotatelogs.WithLinkName(viper.GetString("log.dir")),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(viper.GetInt("log.max-days-file-age"))),
		rotatelogs.WithRotationTime(time.Hour*24*time.Duration(viper.GetInt("log.rotation-time-in-days"))),
	)

	if err != nil {
		panic(err.Error())
	}

	log.Log.SetOutput(writer)

	hook, err := logrus_sentry.NewAsyncSentryHook(viper.GetString("SENTRY_DSN"), []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	})

	if err != nil {
		panic(err.Error())
	}

	hook.SetEnvironment(viper.GetString("APPLICATION_ENV"))
	hook.SetDefaultLoggerName(applicationName)
	hook.StacktraceConfiguration.Enable = false
	hook.Timeout = time.Second * 20

	log.Log.Hooks.Add(hook)
}
