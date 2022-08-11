/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "批量拉取仓库",
	Long:  `执行脚本会遍历work_dir目录下中的git仓库，并执行分支拉取操作。`,
	Run: func(cmd *cobra.Command, args []string) {
		workDir, _ := cmd.Flags().GetString("work_dir")
		force, _ := cmd.Flags().GetBool("force")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("force: %s", force)

		localDir := util.GetWorkDir(workDir)
		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}
		for _, repo := range repos {
			ok := util.GitPull(filepath.Join(localDir, repo), force)
			if ok {
				log.Infof("pull %s done.", repo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().BoolVarP(&log.DEBUG, "debug", "x", false, "debug")
	pullCmd.Flags().StringP("work_dir", "w", "master", "克隆代码的存放路径")
	pullCmd.MarkFlagRequired("work_dir")
	pullCmd.Flags().BoolP("force", "f", false, "强制拉取")
}
