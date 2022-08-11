package util

import (
	"fmt"
	"gitmm/log"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetWorkDir(workDir string) string {
	var localDir string
	if filepath.IsAbs(workDir) {
		localDir = workDir
	} else {
		pwd, _ := os.Getwd()
		localDir = filepath.Join(pwd, workDir)
	}

	if PathExists(localDir) {
		return localDir
	}

	os.MkdirAll(localDir, fs.ModeDir)
	return localDir
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GitClone(originGroup string, repo string, workDir string, workBranch string) bool {
	localDir := filepath.Join(GetWorkDir(workDir), repo)
	remoteAddr := fmt.Sprintf("%s/%s.git", originGroup, repo)
	command := fmt.Sprintf("git clone -b %s -- %s %s", workBranch, remoteAddr, localDir)
	return Out(Execute(command))
}

func GitSync(mainGroup string, originGroup string, repo string, workDir string) bool {
	mainRemote := fmt.Sprintf("%s/%s.git", mainGroup, repo)
	originRemote := fmt.Sprintf("%s/%s.git", originGroup, repo)
	localDir := filepath.Join(GetWorkDir(workDir), repo)

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

func FindGit(dirPth string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		if isGit(filepath.Join(dirPth, fi.Name())) {
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

func isGit(dirPth string) bool {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return false
	}
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		if strings.EqualFold(".git", fi.Name()) {
			return true
		}
	}
	return false
}

func GitPull(localRepo string, force bool) bool {
	var command string

	command = fmt.Sprintf("git -C %s symbolic-ref --short HEAD", localRepo)
	branch, ret := GetOut(Execute(command))

	command = fmt.Sprintf("git -C %s fetch --all -v", localRepo)
	ret = Out(Execute(command))
	if force {
		command = fmt.Sprintf("git -C %s reset --hard refs/remotes/origin/%s", localRepo, branch)
		ret = Out(Execute(command))
	}
	command = fmt.Sprintf("git -C %s pull --all -v", localRepo)
	ret = Out(Execute(command))
	return ret
}

func GitRemote(localRepo string) bool {
	var command string

	command = fmt.Sprintf("git -C %s remote -v", localRepo)
	out, ret := GetOut(Execute(command))
	log.Info(out)

	command = fmt.Sprintf("git -C %s symbolic-ref --short HEAD", localRepo)
	branch, ret := GetOut(Execute(command))
	log.Infof("current branch %s", branch)

	return ret
}
