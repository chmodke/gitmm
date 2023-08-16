// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "分支操作",
	Long:  `分支操作`,
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
