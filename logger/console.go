package logger

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type Logger struct {
	Lv LogLevel
}

func NewLogger(lvstr string) *Logger {
	lv := LogLvStr2Int(lvstr)
	return &Logger{
		Lv: lv,
	}
}

func (l *Logger) Debug(identifier string, format string, args ...interface{}) {
	l.log(DEBUG, identifier, format, args...)
}

func (l *Logger) Trace(identifier string, format string, args ...interface{}) {
	l.log(TRACE, identifier, format, args...)
}

func (l *Logger) Info(identifier string, format string, args ...interface{}) {
	l.log(INFO, identifier, format, args...)
}

func (l *Logger) Warn(identifier string, format string, args ...interface{}) {
	l.log(WARN, identifier, format, args...)
}

func (l *Logger) Error(identifier string, format string, args ...interface{}) {
	l.log(ERROR, identifier, format, args...)
}

func (l *Logger) Fatal(identifier string, format string, args ...interface{}) {
	l.log(FATAL, identifier, format, args...)
}

func (l *Logger) log(lv LogLevel, identifier string, format string, args ...interface{}) {
	if l.Lv <= lv {
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		lvstr := LogLvInt2Str(lv)
		pc, file, line, ok := runtime.Caller(2)
		if !ok {
			return
		}
		funcname := runtime.FuncForPC(pc).Name()
		caller := file + ":" + strconv.Itoa(line)
		fmt.Printf("[%s] [%s] caller: %s, func:[%s] identifier: %s, message: %s\n", lvstr, now.Format("2006-01-02 15:04:05"), caller, funcname, identifier, msg)
	}
}
