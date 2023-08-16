package util

import "github.com/chmodke/gitmm/log"

type Charset string

const (
	UTF8 = Charset("UTF-8")
	GBK  = Charset("GBK")
)

func Execute(workDir string, command string) (outStr string, errStr string, err error) {
	return ExecShell(workDir, command, UTF8)
}
func ExecuteWithCharset(workDir string, command string, charset Charset) (outStr string, errStr string, err error) {
	return ExecShell(workDir, command, charset)
}

func Status(stdout string, stderr string, err error) bool {
	if err != nil {
		log.Println(stderr)
		return false
	} else {
		return true
	}
}

func GetOut(stdout string, stderr string, err error) (string, bool) {
	if err != nil {
		log.Println(stderr)
		return "", false
	} else {
		return stdout, true
	}
}

func GetErr(stdout string, stderr string, err error) (string, string, bool) {
	if err != nil {
		log.Println(stderr)
		return "", stderr, false
	} else {
		return stdout, "", true
	}
}
