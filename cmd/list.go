// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitmm/log"
	"gitmm/util"
	"path/filepath"
	"strings"
)

// listCmd represents the batch command
// name
// status(A D M ?) git status --porcelain
// branch(git branch -l $(git branch --show-current) -v --format='%(upstream)') //%(upstream:remotename)
// lastcommit(git log -n1 --pretty="format:%ad %an" --date=iso)
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "展示工作路径下的Git仓库信息",
	Long:    `执行命令会遍历work_dir中的git仓库，并展示基础信息。`,
	Example: "gitmm list -w tmp",
	RunE: func(cmd *cobra.Command, args []string) error {
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Debugf("work_dir: %s", workDir)
		match, _ := cmd.Flags().GetString("match")
		log.Debugf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Debugf("invert: %s", invert)
		lineNumber, _ := cmd.Flags().GetInt("line-number")
		log.Debugf("line-number: %d", lineNumber)

		localDir, err := util.GetWorkDir(workDir)
		if err != nil {
			log.Error("获取工作路径失败")
			return nil
		}

		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}

		printHead()
		for _, repo := range repos {
			if !util.Match(repo, match, invert) {
				continue
			}

			repoPath := filepath.Join(localDir, repo)
			status := util.GitStatusStatistic(repoPath)
			branchName := util.GitCurrentBranch(repoPath)
			branchTrack := util.GitBranchTrack(repoPath, branchName)
			lastCommit := util.GitLastCommit(repoPath)
			printStatus(repo, branchName, branchTrack, lastCommit, status)
		}
		return nil
	},
}

func printHead() {
	fmt.Printf("%-15s  %-15s  %-12s  %-33s  %-34s\n", "Repo", "BranchName", "TrackTo", "LastCommit", "Status")
	fmt.Printf(strings.Repeat("-", 16) + "+" + strings.Repeat("-", 16) + "+" + strings.Repeat("-", 13) + "+" + strings.Repeat("-", 34) + "+" + strings.Repeat("-", 34) + "\n")
}

func printStatus(repo, branchName, branchTrack, lastCommit string, status map[string]int) {
	statusLine := ""
	if len(status) > 0 {
		for k, v := range status {
			statusLine += fmt.Sprintf("%s-%d ", k, v)
		}
	} else {
		statusLine = "clean"
	}
	fmt.Printf("%-15s  %-15s  %-12s  %-33s  %-34s\n", repo, branchName, branchTrack, lastCommit, statusLine)
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	listCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	listCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
	listCmd.Flags().IntP("line-number", "n", 1, "日志行数")
}
