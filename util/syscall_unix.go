//go:build !windows

package util

import (
	"bytes"
	"gitmm/log"
	"os/exec"
	"strings"
)

func ExecShell(command string, charset Charset) (outStr string, errStr string, err error) {
	log.Debugf("command: %s", command)
	var cmd = exec.Command("/bin/bash", "-c", command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	outStr, errStr = strings.Trim(stdout.String(), "\n"), strings.Trim(stderr.String(), "\n")
	return
}
