// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:     "switch",
	Short:   "批量切换分支",
	Long:    `执行命令会遍历work_dir中的git仓库，并执行分支切换操作。`,
	Example: "gitmm switch -w tmp -b develop",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Debugf("work_dir: %s", workDir)
		branch, _ := cmd.Flags().GetString("branch")
		log.Debugf("branch: %s", branch)
		force, _ := cmd.Flags().GetBool("force")
		log.Debugf("force: %s", force)
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
		for _, repo := range repos {
			var process util.Progress
			process.NewOption(util.RightCut(repo, 18), 0, 4)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := util.GitSwitchBranch(filepath.Join(localDir, repo), branch, force, &process)
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	switchCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	switchCmd.Flags().StringP("branch", "b", "master", "目标分支/tag/commit")
	switchCmd.MarkFlagRequired("branch")
	switchCmd.Flags().BoolP("force", "f", false, "强制切换")
	switchCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	switchCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
