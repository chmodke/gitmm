// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

// remoteRemoveCmd represents the remote command
var remoteRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "批量删除仓库远程信息",
	Long:    `执行命令会遍历work_dir目录下中的git仓库，并删除仓库远程信息。`,
	Example: "gitmm remote remove upstream",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		remote := args[0]
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
			ok := git.RemoteRemove(filepath.Join(localDir, repo), remote)
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
	remoteCmd.AddCommand(remoteRemoveCmd)

	remoteRemoveCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	remoteRemoveCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	remoteRemoveCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
