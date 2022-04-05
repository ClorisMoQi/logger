package main

import (
	"fmt"
	"reflect"

	"github.com/ClorisMoQi/logger/logger"
)

func main() {
	// clog.Fatal("1111111")
	clog := logger.NewCLogger("debug")
	id := 10001
	name := "testLog"
	clog.Debug("test", "This is a debug log, id: %d, name:%s", id, name)

	fmt.Println(reflect.TypeOf(clog))
	var l logger.Logger
	l = logger.NewFLogger("debug", "./", "test1", 10)
	fmt.Println(reflect.TypeOf(l))
	fmt.Println(reflect.ValueOf(l))

	for {
		flog := logger.NewFLogger("debug", "./", "test.log", 10*1024)
		flog.Debug("Debug", "This is a debug log, id:%d, name:%s", id, name)
		flog.Error("Error", "This is an error log, id:%d, name:%s", id, name)
		break
	}	
}
