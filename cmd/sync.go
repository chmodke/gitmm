// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitmm/config"
	"gitmm/git"
	"gitmm/log"
	"gitmm/util"
	"os"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "批量同步主从仓库",
	Long: `执行命令会读取当前目录下repo.yaml配置文件，遍历repos配置项，从upstream强制同步全部内容到origin中，需要用户对origin有强制写权限（取消分支保护）。
注意：会强制以upstream中的内容覆盖origin中的内容。`,
	Example: "gitmm sync",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()
		log.Printf("upstream: %s", config.Upstream)
		log.Printf("origin: %s", config.Origin)
		log.Printf("repos: %s", config.Repos)

		match, _ := cmd.Flags().GetString("match")
		log.Printf("match: %s", match)
		invert, _ := cmd.Flags().GetString("invert-match")
		log.Printf("invert: %s", invert)

		log.Consolef("sync repo from %s to %s.", config.Upstream, config.Origin)

		tmp := fmt.Sprintf("_%s_", util.RandCreator(8))
		for _, repo := range config.Repos {
			var process util.Progress
			process.NewOption(util.RightCut(repo, 18), 0, 9)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := git.GitSync(config.Upstream, config.Origin, repo, tmp, &process)
			if ok {
				process.Finish(OK)
			} else {
				process.Finish(FAIL)
			}
		}
		os.RemoveAll(tmp)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	syncCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
