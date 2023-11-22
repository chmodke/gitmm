package git

import "github.com/chmodke/gitmm/util"

func Fetch(localRepo, branch, remote string, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("fetch").Add(remote).Add(branch)
	ret := util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()
	return ret
}
