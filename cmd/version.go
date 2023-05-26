// Package cmd /*
package cmd

import (
	"github.com/shirou/gopsutil/v3/host"
	"github.com/spf13/cobra"
	"gitmm/git"
	"gitmm/log"
)

var VERSION = "1.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show tool version",
	Long:    `Show tool version`,
	Example: "gitmm version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Consolef("gitmm version %s\n", VERSION)
		log.Consoleln(git.GetGitVersion())
		platform, _, version, _ := host.PlatformInformation()
		kernelArch, _ := host.KernelArch()
		log.Consolef("%s %s %s\n", platform, version, kernelArch)
		log.Consolef("Report Bug: <%s>\n", "https://gitee.com/chmodke/gitmm.git")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
