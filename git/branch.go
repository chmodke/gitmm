package git

import (
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"strings"
)

func CreateBranch(localRepo, newBranch, startPoint string, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("branch").Add(newBranch).Add(startPoint)
	ret := util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()

	return ret
}

func SwitchBranch(localRepo, aimBranch string, force bool, progress *util.Progress) bool {
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

func DeleteBranch(localRepo, branch string, force bool, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("branch").Add("-d")
	if force {
		builder.Add("-D").Add("-f")
	}
	builder.Add(branch)
	ret := util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()

	return ret
}

func RenameBranch(localRepo, oldBranch, newBranch string, progress *util.Progress) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("branch").Add("-m").Add("-M").Add(oldBranch).Add(newBranch)
	ret := util.Status(util.Execute(localRepo, builder.Build()))
	progress.Next()

	return ret
}

func ListBranch(localRepo string) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("branch").Add("-l").Add("--format=%(refname:lstrip=2)")
	out, ret := util.GetOut(util.Execute(localRepo, builder.Build()))
	log.Consoleln(strings.Split(out, "\n"))
	return ret
}
