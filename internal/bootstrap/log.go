package bootstrap

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func InitLog() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              viper.GetString("log.sentry-dsn"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		panic("Sentry initialization failed")
	}
}
