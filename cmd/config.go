// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "生成示例配置文件",
	Long:  `生成示例配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteCfg()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
