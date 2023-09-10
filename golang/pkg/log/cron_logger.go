package log

type CronLogger struct {
	*customLogger
}

func NewCronLogger() *CronLogger {
	return &CronLogger{
		Log,
	}
}

func (r *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	r.Infoln(msg, keysAndValues)
}

// Error logs an error condition.
func (r *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	r.Errorln(err, msg, keysAndValues)
}
