package git

import (
	"github.com/chmodke/gitmm/util"
	"strings"
)

func StatusStatistic(localRepo string) map[string]int {
	status := make(map[string]int)
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git").Add("status --porcelain")

	out, _ := util.GetOut(util.Execute(localRepo, builder.Build()))

	outData := strings.Split(out, "\n")
	for _, line := range outData {
		if len(line) > 0 {
			s := strings.Split(strings.TrimLeft(line, " "), " ")[0]
			status[s]++
		}
	}
	return status
}

func CurrentBranch(localRepo string) string {
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git").Add("branch")
	builder.Add("--show-current")

	out, _ := util.GetOut(util.Execute(localRepo, builder.Build()))

	if len(out) == 0 {
		builder.Reset()
		builder.Add("git").Add("tag")
		builder.Add("--points-at HEAD")

		out, _ = util.GetOut(util.Execute(localRepo, builder.Build()))
	}
	if len(out) == 0 {
		builder.Reset()
		builder.Add("git").Add("rev-parse")
		builder.Add("--short HEAD")

		out, _ = util.GetOut(util.Execute(localRepo, builder.Build()))
	}
	return out
}

func BranchTrack(localRepo, branchName string) string {
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git").Add("branch")
	builder.Add("-l")
	builder.Add(branchName)
	builder.Add("-v --format=%(upstream:lstrip=2)")

	out, _ := util.GetOut(util.Execute(localRepo, builder.Build()))
	return out
}

func LastCommit(localRepo string) string {
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git").Add("log")
	builder.Add("-n1")
	builder.Add("--pretty=\"format:%ad %an\" --date=\"format:%Y/%m/%d %H:%M:%S\"")

	out, _ := util.GetOut(util.Execute(localRepo, builder.Build()))
	return out
}
