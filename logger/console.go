package logger

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(identifier string, msg string) {
    now := time.Now()
    _, file, line, ok := runtime.Caller(1)
    if !ok {
        return
    }
    caller := file + ":" + strconv.Itoa(line)
	fmt.Printf("level: Debug, ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
}

func (l *Logger) Trace(identifier string, msg string) {
	now := time.Now()
    _, file, line, ok := runtime.Caller(1)
    if !ok {
        return
    }
    caller := file + ":" + strconv.Itoa(line)
	fmt.Printf("level: Trace, ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
}

func (l *Logger) Info(identifier string, msg string) {
	now := time.Now()
    _, file, line, ok := runtime.Caller(1)
    if !ok {
        return
    }
    caller := file + ":" + strconv.Itoa(line)
	fmt.Printf("Info: Debug, ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
}

func (l *Logger) Warn(identifier string, msg string) {
	now := time.Now()
    _, file, line, ok := runtime.Caller(1)
    if !ok {
        return
    }
    caller := file + ":" + strconv.Itoa(line)
	fmt.Printf("Warn: Debug, ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
}

func (l *Logger) Error(identifier string, msg string) {
	now := time.Now()
    _, file, line, ok := runtime.Caller(1)
    if !ok {
        return
    }
    caller := file + ":" + strconv.Itoa(line)
	fmt.Printf("Error: Debug, ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
}

func (l *Logger) Fatal(identifier string, msg string) {
	now := time.Now()
    _, file, line, ok := runtime.Caller(1)
    if !ok {
        return
    }
    caller := file + ":" + strconv.Itoa(line)
	fmt.Printf("level: Fatal, ts: [%s], caller: %s, identifier: %s, message: \"%s\"\n", now.Format("2006-01-02 15:04:05"), caller, identifier, msg)
}