package logger

import "strings"

type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

func parseLogLevel(s string) LogLevel {
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
