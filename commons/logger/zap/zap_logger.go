/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package zap

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const InternalErrorCode = 2

type zapLogger struct {
	cfg           *config.LoggerConfig
	loggerContext string
	sugarLogger   *zap.SugaredLogger
}

//NewZapLogger creates & initialize logger.Logger implementation using zapcore library
func NewZapLogger(cfg *config.LoggerConfig) *zapLogger {
	logger := zapLogger{cfg: cfg}
	logger.initLogger()
	return &logger
}

//CreateLoggerContext returns instance of logger with prefix. All logs will have [prefix] on the beginning of each log message
func CreateLoggerContext(logger zapLogger, prefix string) zapLogger {
	logger.loggerContext = "[" + prefix + "] "
	return logger
}

var loggerLevelMap = map[string]zapcore.Level{
	/*For mapping config logger to logger levels*/
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *zapLogger) getLoggerLevel(cfg *config.LoggerConfig) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *zapLogger) initLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.Lock(os.Stdout)

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	if len(l.cfg.TimeEncoderFormat) > 1 {
		encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(l.cfg.TimeEncoderFormat)
	} else {
		encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.ConsoleSeparator = " "

	if l.cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err == nil {
		//l.sugarLogger.Error(err)
		os.Exit(InternalErrorCode)
	}
}

func (l zapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l zapLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l zapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l zapLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l zapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l zapLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l zapLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l zapLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l zapLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l zapLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l zapLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l zapLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
