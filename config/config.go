package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Filename string
	Force    bool
	Saving   bool
	URL      string
}

func NewConfig() (*Config, error) {
	c := &Config{}

	flag.BoolVar(&c.Force, "f", false, "overwrite existing files")
	flag.BoolVar(&c.Saving, "o", false, "save output to 'filename', or leave empty to use a best guess")
	flag.Parse()

	switch {
	case flag.NArg() == 1:
		c.URL = flag.Arg(0)
	case flag.NArg() == 2 && c.Saving:
		c.Filename = flag.Arg(0)
		c.URL = flag.Arg(1)
	case flag.NArg() == 0:
		fmt.Fprintf(os.Stderr, "ERROR: path or URL required\n\n")
		fallthrough
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	return c, nil
}
