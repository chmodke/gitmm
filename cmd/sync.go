/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitmm/config"
	"gitmm/util"
	"os"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "批量同步主从仓库",
	Long: `执行脚本会读取当前目录下repo.yaml配置文件，遍历repos配置项，从main_group强制同步全部内容到origin_group中，需要用户对origin_group有强制写权限（取消分支保护）。
注意：会强制以main_group中的内容覆盖origin_group中的内容。`,
	Example: "gitmm sync",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("main_group: %s\n", config.MainGroup)
		fmt.Printf("origin_group: %s\n", config.OriginGroup)
		fmt.Printf("repos: %s\n", config.Repos)
		for _, repo := range config.Repos {
			util.GitSync(config.MainGroup, config.OriginGroup, repo, "tmp")
		}
		os.RemoveAll("tmp")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
