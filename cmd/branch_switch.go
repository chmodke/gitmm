// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

// branchSwitchCmd represents the switch command
var branchSwitchCmd = &cobra.Command{
	Use:     "switch",
	Short:   "批量切换分支",
	Long:    `执行命令会遍历work_dir中的git仓库，并执行分支切换操作。`,
	Example: "gitmm branch switch -w tmp develop",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		branch := args[0]
		force, _ := cmd.Flags().GetBool("force")
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
		for _, repo := range repos {
			var process util.Progress
			process.NewOption(util.RightCut(repo, 18), 0, 4)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := git.SwitchBranch(filepath.Join(localDir, repo), branch, force, &process)
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
	},
}

func init() {
	branchCmd.AddCommand(branchSwitchCmd)

	branchSwitchCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	branchSwitchCmd.Flags().BoolP("force", "f", false, "强制切换")
	branchSwitchCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	branchSwitchCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
