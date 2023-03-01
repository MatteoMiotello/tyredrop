package log

import "github.com/sirupsen/logrus"

type customLogger struct {
	*logrus.Logger
	entry *entry
}

type entry struct {
	logEntry *logrus.Entry
}

var Log *customLogger

func New(envName string) *customLogger {
	l := logrus.New()

	Log = &customLogger{
		Logger: l,
		entry:  buildEntry(l, envName),
	}

	return Log
}

func buildEntry(l *logrus.Logger, envName string) *entry {
	return &entry{l.WithField("env", envName)}
}

func GetEntry() *entry {
	return Log.entry
}

func WithField(key string, value interface{}) *entry {
	return Log.entry.WithField(key, value)
}

func (e *entry) WithField(key string, value interface{}) *entry {
	return &entry{e.logEntry.WithField(key, value)}
}

func (e *entry) Info(args ...interface{}) {
	e.logEntry.Info(args)
}

func (e *entry) Warn(args ...interface{}) {
	e.logEntry.Warn(args)
}

func (e *entry) Panic(args ...interface{}) {
	e.logEntry.Panic(args)
}

func (e *entry) Fatal(args ...interface{}) {
	e.logEntry.Fatal(args)
}

func (e *entry) Error(args ...interface{}) {
	e.logEntry.Error(args)
}

func Info(args ...interface{}) {
	Log.entry.Info(args)
}

func Warn(args ...interface{}) {
	Log.entry.Warn(args)
}

func Panic(args ...interface{}) {
	Log.entry.Panic(args)
}

func Fatal(args ...interface{}) {
	Log.entry.Fatal(args)
}

func Error(args ...interface{}) {
	Log.entry.Error(args)
}
