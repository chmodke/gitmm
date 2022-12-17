package util

import (
	"fmt"
	"strings"
)

type CmdBuilder struct {
	commands []string
}

func (b *CmdBuilder) Add(command string) *CmdBuilder {
	b.AddWithCond(command, true)
	return b
}

func (b *CmdBuilder) AddWithArg(command string, args ...interface{}) *CmdBuilder {
	b.Add(fmt.Sprintf(command, args...))
	return b
}

func (b *CmdBuilder) AddWithCond(command string, cond bool) *CmdBuilder {
	if cond {
		b.commands = append(b.commands, command)
	}
	return b
}

func (b *CmdBuilder) AddWithArgAndCond(command string, cond bool, args ...interface{}) *CmdBuilder {
	if cond {
		b.Add(fmt.Sprintf(command, args...))
	}
	return b
}

func (b *CmdBuilder) Build() string {
	if b.commands == nil {
		return ""
	}
	return strings.Join(b.commands, " ")
}

func (b *CmdBuilder) Reset() *CmdBuilder {
	if b.commands == nil {
		return b
	}
	b.commands = b.commands[:0]
	return b
}
