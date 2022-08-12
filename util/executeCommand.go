package util

import (
	"bytes"
	"gitmm/log"
	"os/exec"
	"runtime"
	"strings"
)

func Execute(command string) (outStr string, errStr string, err error) {
	log.Debugf("command: %s", command)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/c", command)
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	outStr, errStr = strings.Trim(stdout.String(), "\n"), strings.Trim(stderr.String(), "\n")
	return
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
