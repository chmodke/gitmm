// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

// remoteShowCmd represents the remote command
var remoteShowCmd = &cobra.Command{
	Use:     "show",
	Short:   "批量查看仓库远程信息",
	Long:    `执行命令遍历work_dir目录下中的git仓库，并查看仓库远程信息。`,
	Example: "gitmm remote show\n查看当前工作目录下所有仓库的远程信息",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
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
		result := make(map[string]string)
		for _, repo := range repos {
			if !util.Match(repo, match, invert) {
				result[repo] = SKIP
				continue
			}
			log.Consoleln(repo)
			ok := git.RemoteShow(filepath.Join(localDir, repo))
			log.Consoleln("")
			if ok {
				result[repo] = OK
			} else {
				result[repo] = FAIL
			}
		}
		util.ExecStatistic("remote", result)
	},
}

func init() {
	remoteCmd.AddCommand(remoteShowCmd)

	remoteShowCmd.Flags().StringP("work_dir", "w", ".", "可选，本地代码的存放路径")
	remoteShowCmd.Flags().StringP("match", "m", "", "可选，仓库过滤条件，golang正则表达式")
	remoteShowCmd.Flags().StringP("invert-match", "i", "", "可选，仓库反向过滤条件，golang正则表达式")
}
