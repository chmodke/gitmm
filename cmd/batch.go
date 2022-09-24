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
	Long:    `执行命令会遍历work_dir中的git仓库，并执行提供的git命令。`,
	Example: "gitmm batch -w tmp 'log --oneline -n1'",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("请提供要执行的命令")
		}
		workDir, _ := cmd.Flags().GetString("work_dir")
		grep, _ := cmd.Flags().GetString("grep")
		log.Debugf("work_dir: %s", workDir)
		log.Debugf("grep: %s", grep)

		gitCommand := args[0]
		gitCommand = strings.TrimLeft(gitCommand, "git ")
		log.Debugf("git command: %s", gitCommand)

		localDir, err := util.GetWorkDir(workDir)
		if err != nil {
			log.Error("获取工作路径失败")
			return nil
		}

		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}

		result := make(map[string]string)
		for _, repo := range repos {
			if len(grep) > 0 && !util.Match(grep, repo) {
				log.Info(util.LeftAlign(fmt.Sprintf("skip execute command at %s.\n", repo), 2, "-"))
				result[repo] = SKIP
				continue
			}
			log.Info(util.LeftAlign(fmt.Sprintf("start execute command at %s.", repo), 2, "-"))
			ok := util.GitCommand(filepath.Join(localDir, repo), gitCommand)
			if ok {
				log.Info(util.LeftAlign(fmt.Sprintf("execute command at %s done.\n", repo), 2, "-"))
				result[repo] = OK
			} else {
				log.Error(util.LeftAlign(fmt.Sprintf("execute command at %s fail.\n", repo), 2, "-"))
				result[repo] = FAIL
			}
		}
		util.ExecStatistic("batch", result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	batchCmd.Flags().StringP("grep", "g", "", "仓库过滤条件，golang正则表达式")
}
