/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/chmodke/gitmm/cmd"
	"github.com/chmodke/gitmm/log"
)

func main() {
	cmd.Execute()
	log.FlushAndClose()
}
