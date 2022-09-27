//go:build windows

package util

import (
	"bytes"
	"gitmm/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"strings"
	"syscall"
)

type Charset string

const (
	UTF8 = Charset("UTF-8")
	GBK  = Charset("GBK")
)

func ExecShell(command string) (outStr string, errStr string, err error) {
	log.Debugf("command: %s", command)
	var cmd = exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c " + command}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	outStr = strings.Trim(ConvertByte2String(stdout.Bytes(), GBK), "\n")
	errStr = strings.Trim(ConvertByte2String(stderr.Bytes(), GBK), "\n")
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
