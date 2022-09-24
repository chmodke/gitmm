// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
		force, _ := cmd.Flags().GetBool("force")
		grep, _ := cmd.Flags().GetString("grep")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("force: %s", force)
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
		result := make(map[string]string)
		for _, repo := range repos {
			if len(grep) > 0 && !util.Match(grep, repo) {
				log.Info(util.LeftAlign(fmt.Sprintf("skip pull %s.\n", repo), 2, "-"))
				result[repo] = SKIP
				continue
			}
			log.Info(util.LeftAlign(fmt.Sprintf("start pull %s.", repo), 2, "-"))
			ok := util.GitPull(filepath.Join(localDir, repo), force)
			if ok {
				result[repo] = OK
				log.Info(util.LeftAlign(fmt.Sprintf("pull %s done.\n", repo), 2, "-"))
			} else {
				result[repo] = FAIL
				log.Error(util.LeftAlign(fmt.Sprintf("pull %s fail.\n", repo), 2, "-"))
			}
		}
		util.ExecStatistic("pull", result)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	pullCmd.Flags().BoolP("force", "f", false, "强制拉取")
	pullCmd.Flags().StringP("grep", "g", "", "仓库过滤条件，golang正则表达式")
}
