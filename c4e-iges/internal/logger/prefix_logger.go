package logger

type prefixLogger struct {
	baseLogger Logger
	prefix     string
}

func NewPrefixLogger(logger Logger, prefix string) *prefixLogger {
	return &prefixLogger{baseLogger: logger, prefix: prefix}
}

func (l *prefixLogger) InitLogger() {

}

func (l *prefixLogger) Debug(args ...interface{}) {
	l.baseLogger.Debugf(l.prefix+" %v ", args...)
}

func (l *prefixLogger) Debugf(template string, args ...interface{}) {
	l.baseLogger.Debugf(l.prefix+" "+template, args...)
}

func (l *prefixLogger) Info(args ...interface{}) {
	l.baseLogger.Infof(l.prefix+" %v ", args...)
}

func (l *prefixLogger) Infof(template string, args ...interface{}) {
	l.baseLogger.Infof(l.prefix+" "+template, args...)
}

func (l *prefixLogger) Warn(args ...interface{}) {
	l.baseLogger.Warnf(l.prefix+" %v ", args...)
}

func (l *prefixLogger) Warnf(template string, args ...interface{}) {
	l.baseLogger.Warnf(l.prefix+" "+template, args...)
}

func (l *prefixLogger) Error(args ...interface{}) {
	l.baseLogger.Errorf(l.prefix+" %v ", args...)
}

func (l *prefixLogger) Errorf(template string, args ...interface{}) {
	l.baseLogger.Errorf(l.prefix+" "+template, args...)
}

func (l *prefixLogger) DPanic(args ...interface{}) {
	l.baseLogger.DPanicf(l.prefix+" %v ", args...)
}

func (l *prefixLogger) DPanicf(template string, args ...interface{}) {
	l.baseLogger.DPanicf(l.prefix+" "+template, args...)
}

func (l *prefixLogger) Panic(args ...interface{}) {
	l.baseLogger.Panicf(l.prefix+" %v ", args...)
}

func (l *prefixLogger) Panicf(template string, args ...interface{}) {
	l.baseLogger.Panicf(l.prefix+" "+template, args...)
}

func (l *prefixLogger) Fatal(args ...interface{}) {
	l.baseLogger.Fatalf(l.prefix+" %v ", args...)
}

func (l *prefixLogger) Fatalf(template string, args ...interface{}) {
	l.baseLogger.Fatalf(l.prefix+" "+template, args...)
}
