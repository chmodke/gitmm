/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitmm/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitmm",
	Short: "git 多仓库管理工具",
	Long:  "一个git多仓库管理工具，通过简单的配置就可以批量管理多个仓库。",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.LoadCfg()
}
