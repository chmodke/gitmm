// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

// listCmd represents the batch command
// name
// status(A D M ?) git status --porcelain
// branch(git branch -l $(git branch --show-current) -v --format='%(upstream:lstrip=2)') //%(upstream:remotename)
// lastcommit(git log -n1 --pretty="format:%ad %an" --date=iso)
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "展示工作路径下的Git仓库信息",
	Long:    `执行命令会遍历work_dir中的git仓库，并展示基础信息。`,
	Example: "gitmm list -w tmp",
	RunE: func(cmd *cobra.Command, args []string) error {
		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Printf("work_dir: %s", workDir)
		match, _ := cmd.Flags().GetString("match")
		log.Printf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Printf("invert: %s", invert)
		lineNumber, _ := cmd.Flags().GetInt("line-number")
		log.Printf("line-number: %d", lineNumber)

		localDir, err := git.GetWorkDir(workDir)
		if err != nil {
			log.Consoleln("获取工作路径失败")
			return nil
		}

		repos, err := git.FindGit(localDir)
		if err != nil {
			log.Consoleln("获取本地仓库失败")
		}

		printHead()
		for _, repo := range repos {
			if !util.Match(repo, match, invert) {
				continue
			}

			repoPath := filepath.Join(localDir, repo)
			status := git.StatusStatistic(repoPath)
			branchName := git.CurrentBranch(repoPath)
			branchTrack := git.BranchTrack(repoPath, branchName)
			lastCommit := git.LastCommit(repoPath)
			printStatus(repo, branchName, branchTrack, lastCommit, status)
		}
		printSplit(17, 17, 23, 35, 34)
		return nil
	},
}

func printLine(one, two, three, four, five string) {
	log.Consolef("%-18s%-18s%-24s%-36s%-34s\n", one, two, three, four, five)
}

func printSplit(one, two, three, four, five int) {
	log.Consolef(strings.Repeat("-", one) + "+" + strings.Repeat("-", two) + "+" + strings.Repeat("-", three) + "+" + strings.Repeat("-", four) + "+" + strings.Repeat("-", five) + "\n")
}

func printHead() {
	printLine("Repo", "BranchName", "TrackTo", "LastCommit", "Status")
	printSplit(17, 17, 23, 35, 34)
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
	printLine(util.RightCut(repo, 16), util.LeftCut(branchName, 16), util.LeftCut(branchTrack, 23), util.LeftCut(lastCommit, 34), statusLine)
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	listCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	listCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
	listCmd.Flags().IntP("line-number", "n", 1, "日志行数")
}
