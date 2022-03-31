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
    lv := parseLogLevel(lvstr)
	return &Logger{
        Lv: lv,
    }
}

func (l *Logger) Debug(identifier string, msg string) {
    if l.Lv >= DEBUG {
        now := time.Now()
        _, file, line, ok := runtime.Caller(1)
        if !ok {
            return
        }
        caller := file + ":" + strconv.Itoa(line)
        fmt.Printf("[Debug], ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
    }
}

func (l *Logger) Trace(identifier string, msg string) {
    if l.Lv >= TRACE {
        now := time.Now()
        _, file, line, ok := runtime.Caller(1)
        if !ok {
            return
        }
        caller := file + ":" + strconv.Itoa(line)
	    fmt.Printf("[Trace], ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
    }
}

func (l *Logger) Info(identifier string, msg string) {
    if l.Lv >= INFO {
        now := time.Now()
        _, file, line, ok := runtime.Caller(1)
        if !ok {
            return
        }
        caller := file + ":" + strconv.Itoa(line)
        fmt.Printf("[Info] [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
    }
}

func (l *Logger) Warn(identifier string, msg string) {
    if l.Lv >= WARN {
        now := time.Now()
        _, file, line, ok := runtime.Caller(1)
        if !ok {
            return
        }
        caller := file + ":" + strconv.Itoa(line)
        fmt.Printf("[Warn] [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
    }
}

func (l *Logger) Error(identifier string, msg string) {
    if l.Lv >= ERROR {
        now := time.Now()
        _, file, line, ok := runtime.Caller(1)
        if !ok {
            return
        }
        caller := file + ":" + strconv.Itoa(line)
        fmt.Printf("[Error] [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
    }
}

func (l *Logger) Fatal(identifier string, msg string) {
    if l.Lv >= FATAL {
        now := time.Now()
        _, file, line, ok := runtime.Caller(1)
        if !ok {
            return
        }
        caller := file + ":" + strconv.Itoa(line)
	    fmt.Printf("[Fatal] [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
    }
}