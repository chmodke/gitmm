// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

// branchCreateCmd represents the create command
var branchCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "批量创建分支",
	Long:    `执行命令会遍历work_dir中的git仓库，并执行分支创建操作。`,
	Example: "git branch create -w tmp -b develop",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Printf("work_dir: %s", workDir)
		newBranch, _ := cmd.Flags().GetString("new_branch")
		log.Printf("new_branch: %s", newBranch)
		startPoint, _ := cmd.Flags().GetString("refs")
		log.Printf("refs: %s", startPoint)
		match, _ := cmd.Flags().GetString("match")
		log.Printf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Printf("invert: %s", invert)

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
			ok := git.CreateBranch(filepath.Join(localDir, repo), newBranch, startPoint, &process)
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
	},
}

func init() {
	branchCmd.AddCommand(branchCreateCmd)

	branchCreateCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	branchCreateCmd.Flags().StringP("new_branch", "b", "master", "新分支名称")
	branchCreateCmd.MarkFlagRequired("new_branch")
	branchCreateCmd.Flags().StringP("refs", "r", "HEAD", "新分支起点")
	branchCreateCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	branchCreateCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
