package util

import (
	"fmt"
	"github.com/chmodke/gitmm/log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Title(text string, width int, padding string) string {
	length := len(text)
	paddingLen := (width - length) / 2
	return fmt.Sprintf("%s%s%s", strings.Repeat(padding, paddingLen), text, strings.Repeat(padding, width-length-paddingLen))
}

func LeftAlign(text string, width int, padding string) string {
	return fmt.Sprintf("%s %s", strings.Repeat(padding, width), text)
}

func Match(text string, regex string, invertRegex string) bool {
	if len(regex) == 0 && len(invertRegex) == 0 {
		return true
	}
	var result = true
	if len(regex) > 0 {
		r, _ := regexp.Compile(regex)
		result = result && r.MatchString(text)
	}
	if result && len(invertRegex) > 0 {
		ir, _ := regexp.Compile(invertRegex)
		result = result && !ir.MatchString(text)
	}
	return result
}

func ExecStatistic(text string, result map[string]string) {
	log.Consoleln(Title("", 80, "#"))
	log.Consoleln(LeftAlign(fmt.Sprintf("%s statistics", text), 2, " "))
	log.Consoleln(Title("", 80, "-"))
	maxLen := 0
	for k := range result {
		if len(k) > maxLen {
			maxLen = len(k)
		}
	}

	formatter := "  %-" + strconv.Itoa(maxLen) + "s \t%s\n"
	for k, v := range result {
		log.Consolef(formatter, k, v)
	}
	log.Consoleln(Title("", 80, "#"))
}

func RandCreator(l int) string {
	str := "abcdefghigklmnopqrstuvwxyz"
	strList := []byte(str)

	result := []byte{}
	i := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i = i + 1
	}
	return string(result)
}

func RightCut(str string, l int) string {
	if l <= 0 {
		return ""
	}
	if len(str) <= l {
		return str
	}
	return str[:l-2] + ".."
}

func LeftCut(str string, l int) string {
	if l <= 0 {
		return ""
	}
	if len(str) <= l {
		return str
	}
	start := len(str) - l + 2
	return ".." + str[start:]
}

func MaxLen(arr []string) int {
	max := 0
	for _, a := range arr {
		if len(a) > max {
			max = len(a)
		}
	}
	return max
}
