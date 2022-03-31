package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open file failed: %v\n", err)
		return
	}
	// log.SetOutput(os.Stdout) //输出到命令行
	log.SetOutput(fileObj) //输出到文件fileObj
	for {
		log.Printf("This is a test log.")
		time.Sleep(time.Second * 2)
	}
}