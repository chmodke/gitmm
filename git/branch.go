package git

import "gitmm/util"

func GitCreateBranch(localRepo, newBranch, startPoint string, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("branch").Add(newBranch).Add(startPoint)
	ret := util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()

	return ret
}

func GitSwitchBranch(localRepo, aimBranch string, force bool, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}

	ret := false
	if force {
		builder.Reset()
		builder.Add("git").Add("clean -df")
		ret = util.Status(util.Execute(localRepo, builder.Build()))
		progress.Next()
		builder.Reset()
		builder.Add("git").Add("reset --hard")
		ret = util.Status(util.Execute(localRepo, builder.Build()))
		progress.Next()
		builder.Reset()
		builder.Add("git").Add("fetch --all")
		ret = util.Status(util.Execute(localRepo, builder.Build()))
		progress.Next()
	} else {
		progress.Skip(3)
	}

	builder.Reset()
	builder.Add("git").Add("checkout").Add(aimBranch)
	ret = util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()
	return ret
}
