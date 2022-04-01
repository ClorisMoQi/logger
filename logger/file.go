package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

type FLogger struct {
	Lv          LogLevel
	FilePath    string //日志文件保存的路径
	FileName    string //日志文件保存的文件名
	ErrFileName string
    FileObj     *os.File
    ErrFileObj  *os.File
	MaxSize     int64
}

func NewFLogger(lvstr, filepath, filename string, maxsize int64) *FLogger {
	logLv := LogLvStr2Int(lvstr)
    fl := &FLogger{
		Lv:       logLv,
		FilePath: filepath,
		FileName: filename,
		MaxSize:  maxsize,
	}
    err := fl.initFile()
    if err != nil {
        panic(err) //日志文件没有成功打开
    }
	return fl
}

func (l *FLogger) Debug(identifier string, format string, args ...interface{}) {
	l.log(DEBUG, identifier, format, args...)
}

func (l *FLogger) Trace(identifier string, format string, args ...interface{}) {
	l.log(TRACE, identifier, format, args...)
}

func (l *FLogger) Info(identifier string, format string, args ...interface{}) {
	l.log(INFO, identifier, format, args...)
}

func (l *FLogger) Warn(identifier string, format string, args ...interface{}) {
	l.log(WARN, identifier, format, args...)
}

func (l *FLogger) Error(identifier string, format string, args ...interface{}) {
	l.log(ERROR, identifier, format, args...)
}

func (l *FLogger) Fatal(identifier string, format string, args ...interface{}) {
	l.log(FATAL, identifier, format, args...)
}

func (l *FLogger) log(lv LogLevel, identifier string, format string, args ...interface{}) {
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
		fmt.Fprintf(l.FileObj, "[%s] [%s] caller: %s, func:[%s] identifier: %s, message: %s\n", lvstr, now.Format("2006-01-02 15:04:05"), caller, funcname, identifier, msg)
        if lv >= ERROR {
		    fmt.Fprintf(l.ErrFileObj, "[%s] [%s] caller: %s, func:[%s] identifier: %s, message: %s\n", lvstr, now.Format("2006-01-02 15:04:05"), caller, funcname, identifier, msg)
        }
    }
}

func (l *FLogger) initFile() error {
    fullpath := path.Join(l.FilePath, l.FileName)
    fileObj, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("open log file failed, err: %v\n", err)
        return err
    }
    errfullpath := path.Join(l.FilePath, l.ErrFileName)
    errfileObj, err := os.OpenFile(errfullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("open log file failed, err: %v\n", err)
        return err
    }
    l.FileObj = fileObj
    l.ErrFileObj = errfileObj
    return nil
}

func (l *FLogger) close() {
    l.FileObj.Close()
    l.ErrFileObj.Close()
}