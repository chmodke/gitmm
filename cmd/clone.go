// Package cmd /*
package cmd

import (
	"gitmm/config"
	"gitmm/log"
	"gitmm/util"
	"strings"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:     "clone",
	Short:   "批量克隆仓库",
	Long:    "执行脚本会读取当前目录下repo.yaml配置文件，遍历repos配置项，从origin_group克隆代码到当前目录下work_dir指定的文件夹中。",
	Example: "gitmm clone -w tmp -b master",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()

		workDir, _ := cmd.Flags().GetString("work_dir")
		workBranch, _ := cmd.Flags().GetString("work_branch")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("work_branch: %s", workBranch)
		log.Debugf("origin_group: %s", config.OriginGroup)
		log.Debugf("repos: %s", config.Repos)

		for _, repo := range config.Repos {
			ok := util.GitClone(config.OriginGroup, repo, workDir, workBranch)
			if ok {
				log.Infof("clone %s done.", repo)
			} else {
				log.Infof("clone %s fail.", repo)
			}
			log.Info(strings.Repeat("-", 80))
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringP("work_dir", "w", "master", "克隆代码的存放路径")
	cloneCmd.MarkFlagRequired("work_dir")
	cloneCmd.Flags().StringP("work_branch", "b", "master", "克隆代码的分支")
}
