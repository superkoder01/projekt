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
package logger

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var loggerLevelMap = map[string]zapcore.Level{
	/*For mapping config logger to email_service logger levels*/
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(cfg *configuration.Logger) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func NewApiLogger(cfg *configuration.Logger) *zap.SugaredLogger {
	logLevel := getLoggerLevel(cfg)

	logWriter := zapcore.Lock(os.Stdout)

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.ConsoleSeparator = " "

	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	sugarLogger := logger.Sugar()
	if err := sugarLogger.Sync(); err != nil {
		/*todo fix this log*/
		//l.sugarLogger.Error(err)
	}

	return sugarLogger
}
