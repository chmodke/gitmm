// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "生成示例配置文件，校验配置文件",
	Long:  `生成示例配置文件，校验配置文件`,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
