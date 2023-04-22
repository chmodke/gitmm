// Package cmd /*
package cmd

import (
	"errors"
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
		log.Printf("work_dir: %s", workDir)
		match, _ := cmd.Flags().GetString("match")
		log.Printf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Printf("invert: %s", invert)

		gitCommand := args[0]
		gitCommand = strings.TrimLeft(gitCommand, "git ")
		log.Printf("git command: %s", gitCommand)

		localDir, err := util.GetWorkDir(workDir)
		if err != nil {
			log.Consoleln("获取工作路径失败")
			return nil
		}

		repos, err := util.FindGit(localDir)
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
			ok := util.GitCommand(filepath.Join(localDir, repo), gitCommand)
			log.Consoleln("")
			if ok {
				result[repo] = OK
			} else {
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
	batchCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	batchCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
