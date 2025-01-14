//go:build windows

package util

import (
	"bytes"
	"github.com/chmodke/gitmm/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"strings"
	"syscall"
)

func ExecShell(workDir string, command string, charset Charset) (outStr string, errStr string, err error) {
	log.Printf("exec command: [%s] at [%s].", command, workDir)
	var cmd = exec.Command("cmd.exe")
	cmd.Dir = workDir
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c " + command}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	outStr = strings.Trim(ConvertByte2String(stdout.Bytes(), charset), "\n")
	errStr = strings.Trim(ConvertByte2String(stderr.Bytes(), charset), "\n")
	return
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GBK:
		var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
