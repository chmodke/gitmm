/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gitmm/config"
	"gitmm/util"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitmm",
	Short: "git 多仓库管理工具",
	Long:  "一个git多仓库管理工具，通过简单的配置就可以批量管理多个仓库。",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	command := "git --version"
	out, ok := util.GetOut(util.Execute(command))
	if !ok {
		fmt.Println("执行git失败，请检查是否安装git，或者环境变量配置错误。")
		os.Exit(1)
	}

	r, _ := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+")
	ver := r.FindString(out)

	if !compareVersion(ver, "2.28.0") {
		fmt.Println("git版本低于2.28.0，部分功能不可用。")
	}

	config.LoadCfg()
}

func compareVersion(ver1 string, ver2 string) bool {
	return false
}
