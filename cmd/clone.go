// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitmm/config"
	"gitmm/log"
	"gitmm/util"
)

var cloneCmd = &cobra.Command{
	Use:     "clone",
	Short:   "批量克隆仓库",
	Long:    "执行命令会读取当前目录下repo.yaml配置文件，遍历repos配置项，从origin克隆代码到当前目录下work_dir指定的文件夹中。",
	Example: "gitmm clone -w tmp -b master",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()
		log.Debugf("origin: %s", config.Origin)
		log.Debugf("repos: %s", config.Repos)

		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Debugf("work_dir: %s", workDir)
		workBranch, _ := cmd.Flags().GetString("work_branch")
		log.Debugf("work_branch: %s", workBranch)
		match, _ := cmd.Flags().GetString("match")
		log.Debugf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Debugf("invert: %s", invert)

		result := make(map[string]string)
		for _, repo := range config.Repos {
			if !util.Match(repo, match, invert) {
				log.Info(util.LeftAlign(fmt.Sprintf("skip clone %s.\n", repo), 2, "-"))
				result[repo] = SKIP
				continue
			}
			log.Info(util.LeftAlign(fmt.Sprintf("start clone %s.", repo), 2, "-"))
			ok := util.GitClone(config.Origin, repo, workDir, workBranch)
			if ok {
				log.Info(util.LeftAlign(fmt.Sprintf("clone %s done.\n", repo), 2, "-"))
				result[repo] = OK
			} else {
				log.Error(util.LeftAlign(fmt.Sprintf("clone %s fail.\n", repo), 2, "-"))
				result[repo] = FAIL
			}
		}
		util.ExecStatistic("clone", result)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringP("work_dir", "w", "master", "克隆代码的存放路径")
	cloneCmd.MarkFlagRequired("work_dir")
	cloneCmd.Flags().StringP("work_branch", "b", "master", "克隆代码的分支")
	cloneCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	cloneCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
