// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "批量分支操作",
	Long:  `批量分支操作`,
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
