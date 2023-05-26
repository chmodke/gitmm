package git

import "gitmm/util"

// GitPull is entry function
func GitPull(localRepo string, force bool, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("symbolic-ref --short HEAD")
	branch, ret := util.GetOut(util.Execute(builder.Build()))
	progress.Next()

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("fetch --all -v")
	ret = util.Status(util.Execute(builder.Build()))
	progress.Next()
	if force {
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("reset --hard")
		builder.AddWithArg("refs/remotes/origin/%s", branch)
		ret = util.Status(util.Execute(builder.Build()))
	}
	progress.Next()
	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("pull --all -v")
	ret = util.Status(util.Execute(builder.Build()))
	progress.Next()
	return ret
}
