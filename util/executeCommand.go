package util

import (
	"bytes"
	"os/exec"
	"runtime"
)

func Execute(command string) (outStr string, errStr string, err error) {
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
	outStr, errStr = stdout.String(), stderr.String()
	return
}
