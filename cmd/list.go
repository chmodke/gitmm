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

		localDir, err := util.GetWorkDir(workDir)
		if err != nil {
			log.Error("获取工作路径失败")
			return nil
		}

		repos, err := util.FindGit(localDir)
		if err != nil {
			log.Error("获取本地仓库失败")
		}

		var commands []string
		commands = append(commands, "git")
		commands = append(commands, "-C %s")
		commands = append(commands, "log -n1")
		commands = append(commands, "--pretty=\"format:%%ad %%h %%d %%n%%s%%n\"")
		commands = append(commands, "--date=iso")
		preCmd := strings.Join(commands, " ")

		result := make(map[string]string)
		for _, repo := range repos {
			if !util.Match(repo, match, invert) {
				result[repo] = SKIP
				continue
			}

			command := fmt.Sprintf(preCmd, filepath.Join(localDir, repo))
			info, _ := util.GetOut(util.Execute(command))
			log.InfoOf("%s\n%s\n", repo, info)
			result[repo] = OK
		}
		util.ExecStatistic("list", result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	listCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	listCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
