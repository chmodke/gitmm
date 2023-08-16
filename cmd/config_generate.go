/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/chmodke/gitmm/config"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "生成示例配置文件",
	Long:    `生成示例配置文件`,
	Example: "gitmm config generate",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteCfg()
	},
}

func init() {
	configCmd.AddCommand(generateCmd)
}
