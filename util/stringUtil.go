package util

import (
	"fmt"
	"strings"
)

func Title(text string, width int, padding string) string {
	length := len(text)
	paddingLen := (width - length) / 2
	return fmt.Sprintf("%s%s%s", strings.Repeat(padding, paddingLen), text, strings.Repeat(padding, width-length-paddingLen))
}
