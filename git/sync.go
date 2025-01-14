package git

import (
	"fmt"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"path/filepath"
	"strings"
)

// Sync is entry function
func Sync(upstream string, origin string, repo string, workDir string, progress *util.Progress) bool {
	upstreamRemote := fmt.Sprintf("%s/%s.git", upstream, repo)
	originRemote := fmt.Sprintf("%s/%s.git", origin, repo)

	workPath, err := GetWorkDir(workDir)
	if err != nil {
		return false
	}
	localDir := filepath.Join(workPath, repo)

	var command string
	var ret bool

	log.Printf("sync %s, from %s to %s.", repo, upstream, origin)
	log.Println("1.1 init local repo.")
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("init").Add(repo)
	ret = util.Status(util.Execute(workPath, builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Println("2.1 add upstream fetch url.")
	builder.Reset()
	builder.Add("git").Add("remote").Add("add upstream").Add(upstreamRemote)
	ret = util.Status(util.Execute(localDir, builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Println("2.2 fetch upstream all.")
	builder.Reset()
	builder.Add("git").Add("fetch").Add("--all --prune --tags")
	ret = util.Status(util.Execute(localDir, builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("branch -r")
	var out string
	out, ret = util.GetOut(util.Execute(localDir, builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Println("2.3 track all branch.")
	builder.Reset()
	builder.Add("git").Add("branch -f")
	builder.Add("--track %s").Add("upstream/%s")
	tpl := builder.Build()
	for _, s := range strings.Split(out, "\n") {
		branchName, ok := getBranchName(s)
		if ok {
			command = fmt.Sprintf(tpl, branchName, branchName)
			ret = util.Status(util.Execute(localDir, command))
		}
	}
	progress.Next()

	log.Println("2.4 remove upstream fetch url.")
	builder.Reset()
	builder.Add("git").Add("remote").Add("remove upstream")
	ret = util.Status(util.Execute(localDir, builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Println("3.1 add origin url.")
	builder.Reset()
	builder.Add("git").Add("remote").Add("add origin").Add(originRemote)
	ret = util.Status(util.Execute(localDir, builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	log.Println("3.2 push origin all.")
	builder.Reset()
	builder.Add("git").Add("push").Add("origin --all -f")
	ret = util.Status(util.Execute(localDir, builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("push").Add("origin --tags -f")
	ret = util.Status(util.Execute(localDir, builder.Build()))
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
