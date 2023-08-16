// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
)

// remoteCmd represents the branch command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "远程地址管理",
	Long:  `远程地址管理`,
}

func init() {
	rootCmd.AddCommand(remoteCmd)
}
