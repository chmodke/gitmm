package git

import (
	"gitmm/util"
	"strings"
)

func GitStatusStatistic(localRepo string) map[string]int {
	status := make(map[string]int)
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("status --porcelain")

	out, _ := util.GetOut(util.Execute(builder.Build()))

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
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("branch")
	builder.Add("--show-current")

	out, _ := util.GetOut(util.Execute(builder.Build()))

	if len(out) == 0 {
		builder.Reset()
		builder.Add("git")
		builder.AddWithArg("-C %s", localRepo)
		builder.Add("tag")
		builder.Add("--points-at HEAD")

		out, _ = util.GetOut(util.Execute(builder.Build()))
	}
	if len(out) == 0 {
		builder.Reset()
		builder.Add("git")
		builder.AddWithArg("-C %s", localRepo)
		builder.Add("rev-parse")
		builder.Add("--short HEAD")

		out, _ = util.GetOut(util.Execute(builder.Build()))
	}
	return out
}

func GitBranchTrack(localRepo, branchName string) string {
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("branch")
	builder.Add("-l")
	builder.Add(branchName)
	builder.Add("-v --format=%(upstream:remotename)")

	out, _ := util.GetOut(util.Execute(builder.Build()))
	return out
}

func GitLastCommit(localRepo string) string {
	builder := &util.CmdBuilder{}
	builder.Reset()
	builder.Add("git")
	builder.AddWithArg("-C %s", localRepo)
	builder.Add("log")
	builder.Add("-n1")
	builder.Add("--pretty=\"format:%ad %an\" --date=\"format:%Y/%m/%d %H:%M:%S\"")

	out, _ := util.GetOut(util.Execute(builder.Build()))
	return out
}
