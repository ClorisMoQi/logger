package main

import (
	"github.com/ClorisMoQi/logger/logger"
)

func main() {
	// log.Fatal("1111111")
	logger := logger.NewLogger("debug")
	id := 10010
	name := "testLog"
	logger.Debug("test", "This is a debug log, id: %d, name:%s\n", id, name)
}
