package util

import (
	"bytes"
	"gitmm/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"runtime"
	"strings"
)

type Charset string

const (
	UTF8 = Charset("UTF-8")
	GBK  = Charset("GBK")
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

	if runtime.GOOS == "windows" {
		outStr = strings.Trim(ConvertByte2String(stdout.Bytes(), GBK), "\n")
		errStr = strings.Trim(ConvertByte2String(stderr.Bytes(), GBK), "\n")
	} else {
		outStr, errStr = strings.Trim(stdout.String(), "\n"), strings.Trim(stderr.String(), "\n")
	}
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
