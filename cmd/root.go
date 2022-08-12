// Package cmd /*
package cmd

import (
	"fmt"
	"gitmm/log"
	"gitmm/util"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitmm",
	Short: "git多仓库管理工具",
	Long:  "git多仓库管理工具，通过简单的配置对仓库进行批量管理。",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&log.DEBUG, "debug", "x", false, "debug")
	command := "git --version"
	out, ok := util.GetOut(util.Execute(command))
	if !ok {
		fmt.Println("执行git失败，请检查是否安装git，或者环境变量配置错误。")
		os.Exit(1)
	}

	r, _ := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+")
	ver := r.FindString(out)

	if !newVersion(ver, "2.28.0") {
		fmt.Println("git版本低于2.28.0，部分功能不可用。")
	}

}

func newVersion(ver1 string, ver2 string) bool {
	vers1 := strings.Split(ver1, ".")
	vers2 := strings.Split(ver2, ".")

	length := len(vers1)
	if len(vers2) < length {
		length = len(vers2)
	}

	for i := 0; i < length; i++ {
		v1, _ := strconv.Atoi(vers1[i])
		v2, _ := strconv.Atoi(vers2[i])
		if v1 > v2 {
			return true
		} else if v1 < v2 {
			return false
		}
	}

	return false
}
