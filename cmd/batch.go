// Package cmd /*
package cmd

import (
	"errors"
	"fmt"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:     "batch",
	Short:   "批量执行提供的git命令",
	Long:    `执行脚本会遍历work_dir中的git仓库，并执行提供的git命令。`,
	Example: "gitmm batch -w tmp 'log --oneline -n1'",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("请提供要执行的命令")
		}
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Debugf("work_dir: %s", workDir)

		gitCommand := args[0]
		gitCommand = strings.TrimLeft(gitCommand, "git ")
		log.Debugf("git command: %s", gitCommand)

		localDir := util.GetWorkDir(workDir)
		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}

		for _, repo := range repos {
			log.Info(util.Title(fmt.Sprintf("start execute command at %s.", repo), 80, "-"))
			ok := util.GitCommand(filepath.Join(localDir, repo), gitCommand)
			if ok {
				log.Info(util.Title(fmt.Sprintf("execute command %s done.", repo), 80, "-"))
			} else {
				log.Error(util.Title(fmt.Sprintf("execute command %s fail.", repo), 80, "-"))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
}
