// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "批量拉取仓库",
	Long:    `执行命令会遍历work_dir目录下中的git仓库，并执行分支拉取操作。`,
	Example: "gitmm fetch -w tmp -b master -u upstream",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		branch, _ := cmd.Flags().GetString("branch")
		remote, _ := cmd.Flags().GetString("remote")
		match, _ := cmd.Flags().GetString("match")
		invert, _ := cmd.Flags().GetString("invert-match")

		localDir, err := git.GetWorkDir(workDir)
		if err != nil {
			log.Consoleln("获取工作路径失败")
			return
		}
		repos, err := git.FindGit(localDir)
		if err != nil {
			log.Consoleln("获取本地仓库失败")
		}
		for _, repo := range repos {
			var process util.Progress
			process.NewOption(util.RightCut(repo, 18), 0, 1)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := git.Fetch(filepath.Join(localDir, repo), branch, remote, &process)
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	fetchCmd.Flags().StringP("remote", "u", "origin", "上游")
	fetchCmd.Flags().StringP("branch", "b", "master", "分支")
	fetchCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	fetchCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
