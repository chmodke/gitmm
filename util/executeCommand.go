package util

import "gitmm/log"

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
		log.Debug(stderr)
		return false
	} else {
		return true
	}
}

func GetOut(stdout string, stderr string, err error) (string, bool) {
	if err != nil {
		log.Debug(stderr)
		return "", false
	} else {
		return stdout, true
	}
}

func GetErr(stdout string, stderr string, err error) (string, string, bool) {
	if err != nil {
		log.Debug(stderr)
		return "", stderr, false
	} else {
		return stdout, "", true
	}
}
