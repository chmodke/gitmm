// Package cmd /*
package cmd

import (
	"github.com/chmodke/gitmm/git"
	"github.com/chmodke/gitmm/log"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	OK   = "ok"
	FAIL = "fail"
	SKIP = "skip"
)

var rootCmd = &cobra.Command{
	Use:   "gitmm",
	Short: "git多仓库管理工具",
	Long:  "git多仓库管理工具，通过简单的配置对仓库进行批量管理。",
}

func Execute() {
	log.Printf("main command: [gitmm %v]", strings.Join(os.Args[1:], " "))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = VERSION
	rootCmd.Flags().BoolP("version", "v", false, "show tool version.")
	version := git.GetGitVersion()
	checkGitVersion(version)
}

func checkGitVersion(version string) {
	r, _ := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+")
	ver := r.FindString(version)

	if !newVersion(ver, "2.28.0") {
		log.Consoleln("git版本低于2.28.0，部分功能不可用。")
		log.Consoleln("下载地址: <https://repo.huaweicloud.com/git-for-windows/>")
	}
}

func newVersion(ver1 string, ver2 string) bool {
	part1 := strings.Split(ver1, ".")
	part2 := strings.Split(ver2, ".")

	length := len(part1)
	if len(part2) < length {
		length = len(part2)
	}

	for i := 0; i < length; i++ {
		v1, _ := strconv.Atoi(part1[i])
		v2, _ := strconv.Atoi(part2[i])
		if v1 > v2 {
			return true
		} else if v1 < v2 {
			return false
		}
	}

	return false
}
