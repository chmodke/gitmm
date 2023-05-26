// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/git"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "批量拉取仓库",
	Long:    `执行命令会遍历work_dir目录下中的git仓库，并执行分支拉取操作。`,
	Example: "gitmm pull -w tmp",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Printf("work_dir: %s", workDir)
		force, _ := cmd.Flags().GetBool("force")
		log.Printf("force: %t", force)
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
			process.NewOption(util.RightCut(repo, 18), 0, 4)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := git.GitPull(filepath.Join(localDir, repo), force, &process)
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	pullCmd.Flags().BoolP("force", "f", false, "强制拉取")
	pullCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	pullCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
