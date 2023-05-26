package git

import "gitmm/util"

func GitCreateBranch(localRepo, newBranch, startPoint string, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("branch").Add(newBranch).Add(startPoint)
	ret := util.Status(util.Execute(builder.Build()))
	progress.Next()

	return ret
}

func GitSwitchBranch(localRepo, aimBranch string, force bool, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}

	ret := false
	if force {
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("clean -df")
		ret = util.Status(util.Execute(builder.Build()))
		progress.Next()
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("reset --hard")
		ret = util.Status(util.Execute(builder.Build()))
		progress.Next()
		builder.Reset()
		builder.Add("git").Add("-C").Add(localRepo)
		builder.Add("fetch --all")
		ret = util.Status(util.Execute(builder.Build()))
		progress.Next()
	} else {
		progress.Skip(3)
	}

	builder.Reset()
	builder.Add("git").Add("-C").Add(localRepo)
	builder.Add("checkout").Add(aimBranch)
	ret = util.Status(util.Execute(builder.Build()))
	progress.Next()
	return ret
}
