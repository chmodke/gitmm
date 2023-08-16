// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/chmodke/gitmm/config"
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// remoteAddCmd represents the remote command
var remoteAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "批量添加仓库远程信息",
	Long:    `执行命令会遍历work_dir目录下中的git仓库，并添加仓库远程信息。`,
	Example: "gitmm remote add -u upstream",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()

		workDir, _ := cmd.Flags().GetString("work_dir")
		log.Printf("work_dir: %s", workDir)
		remote, _ := cmd.Flags().GetString("remote")
		log.Printf("remote: %s", remote)
		match, _ := cmd.Flags().GetString("match")
		log.Printf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Printf("invert: %s", invert)

		url, ok := config.Remote[remote]
		if !ok {
			log.Consolef("未配置%s远端地址\n", remote)
			os.Exit(1)
		}

		localDir, err := git.GetWorkDir(workDir)
		if err != nil {
			log.Consoleln("获取工作路径失败")
			return
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
			ok := git.RemoteAdd(filepath.Join(localDir, repo), remote, fmt.Sprintf("%s/%s.git", url, repo))
			log.Consoleln("")
			if ok {
				result[repo] = OK
			} else {
				result[repo] = FAIL
			}
		}
		util.ExecStatistic("remote", result)
	},
}

func init() {
	remoteCmd.AddCommand(remoteAddCmd)

	remoteAddCmd.Flags().StringP("work_dir", "w", ".", "本地代码的存放路径")
	remoteAddCmd.Flags().StringP("remote", "u", "origin", "远程名称")
	remoteAddCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	remoteAddCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
