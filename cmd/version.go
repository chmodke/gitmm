/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show tool version",
	Long:    `Show tool version`,
	Example: "gitmm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitmm 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
