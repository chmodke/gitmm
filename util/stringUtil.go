package util

import (
	"fmt"
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

func Match(regex string, text string) bool {
	r, _ := regexp.Compile(regex)
	return r.MatchString(text)
}

func ExecStatistic(text string, result map[string]string) {
	fmt.Println(Title("", 80, "#"))
	fmt.Println(LeftAlign(fmt.Sprintf("%s statistics", text), 2, " "))
	fmt.Println(Title("", 80, "-"))
	maxLen := 0
	for k := range result {
		if len(k) > maxLen {
			maxLen = len(k)
		}
	}

	formatter := "  %-" + strconv.Itoa(maxLen) + "s \t%s\n"
	for k, v := range result {
		fmt.Printf(formatter, k, v)
	}
	fmt.Println(Title("", 80, "#"))
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
