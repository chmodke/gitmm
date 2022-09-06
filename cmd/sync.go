// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitmm/config"
	"gitmm/log"
	"gitmm/util"
	"os"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "批量同步主从仓库",
	Long: `执行命令会读取当前目录下repo.yaml配置文件，遍历repos配置项，从main_group强制同步全部内容到origin_group中，需要用户对origin_group有强制写权限（取消分支保护）。
注意：会强制以main_group中的内容覆盖origin_group中的内容。`,
	Example: "gitmm sync",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()
		log.Debugf("main_group: %s", config.MainGroup)
		log.Debugf("origin_group: %s", config.OriginGroup)
		log.Debugf("repos: %s", config.Repos)

		grep, _ := cmd.Flags().GetString("grep")
		log.Debugf("grep: %s", grep)
		result := make(map[string]string)
		for _, repo := range config.Repos {
			if len(grep) > 0 && !util.Match(grep, repo) {
				log.Info(util.LeftAlign(fmt.Sprintf("skip %s sync.", repo), 2, "-"))
				result[repo] = SKIP
				continue
			}
			log.Info(util.LeftAlign(fmt.Sprintf("start %s sync.", repo), 2, "-"))
			ok := util.GitSync(config.MainGroup, config.OriginGroup, repo, "tmp")
			if ok {
				log.Info(util.LeftAlign(fmt.Sprintf("sync %s done.", repo), 2, "-"))
				result[repo] = OK
			} else {
				log.Error(util.LeftAlign(fmt.Sprintf("sync %s fail.", repo), 2, "-"))
				result[repo] = FAIL
			}
		}
		util.ExecStatistic("sync", result)
		os.RemoveAll("tmp")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringP("grep", "g", "", "仓库过滤条件，golang正则表达式")
}
