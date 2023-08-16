package git

import "github.com/chmodke/gitmm/log"
import "github.com/chmodke/gitmm/util"

// RemoteShow show remote address
func RemoteShow(localRepo string) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("remote -v")
	out, ret := util.GetOut(util.Execute(localRepo, builder.Build()))
	log.Consoleln(out)

	return ret
}

// RemoteAdd add remote address
func RemoteAdd(localRepo, remote, url string) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("remote").Add("add").Add(remote).Add(url)
	_, ret := util.GetOut(util.Execute(localRepo, builder.Build()))
	RemoteShow(localRepo)

	return ret
}

// RemoteRemove add remote address
func RemoteRemove(localRepo, remote string) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git").Add("remote").Add("remove").Add(remote)
	_, ret := util.GetOut(util.Execute(localRepo, builder.Build()))
	RemoteShow(localRepo)

	return ret
}
