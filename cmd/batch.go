// Package cmd /*
package cmd

import (
	"errors"
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:     "batch [flags] -- command",
	Short:   "批量执行提供的git命令",
	Long:    `执行命令遍历work_dir中的git仓库，并执行提供的git命令。`,
	Example: "gitmm batch -- log --oneline -n1\n查看当前工作目录下所有仓库最新一条提交记录",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("请提供要执行的命令")
		}
		workDir, _ := cmd.Flags().GetString("work_dir")
		match, _ := cmd.Flags().GetString("match")
		invert, _ := cmd.Flags().GetString("invert-match")

		gitCommand := args
		if args[0] == "git" {
			gitCommand = args[1:]
		}
		log.Printf("git command: %v", strings.Join(gitCommand, " "))

		localDir, err := git.GetWorkDir(workDir)
		if err != nil {
			log.Consoleln("获取工作路径失败")
			return nil
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
			ok := git.Command(filepath.Join(localDir, repo), gitCommand)
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

	batchCmd.Flags().StringP("work_dir", "w", ".", "可选，本地代码的存放路径")
	batchCmd.Flags().StringP("match", "m", "", "可选，仓库过滤条件，golang正则表达式")
	batchCmd.Flags().StringP("invert-match", "i", "", "可选，仓库反向过滤条件，golang正则表达式")
}
