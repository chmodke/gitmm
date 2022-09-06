package util

import (
	"fmt"
	"regexp"
	"strings"
)

func Title(text string, width int, padding string) string {
	length := len(text)
	paddingLen := (width - length) / 2
	return fmt.Sprintf("%s %s %s", strings.Repeat(padding, paddingLen), text, strings.Repeat(padding, width-length-paddingLen-2))
}

func LeftAlign(text string, width int, padding string) string {
	return fmt.Sprintf("%s %s", strings.Repeat(padding, width), text)
}

func Match(regex string, text string) bool {
	r, _ := regexp.Compile(regex)
	return r.MatchString(text)
}
