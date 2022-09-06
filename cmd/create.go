// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "批量创建分支",
	Long:    `执行命令会遍历work_dir中的git仓库，并执行分支创建操作。`,
	Example: "git create -w tmp -b develop",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		newBranch, _ := cmd.Flags().GetString("new_branch")
		startPoint, _ := cmd.Flags().GetString("refs")
		grep, _ := cmd.Flags().GetString("grep")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("new_branch: %s", newBranch)
		log.Debugf("refs: %s", startPoint)
		log.Debugf("grep: %s", grep)

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
			if len(grep) > 0 && !util.Match(grep, repo) {
				log.Info(util.LeftAlign(fmt.Sprintf("skip create branch at %s.", repo), 2, "-"))
				continue
			}
			log.Info(util.LeftAlign(fmt.Sprintf("start create branch at %s.", repo), 2, "-"))
			ok := util.GitCreateBranch(filepath.Join(localDir, repo), newBranch, startPoint)
			if ok {
				log.Info(util.LeftAlign(fmt.Sprintf("%s create branch done.", repo), 2, "-"))
			} else {
				log.Error(util.LeftAlign(fmt.Sprintf("%s create branch fail.", repo), 2, "-"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	createCmd.Flags().StringP("new_branch", "b", "master", "新分支名称")
	createCmd.MarkFlagRequired("new_branch")
	createCmd.Flags().StringP("refs", "r", "HEAD", "新分支起点")
	createCmd.Flags().StringP("grep", "g", "", "仓库过滤条件")
}
