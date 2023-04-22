// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

const (
	OK   = "ok"
	FAIL = "fail"
	SKIP = "skip"
)

var rootCmd = &cobra.Command{
	Use:   "gitmm",
	Short: "git多仓库管理工具",
	Long:  "git多仓库管理工具，通过简单的配置对仓库进行批量管理。",
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
	version := GetGitVersion()
	CheckGitVersion(version)
}
