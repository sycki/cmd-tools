package portcheck

import (
	"os"
	"github.com/spf13/pflag"
	"fmt"
)

type Config struct {
	Host string
	Port string
	Protocol string
	Timeout int
}

func NewFromCmd() *Config {
	config := getDefaultConfig()
	cmd := pflag.CommandLine
	cmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <ip> <port>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		cmd.PrintDefaults()
	}

	cmd.IntVarP(&config.Timeout, "timeout", "w", config.Timeout, "Timeout for wait connection.")

	isTcp := cmd.BoolP("tcp", "t", true, "Use TCP protocol")
	isUdp := cmd.BoolP("udp", "u", false, "Use UDP protocol")

	pflag.Parse()

	if *isTcp { config.Protocol = "tcp" }
	if *isUdp { config.Protocol = "udp" }

	args := pflag.Args()
	if len(args) != 2 {
		cmd.Usage()
		os.Exit(35)
	}

	config.Host = args[0]
	config.Port = args[1]

	return config
}

func getDefaultConfig() *Config {
	return &Config{
		Protocol: "tcp",
		Timeout: 1,
	}
}
