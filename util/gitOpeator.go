package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GitClone(originGroup string, repo string, workDir string, workBranch string) bool {
	pwd, _ := os.Getwd()
	remoteAddr := fmt.Sprintf("%s/%s.git", originGroup, repo)
	localDir := filepath.Join(pwd, workDir, repo)
	command := fmt.Sprintf("git clone -b %s -- %s %s", workBranch, remoteAddr, localDir)
	return Out(Execute(command))
}

func GitSync(mainGroup string, originGroup string, repo string, workDir string) bool {
	pwd, _ := os.Getwd()
	mainRemote := fmt.Sprintf("%s/%s.git", mainGroup, repo)
	originRemote := fmt.Sprintf("%s/%s.git", originGroup, repo)
	localDir := filepath.Join(pwd, workDir, repo)

	var command string
	var ret bool

	command = fmt.Sprintf("git init -b %s %s", "master", localDir)
	ret = Out(Execute(command))

	command = fmt.Sprintf("git -C %s remote add main %s", localDir, mainRemote)
	ret = Out(Execute(command))

	command = fmt.Sprintf("git -C %s fetch --all --prune --tags", localDir)
	ret = Out(Execute(command))

	command = fmt.Sprintf("git -C %s branch -r", localDir)
	var out string
	out, ret = GetOut(Execute(command))

	if !ret {
		return ret
	}
	for _, s := range strings.Split(out, "\n") {
		branchName, ok := getBranchName(s)
		if ok {
			command = fmt.Sprintf("git -C %s branch -f --track %s main/%s", localDir, branchName, branchName)
			ret = Out(Execute(command))
		}
	}

	command = fmt.Sprintf("git -C %s remote remove main", localDir)
	ret = Out(Execute(command))

	command = fmt.Sprintf("git -C %s remote add origin %s", localDir, originRemote)
	ret = Out(Execute(command))

	command = fmt.Sprintf("git -C %s push origin --all -f", localDir)
	ret = Out(Execute(command))

	command = fmt.Sprintf("git -C %s push origin --tags -f", localDir)
	ret = Out(Execute(command))
	return ret
}

/*
getBranchName 从main/master获取分支的名称master
*/
func getBranchName(str string) (string, bool) {
	realStr := strings.Trim(str, " ")
	if len(realStr) == 0 {
		return "", false
	}
	if !strings.Contains(realStr, "/") {
		return realStr, false
	}
	return strings.Split(realStr, "/")[1], true
}
