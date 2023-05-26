package git

import "gitmm/log"
import "gitmm/util"

func GitRemote(localRepo string) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("remote -v")
	out, ret := util.GetOut(util.Execute(localRepo, builder.Build()))
	log.Consoleln(out)

	return ret
}
