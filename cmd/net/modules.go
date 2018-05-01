package main

import (
	"fmt"
	"os"
)

type Commands struct {
	childCmds map[string]string
	short     map[string]string
	funcs     map[string]func()
}

func (c *Commands) Add(cmd, describe string, f func()) {
	c.childCmds[cmd] = describe
	c.funcs[cmd] = f
}

func (c *Commands) AddP(cmd, short, describe string, f func()) {
	c.childCmds[cmd] = fmt.Sprintf("%s (aka '%s')", describe, short)
	c.short[short] = cmd
	c.funcs[cmd] = f
}

func (c *Commands) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <child command> [options] <args>\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "All child command:\n")
	for k, v := range c.childCmds {
		fmt.Fprintf(os.Stderr, "  %s\t\t%s\n", k, v)
	}
}

func NewCommands() *Commands {
	initNum := 3
	return &Commands{
		childCmds: make(map[string]string, initNum),
		short:     make(map[string]string, initNum),
		funcs:     make(map[string]func(), initNum),
	}
}
