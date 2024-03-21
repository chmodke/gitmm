/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/chmodke/gitmm/config"
	"github.com/chmodke/gitmm/log"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:     "verify",
	Short:   "校验配置文件",
	Long:    `校验配置文件`,
	Example: "gitmm config verify\n校验在当前目录下repo.yaml文件格式是否正确",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()
		log.Consolef("upstream: %s", config.Upstream)
		log.Consolef("origin: %s", config.Origin)
		log.Consolef("repos: %s", config.Repos)
	},
}

func init() {
	configCmd.AddCommand(verifyCmd)
}
