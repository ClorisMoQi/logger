package logger

import (
	"fmt"
	"runtime"
	"time"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(msg string) {
    now := time.Now()
    _, file, line, ok := runtime.Caller(1)
	fmt.Printf("Level: Debug, ts: [%s], \"%s\"\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l *Logger) Trace(msg string) {
	fmt.Println(msg)
}

func (l *Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l *Logger) Warn(msg string) {
	fmt.Println(msg)
}

func (l *Logger) Error(msg string) {
	fmt.Println(msg)
}

func (l *Logger) Fatal(msg string) {
	fmt.Println(msg)
}