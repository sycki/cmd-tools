package lookuprsv

import (
	"os"
	"github.com/spf13/pflag"
)

type Config struct {
	Host string
	Wrap bool
	Format string
	Delimiter string
	Timeout int
}

func NewFromCmd() *Config {
	config := getDefaultConfig()
	cmd := pflag.CommandLine

	cmd.BoolVar(&config.Wrap, "wrap", config.Wrap, "Automatic wrap for output for srv result")
	cmd.StringVar(&config.Format, "format", config.Format, "Format output for srv result")
	cmd.IntVarP(&config.Timeout, "timeout", "w", config.Timeout, "Timeout for wait connection.")

	pflag.Parse()

	args := pflag.Args()
	if len(args) < 1 {
		println("Must define host name!")
		os.Exit(25)
	}

	config.Host = args[0]

	if config.Wrap {
		config.Delimiter = "\n"
	}

	return config
}

func getDefaultConfig() *Config {
	return &Config{
		Timeout: 1,
		Wrap: false,
		Format: "%s",
		Delimiter: ",",
	}
}
