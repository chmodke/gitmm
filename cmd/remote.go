// Package cmd /*
package cmd

import (
	"fmt"
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

		localDir, err := util.GetWorkDir(workDir)
		if err != nil {
			log.Error("获取工作路径失败")
			return
		}
		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}
		for _, repo := range repos {
			log.Info(util.Title(fmt.Sprintf("get %s remote info.", repo), 80, "-"))
			ok := util.GitRemote(filepath.Join(localDir, repo))
			if ok {
				log.Info(util.Title(fmt.Sprintf("show remote %s done.", repo), 80, "-"))
			} else {
				log.Error(util.Title(fmt.Sprintf("show remote %s fail.", repo), 80, "-"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	remoteCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
}
