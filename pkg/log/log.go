package log

import "github.com/sirupsen/logrus"

type customLogger struct {
	*logrus.Logger
	environmentName string
}

var Log *customLogger

func New(envName string) *customLogger {
	Log = &customLogger{
		logrus.New(),
		envName,
	}

	return Log
}

func (c customLogger) getEntry() *logrus.Entry {
	return c.WithField("env", c.environmentName)
}

func GetEntry() *logrus.Entry {
	return Log.getEntry()
}

func Info(args ...interface{}) {
	Log.getEntry().Info(args)
}

func Warning(args ...interface{}) {
	Log.getEntry().Warn(args)
}

func Panic(args ...interface{}) {
	Log.getEntry().Panic(args)
}

func Fatal(args ...interface{}) {
	Log.getEntry().Fatal(args)
}

func Error(args ...interface{}) {
	Log.getEntry().Error(args)
}
