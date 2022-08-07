package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func GitClone(originGroup string, repo string, workDir string, workBranch string) {
	pwd, _ := os.Getwd()
	remoteAddr := fmt.Sprintf("%s/%s.git", originGroup, repo)
	localDir := filepath.Join(pwd, workDir, repo)
	command := fmt.Sprintf("git clone -b %s -- %s %s", workBranch, remoteAddr, localDir)
	fmt.Printf("command: %s\n", command)

	stdout, stderr, err := Execute(command)
	if err != nil {
		fmt.Println("execute fail")
	} else {
		fmt.Println("execute succeed")
	}
	fmt.Print(stdout)
	fmt.Print(stderr)
}
