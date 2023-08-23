package util

import (
	"fmt"
	"strings"
)

func AreSure(prompt string) bool {
	var (
		input string
	)
	for {
		fmt.Printf("%s (yes/no): ", prompt)
		fmt.Scanln(&input)
		switch strings.ToLower(input) {
		case "yes", "y":
			return true
		case "no", "n":
			return false
		}
	}
}
