package main

import (
	"github.com/sycki/cmd-tools/pkg/portcheck"
	"github.com/sycki/cmd-tools/pkg/lookuprsv"
)

type Modules struct {
	childCmds map[string]string
}

func (m *Modules) PORTCHECK() {
	portcheck.Main()
}

func (m *Modules) LOOKUPSRV() {
	lookuprsv.Main()
}