package logger

// Logger basic interface for logger implementations
type Logger interface {
	Debugf(template string, args ...interface{})
	Debug(args ...interface{})
	Infof(template string, args ...interface{})
	Info(args ...interface{})
	Errorf(template string, args ...interface{})
	Error(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatal(args ...interface{})
}
