// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitmm/util"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var VERSION = "1.0.3"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show tool version",
	Long:    `Show tool version`,
	Example: "gitmm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gitmm version %s\n", VERSION)
		fmt.Println(GetGitVersion())
		fmt.Printf("%s %s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Report Bug: <%s>\n", "https://gitee.com/chmodke/gitmm.git")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func GetGitVersion() string {
	command := "git --version"
	out, ok := util.GetOut(util.Execute(command))
	if !ok {
		fmt.Println("执行git失败，请检查是否安装git，或者环境变量配置错误。")
		fmt.Println("下载地址: <https://repo.huaweicloud.com/git-for-windows/>")
		os.Exit(1)
		return ""
	} else {
		return out
	}
}

func CheckGitVersion(version string) {
	r, _ := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+")
	ver := r.FindString(version)

	if !newVersion(ver, "2.28.0") {
		fmt.Println("git版本低于2.28.0，部分功能不可用。")
		fmt.Println("下载地址: <https://repo.huaweicloud.com/git-for-windows/>")
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
