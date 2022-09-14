// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/config"
	"gitmm/log"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "生成示例配置文件，校验配置文件",
	Long:  `生成示例配置文件，校验配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()
		log.Infof("upstream: %s", config.Upstream)
		log.Infof("origin: %s", config.Origin)
		log.Infof("repos: %s", config.Repos)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
