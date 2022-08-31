/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitmm/config"
	"gitmm/log"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "校验配置文件",
	Long:  `校验配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadCfg()
		log.Infof("main_group: %s", config.MainGroup)
		log.Infof("origin_group: %s", config.OriginGroup)
		log.Infof("repos: %s", config.Repos)
	},
}

func init() {
	configCmd.AddCommand(verifyCmd)
}
