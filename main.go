/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"gitmm/cmd"
	"gitmm/log"
)

func main() {
	cmd.Execute()
	log.FlushAndClose()
}
