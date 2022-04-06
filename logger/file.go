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
    LogChan     chan *logMsg
}

func NewFLogger(lvstr, filepath, filename string, maxsize int64) *FLogger {
	logLv := LogLvStr2Int(lvstr)
    fl := &FLogger{
		Lv:       logLv,
		FilePath: filepath,
		FileName: filename,
        ErrFileName: filename + ".err",
		MaxSize:  maxsize,
        LogChan: make(chan *logMsg, ChanSize),
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
		pc, file, line, ok := runtime.Caller(2)
		if !ok {
			return
		}
		funcname := runtime.FuncForPC(pc).Name()
		// 把日志发送到通道中
        logTmp := &logMsg{
            lv: lv,
            identifier:identifier,
            msg: msg,
            funcName: funcname,
            line:line,
            fileName: file,
            timestamp: now.Format("2006-01-02 15:04:05"),
        }
        select {
        case l.LogChan <- logTmp:
        default:
            // 当LogChan满了，日志被丢掉，保证不出现阻塞
            time.Sleep(time.Millisecond * 500)
        }
        
    }
}

func (l *FLogger) initFile() error {
    fullpath := path.Join(l.FilePath, l.FileName)
    fmt.Println(fullpath)
    fileObj, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("open log file failed, err: %v\n", err)
        return err
    }
    errfullpath := path.Join(l.FilePath, l.ErrFileName)
    errfileObj, err := os.OpenFile(errfullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("open log file failed, err: %v\n", err)
        return err
    }
    l.FileObj = fileObj
    l.ErrFileObj = errfileObj
    go l.write() // 开启后台goroutine写log到文件里
    return nil
}

// func (l *FLogger)close() {
//     l.FileObj.Close()
//     l.ErrFileObj.Close()
// }

func (l *FLogger)checksize(fileObj *os.File) bool {
    fileInfo, err := fileObj.Stat()
    if err != nil {
        fmt.Printf("get file info failed, err:%v\n", err)
        return false
    }
    return fileInfo.Size() >= l.MaxSize
}

func (l *FLogger)splitFile(fileObj *os.File) (*os.File) {
    // 1. 关闭当前日志文件
    fileInfo, err := fileObj.Stat()
    if err != nil {
        return nil
    }
    curFileName := path.Join(l.FilePath, fileInfo.Name())
    fileObj.Close()
    // 2. 备份并重命名文件
    nowstr := time.Now().Format("20060102150405000")
    newFileName := fmt.Sprintf("%s.bak%s", curFileName, nowstr)
    os.Rename(curFileName, newFileName)
    // 3. 打开一个新的日志文件
    fileObjNew, err := os.OpenFile(curFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("Open new log file failed, err:%v\n", err)
        return nil
    }
    // 4. 将打开的新日志文件对象赋值给l.FileObj
    return fileObjNew
}

// write log in background
func (l *FLogger) write() {
    for {
        select {
        case logTmp := <-l.LogChan:
            lvstr := LogLvInt2Str(l.Lv)
            caller := logTmp.fileName + ":" + strconv.Itoa(logTmp.line)
            logMsg := fmt.Sprintf("[%s] [%s] caller: %s, func:[%s] identifier: %s, message: %s\n", logTmp.timestamp, lvstr, caller, logTmp.funcName, logTmp.identifier, logTmp.msg)
            fmt.Fprintf(l.FileObj, logMsg)
            if l.checksize(l.FileObj) {
                // 需要切割日志文件
                fileObjNew := l.splitFile(l.FileObj)
                l.FileObj = fileObjNew
            }
            if logTmp.lv >= ERROR {
                fmt.Fprintf(l.ErrFileObj, logMsg)
                if l.checksize(l.ErrFileObj) {
                    errFileObjNew := l.splitFile(l.ErrFileObj)
                    l.ErrFileObj = errFileObjNew
                }
            }
        default:
            time.Sleep(time.Millisecond * 500)
        }
        
    }
}