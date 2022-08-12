// Package cmd /*
package cmd

import (
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:     "switch",
	Short:   "批量切换分支",
	Long:    `执行脚本会遍历work_dir中的git仓库，并执行分支切换操作。`,
	Example: "gitmm switch -w tmp -b develop",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		branch, _ := cmd.Flags().GetString("branch")
		force, _ := cmd.Flags().GetBool("force")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("branch: %s", branch)
		log.Debugf("force: %s", force)

		localDir := util.GetWorkDir(workDir)
		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}
		for _, repo := range repos {
			ok := util.GitSwitchBranch(filepath.Join(localDir, repo), branch, force)
			if ok {
				log.Infof("%s switch branch done.", repo)
			} else {
				log.Infof("%s switch branch fail.", repo)
			}
			log.Info(strings.Repeat("-", 80))
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	switchCmd.Flags().StringP("work_dir", "w", "master", "本地代码的存放路径")
	switchCmd.MarkFlagRequired("work_dir")
	switchCmd.Flags().StringP("branch", "b", "master", "目标分支/tag/commit")
	switchCmd.MarkFlagRequired("branch")
	switchCmd.Flags().BoolP("force", "f", false, "强制切换")
}
