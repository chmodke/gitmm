package git

import "gitmm/util"

// GitPull is entry function
func GitPull(localRepo string, force bool, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("symbolic-ref --short HEAD")
	branch, ret := util.GetOut(util.Execute(localRepo, builder.Build()))
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("fetch --all -v")
	ret = util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()
	if force {
		builder.Reset()
		builder.Add("git").Add("reset --hard")
		builder.AddWithArg("refs/remotes/origin/%s", branch)
		ret = util.Status(util.Execute(localRepo, builder.Build()))
	}
	progress.Next()
	builder.Reset()
	builder.Add("git").Add("pull --all -v")
	ret = util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()
	return ret
}
