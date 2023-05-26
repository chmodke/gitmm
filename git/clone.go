package git

import (
	"fmt"
	"gitmm/util"
	"path/filepath"
)

//GitClone is entry function
func GitClone(url string, repo string, remote string, workDir string, workBranch string, progress *util.Progress) bool {
	workPath, err := GetWorkDir(workDir)
	if err != nil {
		return false
	}
	localDir := filepath.Join(workPath, repo)
	remoteAddr := fmt.Sprintf("%s/%s.git", url, repo)
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("clone")
	builder.Add("-o").Add(remote).Add("--")
	builder.Add(remoteAddr).Add(localDir)
	ret := util.Status(util.Execute(builder.Build()))
	if !ret {
		return ret
	}
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("-C").Add(localDir)
	builder.Add("checkout").Add(workBranch)
	ret = util.Status(util.Execute(builder.Build()))
	progress.Next()
	return ret
}
