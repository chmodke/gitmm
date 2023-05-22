package log

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"time"
)

var logger *log.Logger
var logFile *os.File
var home, _ = homedir.Dir()
var logPath = fmt.Sprintf("%s/.gitmm", home)
var logfile = fmt.Sprintf("%s/gitmm-%s.log", logPath, time.Now().Format("20060102"))

func init() {
	var err error
	os.MkdirAll(logPath, 0750)
	logFile, err = os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
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
