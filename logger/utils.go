package logger

import "strings"

type Logger interface{
    Debug(identifier string, format string, args ...interface{})
    Trace(identifier string, format string, args ...interface{})
    Info(identifier string, format string, args ...interface{})
    Warn(identifier string, format string, args ...interface{})
    Error(identifier string, format string, args ...interface{})
    Fatal(identifier string, format string, args ...interface{})
}

type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

func LogLvStr2Int(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG
	case "trace":
		return TRACE
    case "info":
        return INFO
    case "warn":
        return WARN
    case "error":
        return ERROR
    case "fatal":
        return FATAL
    default:
        return DEBUG
	}
}

func LogLvInt2Str(lv LogLevel) string {
    switch lv {
    case DEBUG:
        return "DEBUG"
    case TRACE:
        return "TRACE"
    case INFO:
        return "INFO"
    case WARN:
        return "WARN"
    case ERROR:
        return "ERROR"
    case FATAL:
        return "FATAL"
    default:
        return "DEBUG"
    }
}