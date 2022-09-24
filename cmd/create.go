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
		log.Debugf("work_dir: %s", workDir)
		newBranch, _ := cmd.Flags().GetString("new_branch")
		log.Debugf("new_branch: %s", newBranch)
		startPoint, _ := cmd.Flags().GetString("refs")
		log.Debugf("refs: %s", startPoint)
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
				log.Info(util.LeftAlign(fmt.Sprintf("skip create branch at %s.\n", repo), 2, "-"))
				result[repo] = SKIP
				continue
			}
			log.Info(util.LeftAlign(fmt.Sprintf("start create branch at %s.", repo), 2, "-"))
			ok := util.GitCreateBranch(filepath.Join(localDir, repo), newBranch, startPoint)
			if ok {
				log.Info(util.LeftAlign(fmt.Sprintf("%s create branch done.\n", repo), 2, "-"))
				result[repo] = OK
			} else {
				log.Error(util.LeftAlign(fmt.Sprintf("%s create branch fail.\n", repo), 2, "-"))
				result[repo] = FAIL
			}
		}
		util.ExecStatistic("create", result)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	createCmd.Flags().StringP("new_branch", "b", "master", "新分支名称")
	createCmd.MarkFlagRequired("new_branch")
	createCmd.Flags().StringP("refs", "r", "HEAD", "新分支起点")
	createCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	createCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
