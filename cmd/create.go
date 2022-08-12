// Package cmd /*
package cmd

import (
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "批量创建分支",
	Long:    `执行脚本会遍历work_dir中的git仓库，并执行分支创建操作。`,
	Example: "git create -w tmp -b develop",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		newBranch, _ := cmd.Flags().GetString("new_branch")
		startPoint, _ := cmd.Flags().GetString("refs")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("new_branch: %s", newBranch)
		log.Debugf("refs: %s", startPoint)

		localDir := util.GetWorkDir(workDir)
		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}
		for _, repo := range repos {
			ok := util.GitCreateBranch(filepath.Join(localDir, repo), newBranch, startPoint)
			if ok {
				log.Infof("%s create branch done.", repo)
			} else {
				log.Infof("%s create branch fail.", repo)
			}
			log.Info(strings.Repeat("-", 80))
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("work_dir", "w", "master", "本地代码的存放路径")
	createCmd.MarkFlagRequired("work_dir")
	createCmd.Flags().StringP("new_branch", "b", "master", "新分支名称")
	createCmd.MarkFlagRequired("new_branch")
	createCmd.Flags().StringP("refs", "r", "HEAD", "新分支起点")
}
