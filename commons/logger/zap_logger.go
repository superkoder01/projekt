package logger

//ZapLogger extender logger.Logger interface, for Zap logger implementation
type ZapLogger interface {
	Logger
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
}
