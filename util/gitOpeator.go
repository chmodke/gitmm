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

func GetWorkDir(workDir string) (string, error) {
	var localDir string
	if filepath.IsAbs(workDir) {
		localDir = workDir
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		localDir = filepath.Join(pwd, workDir)
	}

	if PathExists(localDir) {
		return localDir, nil
	}

	err := os.MkdirAll(localDir, fs.ModeDir)
	if err != nil {
		return "", err
	}
	return localDir, nil
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

func GitClone(origin string, repo string, workDir string, workBranch string) bool {
	workPath, err := GetWorkDir(workDir)
	if err != nil {
		return false
	}
	localDir := filepath.Join(workPath, repo)
	remoteAddr := fmt.Sprintf("%s/%s.git", origin, repo)
	log.Infof("from %s clone %s.", origin, repo)
	command := fmt.Sprintf("git clone -- %s %s", remoteAddr, localDir)
	ret := Out(Execute(command))
	if !ret {
		return ret
	}

	command = fmt.Sprintf("git -C %s checkout %s", localDir, workBranch)
	if !Status(Execute(command)) {
		log.Warnf("switch to %s fail.", workBranch)
	}
	return ret
}

func GitSync(upstream string, origin string, repo string, workDir string) bool {
	upstreamRemote := fmt.Sprintf("%s/%s.git", upstream, repo)
	originRemote := fmt.Sprintf("%s/%s.git", origin, repo)

	workPath, err := GetWorkDir(workDir)
	if err != nil {
		return false
	}
	localDir := filepath.Join(workPath, repo)

	var command string
	var ret bool

	log.Infof("sync %s, from %s to %s.", repo, upstream, origin)
	log.Info("1.1 init local repo.")
	command = fmt.Sprintf("git init %s", localDir)
	ret = Out(Execute(command))
	if !ret {
		return ret
	}

	log.Info("2.1 add upstream fetch url.")
	command = fmt.Sprintf("git -C %s remote add upstream %s", localDir, upstreamRemote)
	ret = Out(Execute(command))
	if !ret {
		return ret
	}

	log.Info("2.2 fetch upstream all.")
	command = fmt.Sprintf("git -C %s fetch --all --prune --tags", localDir)
	ret = Out(Execute(command))
	if !ret {
		return ret
	}

	command = fmt.Sprintf("git -C %s branch -r", localDir)
	var out string
	out, ret = GetOut(Execute(command))
	if !ret {
		return ret
	}

	log.Info("2.3 track all branch.")
	for _, s := range strings.Split(out, "\n") {
		branchName, ok := getBranchName(s)
		if ok {
			command = fmt.Sprintf("git -C %s branch -f --track %s upstream/%s", localDir, branchName, branchName)
			ret = Out(Execute(command))
		}
	}

	log.Info("2.4 remove upstream fetch url.")
	command = fmt.Sprintf("git -C %s remote remove upstream", localDir)
	ret = Out(Execute(command))
	if !ret {
		return ret
	}

	log.Info("3.1 add origin url.")
	command = fmt.Sprintf("git -C %s remote add origin %s", localDir, originRemote)
	ret = Out(Execute(command))
	if !ret {
		return ret
	}

	log.Info("3.2 push origin all.")
	command = fmt.Sprintf("git -C %s push origin --all -f", localDir)
	ret = Out(Execute(command))
	if !ret {
		return ret
	}

	command = fmt.Sprintf("git -C %s push origin --tags -f", localDir)
	ret = Out(Execute(command))
	return ret
}

/*
getBranchName 从upstream/master获取分支的名称master
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
	log.InfoO(out)

	command = fmt.Sprintf("git -C %s symbolic-ref --short HEAD", localRepo)
	branch, ret := GetOut(Execute(command))
	log.Infof("current branch %s", branch)

	return ret
}

func GitCreateBranch(localRepo, newBranch, startPoint string) bool {
	var command string

	command = fmt.Sprintf("git -C %s branch %s %s", localRepo, newBranch, startPoint)
	ret := Out(Execute(command))

	command = fmt.Sprintf("git -C %s symbolic-ref --short HEAD", localRepo)
	branch, ret := GetOut(Execute(command))
	log.Infof("current branch %s", branch)
	return ret
}

func GitSwitchBranch(localRepo, aimBranch string, force bool) bool {
	var command string

	command = fmt.Sprintf("git -C %s symbolic-ref --short HEAD", localRepo)
	curBranch, ret := GetOut(Execute(command))
	log.Infof("before switch branch %s", curBranch)

	if force {
		command = fmt.Sprintf("git -C %s clean -df", localRepo)
		ret = Out(Execute(command))
		command = fmt.Sprintf("git -C %s reset --hard", localRepo)
		ret = Out(Execute(command))
		command = fmt.Sprintf("git -C %s fetch --all", localRepo)
		ret = Out(Execute(command))
	}

	command = fmt.Sprintf("git -C %s checkout %s", localRepo, aimBranch)
	ret = Out(Execute(command))

	command = fmt.Sprintf("git -C %s symbolic-ref --short HEAD", localRepo)
	curBranch, ret = GetOut(Execute(command))
	log.Infof("after switch branch %s", curBranch)
	return ret
}

func GitCommand(localRepo, gitCommand string) bool {
	var command string

	command = fmt.Sprintf("git -C %s %s", localRepo, gitCommand)
	out, ret := GetOut(Execute(command))
	log.InfoO(out)
	return ret
}
