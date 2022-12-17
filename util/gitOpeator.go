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

//GitClone is call git clone
func GitClone(url string, repo string, remote string, workDir string, workBranch string) bool {
	workPath, err := GetWorkDir(workDir)
	if err != nil {
		return false
	}
	localDir := filepath.Join(workPath, repo)
	remoteAddr := fmt.Sprintf("%s/%s.git", url, repo)
	log.Infof("from %s clone %s.", url, repo)
	builder := &CmdBuilder{}
	builder.Add("git").Add("clone")
	builder.Add("-o").Add(remote).Add("--")
	builder.Add(remoteAddr).Add(localDir)
	ret := Out(Execute(builder.Build()))
	if !ret {
		return ret
	}

	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("checkout").Add(workBranch)
	if !Status(Execute(builder.Build())) {
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
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("remote").Add("add upstream").Add(upstreamRemote)
	ret = Out(Execute(builder.Build()))
	if !ret {
		return ret
	}

	log.Info("2.2 fetch upstream all.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("fetch").Add("--all --prune --tags")
	ret = Out(Execute(builder.Build()))
	if !ret {
		return ret
	}

	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("branch -r")
	var out string
	out, ret = GetOut(Execute(builder.Build()))
	if !ret {
		return ret
	}

	log.Info("2.3 track all branch.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("branch -f")
	builder.Add("--track %s").Add("upstream/%s")
	tpl := builder.Build()
	for _, s := range strings.Split(out, "\n") {
		branchName, ok := getBranchName(s)
		if ok {
			command = fmt.Sprintf(tpl, branchName, branchName)
			ret = Out(Execute(command))
		}
	}

	log.Info("2.4 remove upstream fetch url.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("remote").Add("remove upstream")
	ret = Out(Execute(builder.Build()))
	if !ret {
		return ret
	}

	log.Info("3.1 add origin url.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("remote").Add("add origin").Add(originRemote)
	ret = Out(Execute(builder.Build()))
	if !ret {
		return ret
	}

	log.Info("3.2 push origin all.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("push").Add("origin --all -f")
	ret = Out(Execute(builder.Build()))
	if !ret {
		return ret
	}

	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("push").Add("origin --tags -f")
	ret = Out(Execute(builder.Build()))
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
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("symbolic-ref --short HEAD")
	branch, ret := GetOut(Execute(builder.Build()))

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("fetch --all -v")
	ret = Out(Execute(builder.Build()))
	if force {
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("reset --hard")
		builder.AddWithArg("refs/remotes/origin/%s", branch)
		ret = Out(Execute(builder.Build()))
	}
	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("pull --all -v")
	ret = Out(Execute(builder.Build()))
	return ret
}

func GitRemote(localRepo string) bool {
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("remote -v")
	out, ret := GetOut(Execute(builder.Build()))
	log.InfoO(out)

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("symbolic-ref --short HEAD")
	branch, ret := GetOut(Execute(builder.Build()))
	log.Infof("current branch %s", branch)

	return ret
}

func GitCreateBranch(localRepo, newBranch, startPoint string) bool {
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("branch").Add(newBranch).Add(startPoint)
	ret := Out(Execute(builder.Build()))

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("symbolic-ref --short HEAD")
	branch, ret := GetOut(Execute(builder.Build()))
	log.Infof("current branch %s", branch)
	return ret
}

func GitSwitchBranch(localRepo, aimBranch string, force bool) bool {
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("symbolic-ref --short HEAD")
	curBranch, ret := GetOut(Execute(builder.Build()))
	log.Infof("before switch branch %s", curBranch)

	if force {
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("clean -df")
		ret = Out(Execute(builder.Build()))
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("reset --hard")
		ret = Out(Execute(builder.Build()))
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("fetch --all")
		ret = Out(Execute(builder.Build()))
	}

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("checkout").Add(aimBranch)
	ret = Out(Execute(builder.Build()))

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("symbolic-ref --short HEAD")
	curBranch, ret = GetOut(Execute(builder.Build()))
	log.Infof("after switch branch %s", curBranch)
	return ret
}

func GitCommand(localRepo, gitCommand string) bool {
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add(gitCommand)
	out, ret := GetOut(Execute(builder.Build()))
	log.InfoO(out)
	return ret
}
