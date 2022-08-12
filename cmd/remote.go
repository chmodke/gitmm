// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
	"strings"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:     "remote",
	Short:   "批量查看仓库远程信息",
	Long:    `执行脚本会遍历work_dir目录下中的git仓库，并查看仓库远程信息。`,
	Example: "gitmm remote -w tmp",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Debugf("work_dir: %s", workDir)

		localDir := util.GetWorkDir(workDir)
		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}
		for _, repo := range repos {
			ok := util.GitRemote(filepath.Join(localDir, repo))
			if ok {
				log.Infof("show remote %s done.", repo)
			} else {
				log.Infof("show remote %s fail.", repo)
			}
			log.Info(strings.Repeat("-", 80))
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	remoteCmd.Flags().StringP("work_dir", "w", "master", "本地代码的存放路径")
	remoteCmd.MarkFlagRequired("work_dir")
}
