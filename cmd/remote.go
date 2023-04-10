// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:     "remote",
	Short:   "批量查看仓库远程信息",
	Long:    `执行命令会遍历work_dir目录下中的git仓库，并查看仓库远程信息。`,
	Example: "gitmm remote -w tmp",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Debugf("work_dir: %s", workDir)
		match, _ := cmd.Flags().GetString("match")
		log.Debugf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Debugf("invert: %s", invert)

		localDir, err := util.GetWorkDir(workDir)
		if err != nil {
			log.Error("获取工作路径失败")
			return
		}
		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}
		result := make(map[string]string)
		for _, repo := range repos {
			if !util.Match(repo, match, invert) {
				result[repo] = SKIP
				continue
			}
			log.InfoO(repo)
			ok := util.GitRemote(filepath.Join(localDir, repo))
			log.InfoO("")
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
	rootCmd.AddCommand(remoteCmd)

	remoteCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	remoteCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	remoteCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
