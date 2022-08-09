package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func Execute(command string) (outStr string, errStr string, err error) {
	fmt.Printf("command: %s\n", command)
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

func Out(stdout string, stderr string, err error) bool {
	if err != nil {
		fmt.Println("execute fail")
		fmt.Println(stderr)
		return false
	} else {
		fmt.Println("execute succeed")
		fmt.Println(stdout)
		return true
	}
}

func GetOut(stdout string, stderr string, err error) (string, bool) {
	if err != nil {
		fmt.Println("execute fail")
		fmt.Println(stderr)
		return "", false
	} else {
		fmt.Println("execute succeed")
		return stdout, true
	}
}
