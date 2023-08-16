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
