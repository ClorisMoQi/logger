package main

import (
	"github.com/ClorisMoQi/logger/logger"
)

func main() {
	// log.Fatal("1111111")
	logger := logger.NewCLogger("debug")
	id := 10010
	name := "testLog"
	logger.Debug("test", "This is a debug log, id: %d, name:%s\n", id, name)

	log = logger.NewFLogger("debug", "./", "test.log", 10*1024*1024)
	
}
