package util

import (
	"gitmm/log"
)

type Charset string

const (
	UTF8 = Charset("UTF-8")
	GBK  = Charset("GBK")
)

func Execute(command string) (outStr string, errStr string, err error) {
	return ExecShell(command, UTF8)
}
func ExecuteWithCharset(command string, charset Charset) (outStr string, errStr string, err error) {
	return ExecShell(command, charset)
}

func Status(stdout string, stderr string, err error) bool {
	if err != nil {
		return false
	} else {
		return true
	}
}

func Out(stdout string, stderr string, err error) bool {
	if err != nil {
		log.Error("execute fail")
		log.Error(stderr)
		return false
	} else {
		log.Debug("execute succeed")
		log.Debug(stdout)
		return true
	}
}

func GetOut(stdout string, stderr string, err error) (string, bool) {
	if err != nil {
		log.Error("execute fail")
		log.Error(stderr)
		return "", false
	} else {
		log.Debug("execute succeed")
		return stdout, true
	}
}
