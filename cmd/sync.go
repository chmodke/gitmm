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
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "批量同步主从仓库",
	Long: `执行命令会读取当前目录下repo.yaml配置文件，遍历repos配置项，从upstream强制同步全部内容到origin中，需要用户对origin有强制写权限（取消分支保护）。
注意：会强制以upstream中的内容覆盖origin中的内容。`,
	Example: "gitmm sync",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()

		var (
			ok      bool
			from    string
			fromUrl string
			to      string
			toUrl   string
		)

		from, _ = cmd.Flags().GetString("from")
		to, _ = cmd.Flags().GetString("to")
		match, _ := cmd.Flags().GetString("match")
		invert, _ := cmd.Flags().GetString("invert-match")

		if fromUrl, ok = config.Remote[from]; !ok {
			log.Consolef("未配置%s远端地址\n", from)
			os.Exit(1)
		}

		if toUrl, ok = config.Remote[to]; !ok {
			log.Consolef("未配置%s远端地址\n", to)
			os.Exit(1)
		}

		log.Printf("from: %s", fromUrl)
		log.Printf("to: %s", toUrl)
		log.Printf("repos: %s", config.Repos)
		log.Consolef("sync repo from %s[%s] to %s[%s].", from, fromUrl, to, toUrl)
		sure := util.AreSure("Are you sure you want to continue?")
		if !sure {
			return
		}

		tmp := fmt.Sprintf("_%s_", util.RandCreator(8))
		for _, repo := range config.Repos {
			var process util.Progress
			process.NewOption(util.RightCut(repo, 18), 0, 9)
			if !util.Match(repo, match, invert) {
				process.Finish(SKIP)
				continue
			}
			ok := git.Sync(fromUrl, toUrl, repo, tmp, &process)
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
	syncCmd.Flags().StringP("from", "f", "upstream", "源端仓库地址")
	syncCmd.Flags().StringP("to", "t", "origin", "目标端仓库地址")
	syncCmd.Flags().StringP("match", "m", "", "仓库过滤条件，golang正则表达式")
	syncCmd.Flags().StringP("invert-match", "i", "", "仓库反向过滤条件，golang正则表达式")
}
