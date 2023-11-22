// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

// branchListCmd represents the delete command
var branchListCmd = &cobra.Command{
	Use:     "list",
	Short:   "批量产看分支",
	Long:    `执行命令会遍历work_dir中的git仓库，并执行分支产看操作。`,
	Example: "gitmm branch list -w tmp",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
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
		result := make(map[string]string)
		for _, repo := range repos {
			if !util.Match(repo, match, invert) {
				result[repo] = SKIP
				continue
			}
			log.Consoleln(repo)
			ok := git.ListBranch(filepath.Join(localDir, repo))
			log.Consoleln("")
			if ok {
				result[repo] = OK
			} else {
				result[repo] = FAIL
			}
		}
		util.ExecStatistic("remote", result)
	},
}

func init() {
	branchCmd.AddCommand(branchListCmd)

	branchListCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	branchListCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	branchListCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
