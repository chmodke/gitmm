package log

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger
var logFile *os.File

func init() {
	var err error
	logFile, err = os.OpenFile("gitmm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error :", err)
		return
	}
	logger = log.New(logFile, "", log.LstdFlags)
}

func Printf(format string, v ...any) {
	logger.Printf(format, v...)
}

func Print(v ...any) {
	logger.Print(v...)
}

func Println(v ...any) {
	logger.Println(v...)
}

func FlushAndClose() {
	defer logFile.Close()
}
