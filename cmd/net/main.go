package main

import (
	"fmt"
	"github.com/sycki/cmd-tools/pkg/lookupsrv"
	"github.com/sycki/cmd-tools/pkg/portcheck"
	"os"
)

func main() {
	cmd := NewCommands()

	cmd.AddP("portcheck", "p", "Check the port is open", portcheck.Main)
	cmd.AddP("lookupsrv", "l", "Find peers by host name", lookupsrv.Main)

	if len(os.Args) < 2 {
		cmd.Usage()
		os.Exit(11)
	}

	childCmd := os.Args[1]
	if len(childCmd) == 1 {
		value, ok := cmd.short[childCmd]
		if ok {
			childCmd = value
		} else {
			fmt.Fprintf(os.Stderr, "Unknown child command: %s\n\n", childCmd)
			cmd.Usage()
			os.Exit(13)
		}
	}
	f, ok := cmd.funcs[childCmd]

	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown found command: %s\n\n", childCmd)
		cmd.Usage()
		os.Exit(15)
	}

	os.Args[1] = fmt.Sprintf("%s %s", os.Args[0], childCmd)
	os.Args = os.Args[1:]

	f()

}
