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

//GitClone is entry function
func GitClone(url string, repo string, remote string, workDir string, workBranch string, progress *Progress) bool {
	workPath, err := GetWorkDir(workDir)
	if err != nil {
		return false
	}
	localDir := filepath.Join(workPath, repo)
	remoteAddr := fmt.Sprintf("%s/%s.git", url, repo)
	builder := &CmdBuilder{}
	builder.Add("git").Add("clone")
	builder.Add("-o").Add(remote).Add("--")
	builder.Add(remoteAddr).Add(localDir)
	ret := Status(Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("checkout").Add(workBranch)
	ret = Status(Execute(builder.Build()))
	progress.Next()
	return ret
}

// GitSync is entry function
func GitSync(upstream string, origin string, repo string, workDir string, progress *Progress) bool {
	upstreamRemote := fmt.Sprintf("%s/%s.git", upstream, repo)
	originRemote := fmt.Sprintf("%s/%s.git", origin, repo)

	workPath, err := GetWorkDir(workDir)
	if err != nil {
		return false
	}
	localDir := filepath.Join(workPath, repo)

	var command string
	var ret bool

	log.Debugf("sync %s, from %s to %s.", repo, upstream, origin)
	log.Debug("1.1 init local repo.")
	command = fmt.Sprintf("git init %s", localDir)
	ret = Status(Execute(command))
	if !ret {
		return ret
	}
	progress.Next()

	log.Debug("2.1 add upstream fetch url.")
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("remote").Add("add upstream").Add(upstreamRemote)
	ret = Status(Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Debug("2.2 fetch upstream all.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("fetch").Add("--all --prune --tags")
	ret = Status(Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("branch -r")
	var out string
	out, ret = GetOut(Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Debug("2.3 track all branch.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("branch -f")
	builder.Add("--track %s").Add("upstream/%s")
	tpl := builder.Build()
	for _, s := range strings.Split(out, "\n") {
		branchName, ok := getBranchName(s)
		if ok {
			command = fmt.Sprintf(tpl, branchName, branchName)
			ret = Status(Execute(command))
		}
	}
	progress.Next()

	log.Debug("2.4 remove upstream fetch url.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("remote").Add("remove upstream")
	ret = Status(Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Debug("3.1 add origin url.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("remote").Add("add origin").Add(originRemote)
	ret = Status(Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Debug("3.2 push origin all.")
	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("push").Add("origin --all -f")
	ret = Status(Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("push").Add("origin --tags -f")
	ret = Status(Execute(builder.Build()))
	progress.Next()
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

// GitPull is entry function
func GitPull(localRepo string, force bool, progress *Progress) bool {
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("symbolic-ref --short HEAD")
	branch, ret := GetOut(Execute(builder.Build()))
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("fetch --all -v")
	ret = Status(Execute(builder.Build()))
	progress.Next()
	if force {
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("reset --hard")
		builder.AddWithArg("refs/remotes/origin/%s", branch)
		ret = Status(Execute(builder.Build()))
	}
	progress.Next()
	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("pull --all -v")
	ret = Status(Execute(builder.Build()))
	progress.Next()
	return ret
}

func GitRemote(localRepo string) bool {
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("remote -v")
	out, ret := GetOut(Execute(builder.Build()))
	log.InfoO(out)

	return ret
}

func GitCreateBranch(localRepo, newBranch, startPoint string, progress *Progress) bool {
	builder := &CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("branch").Add(newBranch).Add(startPoint)
	ret := Status(Execute(builder.Build()))
	progress.Next()

	return ret
}

func GitSwitchBranch(localRepo, aimBranch string, force bool, progress *Progress) bool {
	builder := &CmdBuilder{}

	ret := false
	if force {
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("clean -df")
		ret = Status(Execute(builder.Build()))
		progress.Next()
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("reset --hard")
		ret = Status(Execute(builder.Build()))
		progress.Next()
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("fetch --all")
		ret = Status(Execute(builder.Build()))
		progress.Next()
	} else {
		progress.Skip(3)
	}

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("checkout").Add(aimBranch)
	ret = Status(Execute(builder.Build()))
	progress.Next()
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

func GitStatusStatistic(localRepo string) map[string]int {
	status := make(map[string]int)
	builder := &CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("status --porcelain")

	out, _ := GetOut(Execute(builder.Build()))

	outData := strings.Split(out, "\n")
	for _, line := range outData {
		if len(line) > 0 {
			s := strings.Split(strings.TrimLeft(line, " "), " ")[0]
			status[s]++
		}
	}
	return status
}

func GitCurrentBranch(localRepo string) string {
	builder := &CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("branch")
	builder.Add("--show-current")

	out, _ := GetOut(Execute(builder.Build()))
	return out
}

func GitBranchTrack(localRepo, branchName string) string {
	builder := &CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("branch")
	builder.Add("-l")
	builder.Add(branchName)
	builder.Add("-v --format=%(upstream:remotename)")

	out, _ := GetOut(Execute(builder.Build()))
	return out
}

func GitLastCommit(localRepo string) string {
	builder := &CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("log")
	builder.Add("-n1")
	builder.Add("--pretty=\"format:%ad %an\" --date=\"format:%Y/%m/%d %H:%M:%S\"")

	out, _ := GetOut(Execute(builder.Build()))
	return out
}
