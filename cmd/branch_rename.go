// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

// branchRenameCmd represents the switch command
var branchRenameCmd = &cobra.Command{
	Use:     "rename [flags] old_branch new_branch",
	Short:   "批量重命名分支",
	Long:    `执行命令遍历work_dir中的git仓库，并执行分支重命名操作。`,
	Example: "gitmm branch rename develop master\n将当前工作目录下所有仓库的develop分支重命名master分支",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		branches := args[0]
		match, _ := cmd.Flags().GetString("match")
		invert, _ := cmd.Flags().GetString("invert-match")
		newBranch := args[1]

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
			process.NewOption(util.RightCut(repo, 18), 0, 4)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := true
			for _, branch := range strings.Split(branches, ",") {
				ok = git.RenameBranch(filepath.Join(localDir, repo), branch, newBranch, &process)
				if ok {
					break
				}
			}
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
	},
}

func init() {
	branchCmd.AddCommand(branchRenameCmd)

	branchRenameCmd.Flags().StringP("work_dir", "w", ".", "可选，本地代码的存放路径")
	branchRenameCmd.Flags().StringP("match", "m", "", "可选，仓库过滤条件，golang正则表达式")
	branchRenameCmd.Flags().StringP("invert-match", "i", "", "可选，仓库反向过滤条件，golang正则表达式")
}
