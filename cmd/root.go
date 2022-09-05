// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"gitmm/log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gitmm",
	Short: "git多仓库管理工具",
	Long:  "git多仓库管理工具，通过简单的配置对仓库进行批量管理。",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level, _ := cmd.Flags().GetString("debug")
		log.SetLevel(level)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = VERSION
	rootCmd.Flags().BoolP("version", "v", false, "show tool version.")
	rootCmd.PersistentFlags().StringP("debug", "x", "info", "show more detail.")
	version := GetGitVersion()
	CheckGitVersion(version)
}
