package git

import (
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"os"
	"runtime"
)

func GetGitVersion() string {
	command := "git --version"
	var charset = util.UTF8
	if runtime.GOOS == "windows" {
		charset = util.GBK
	}
	out, ok := util.GetOut(util.ExecuteWithCharset("", command, charset))
	if !ok {
		log.Consoleln("执行git失败，请检查是否安装git，或者环境变量配置错误。")
		log.Consoleln("下载地址: <https://repo.huaweicloud.com/git-for-windows/>")
		os.Exit(1)
		return ""
	} else {
		return out
	}
}
