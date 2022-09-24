package log

import (
	"fmt"
	"strings"
	"time"
)

var level = INFO

var logLevelMap = make(map[int8]string)
var logNameMap = make(map[string]int8)

const (
	GlobalFmt       = "%s %5s %s\n"
	DateFormat      = "2006-01-02 15:04:05.000"
	DEBUG      int8 = 1
	INFO       int8 = 3
	WARN       int8 = 5
	ERROR      int8 = 7
)

func init() {
	logLevelMap[DEBUG] = "DEBUG"
	logLevelMap[INFO] = "INFO"
	logLevelMap[WARN] = "WARN"
	logLevelMap[ERROR] = "ERROR"

	logNameMap["DEBUG"] = DEBUG
	logNameMap["D"] = DEBUG
	logNameMap["INFO"] = INFO
	logNameMap["I"] = INFO
	logNameMap["WARN"] = WARN
	logNameMap["W"] = WARN
	logNameMap["ERROR"] = ERROR
	logNameMap["E"] = ERROR

}

func SetLevel(ls string) {
	ls = strings.ToUpper(ls)
	if _, ok := logNameMap[ls]; ok {
		level = logNameMap[ls]
	}
}

func Debug(msg string) {
	if level <= DEBUG {
		write(DEBUG, msg)
	}
}

func Debugf(template string, args ...interface{}) {
	if level <= DEBUG {
		write(DEBUG, fmt.Sprintf(template, args...))
	}
}

func DebugO(msg string) {
	if level <= DEBUG {
		out(msg)
	}
}

func DebugOf(template string, args ...interface{}) {
	if level <= DEBUG {
		out(fmt.Sprintf(template, args...))
	}
}

func Info(msg string) {
	if level <= INFO {
		write(INFO, msg)
	}
}
func Infof(template string, args ...interface{}) {
	if level <= INFO {
		write(INFO, fmt.Sprintf(template, args...))
	}
}

func InfoO(msg string) {
	if level <= INFO {
		out(msg)
	}
}

func InfoOf(template string, args ...interface{}) {
	if level <= INFO {
		out(fmt.Sprintf(template, args...))
	}
}

func Warn(msg string) {
	if level <= WARN {
		write(WARN, msg)
	}
}
func Warnf(template string, args ...interface{}) {
	if level <= WARN {
		write(WARN, fmt.Sprintf(template, args...))
	}
}

func WarnO(msg string) {
	if level <= WARN {
		out(msg)
	}
}

func WarnOf(template string, args ...interface{}) {
	if level <= WARN {
		out(fmt.Sprintf(template, args...))
	}
}

func Error(msg string) {
	if level <= ERROR {
		write(ERROR, msg)
	}
}
func Errorf(template string, args ...interface{}) {
	if level <= ERROR {
		write(ERROR, fmt.Sprintf(template, args...))
	}
}

func ErrorO(msg string) {
	if level <= ERROR {
		out(msg)
	}
}

func ErrorOf(template string, args ...interface{}) {
	if level <= ERROR {
		out(fmt.Sprintf(template, args...))
	}
}

func write(level int8, msg string) {
	fmt.Printf(GlobalFmt, time.Now().Format(DateFormat), logLevelMap[level], msg)
}

func out(msg string) {
	fmt.Println(msg)
}
