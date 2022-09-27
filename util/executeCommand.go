package util

import (
	"gitmm/log"
)

func Execute(command string) (outStr string, errStr string, err error) {
	return ExecShell(command)
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
