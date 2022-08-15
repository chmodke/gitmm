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
	Long:    `执行脚本会遍历work_dir目录下中的git仓库，并执行分支拉取操作。`,
	Example: "gitmm pull -w tmp",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		force, _ := cmd.Flags().GetBool("force")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("force: %s", force)

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
			log.Info(util.Title(fmt.Sprintf("start pull %s.", repo), 80, "-"))
			ok := util.GitPull(filepath.Join(localDir, repo), force)
			if ok {
				log.Info(util.Title(fmt.Sprintf("pull %s done.", repo), 80, "-"))
			} else {
				log.Error(util.Title(fmt.Sprintf("pull %s fail.", repo), 80, "-"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	pullCmd.Flags().BoolP("force", "f", false, "强制拉取")
}
