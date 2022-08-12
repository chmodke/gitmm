// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/config"
	"gitmm/log"
	"gitmm/util"
	"os"
	"strings"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "批量同步主从仓库",
	Long: `执行脚本会读取当前目录下repo.yaml配置文件，遍历repos配置项，从main_group强制同步全部内容到origin_group中，需要用户对origin_group有强制写权限（取消分支保护）。
注意：会强制以main_group中的内容覆盖origin_group中的内容。`,
	Example: "gitmm sync",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()
		log.Debugf("main_group: %s", config.MainGroup)
		log.Debugf("origin_group: %s", config.OriginGroup)
		log.Debugf("repos: %s", config.Repos)
		for _, repo := range config.Repos {
			log.Infof("sync %s start.", repo)
			ok := util.GitSync(config.MainGroup, config.OriginGroup, repo, "tmp")
			if ok {
				log.Infof("sync %s done.", repo)
			} else {
				log.Infof("sync %s fail.", repo)
			}
			log.Info(strings.Repeat("-", 80))
		}
		os.RemoveAll("tmp")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
