//go:build !windows

package util

import (
	"bytes"
	"github.com/chmodke/gitmm/log"
	"os/exec"
	"strings"
)

func ExecShell(workDir string, command string, charset Charset) (outStr string, errStr string, err error) {
	log.Printf("exec command: [%s] at [%s].", command, workDir)
	var cmd = exec.Command("/bin/bash", "-c", command)
	cmd.Dir = workDir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	outStr, errStr = strings.Trim(stdout.String(), "\n"), strings.Trim(stderr.String(), "\n")
	return
}
