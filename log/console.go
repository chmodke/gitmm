package log

import (
	"fmt"
	"log"
	"os"
)

var console *log.Logger

func init() {
	console = log.New(os.Stdout, "", 0)
}

func Consolef(format string, v ...any) {
	console.Printf(format, v...)
}

func Console(v ...any) {
	console.Print(v...)
}

func Consoleln(v ...any) {
	console.Println(v...)
}
func ConsoleOut(format string, v ...any) {
	fmt.Printf(format, v...)
}
