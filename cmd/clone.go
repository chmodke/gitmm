// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/config"
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"os"
)

var cloneCmd = &cobra.Command{
	Use:     "clone",
	Short:   "批量克隆仓库",
	Long:    "执行命令会读取当前目录下repo.yaml配置文件，遍历repos配置项，从origin克隆代码到当前目录下work_dir指定的文件夹中。",
	Example: "gitmm clone -w tmp -b master",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()

		workDir, _ := cmd.Flags().GetString("work_dir")
		workBranch, _ := cmd.Flags().GetString("branch")
		remote, _ := cmd.Flags().GetString("remote")
		match, _ := cmd.Flags().GetString("match")
		invert, _ := cmd.Flags().GetString("invert-match")

		url, ok := config.Remote[remote]
		if !ok {
			log.Consolef("未配置%s远端地址\n", remote)
			os.Exit(1)
		}

		log.Printf("remote-url: %s", url)
		log.Printf("repos: %s", config.Repos)

		for _, repo := range config.Repos {
			var process util.Progress
			process.NewOption(util.RightCut(repo, 18), 0, 2)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := git.Clone(url, repo, remote, workDir, workBranch, &process)
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringP("work_dir", "w", "master", "克隆代码的存放路径")
	cloneCmd.Flags().StringP("branch", "b", "master", "克隆代码的分支")
	cloneCmd.Flags().StringP("remote", "u", "origin", "克隆代码的远程名称")
	cloneCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	cloneCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
