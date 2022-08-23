// Package cmd /*
package cmd

import (
	"fmt"
	"gitmm/util"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show tool version",
	Long:    `Show tool version`,
	Example: "gitmm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gitmm %s\n", VERSION)
		out, _ := util.GetOut(util.Execute("git --version"))
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
