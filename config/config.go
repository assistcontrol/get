package config

import (
	"errors"
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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: get [-f] [-o [filename]] <URL or path>\n")
		flag.PrintDefaults()
	}

	c := &Config{}
	flag.BoolVar(&c.Force, "f", false, "overwrite existing files")
	flag.BoolVar(&c.Saving, "o", false, "save output to `[filename]`, or leave empty to use a best guess")
	flag.Parse()

	switch {
	case flag.NArg() == 1:
		// If there's only one argument, it's a URL
		c.URL = flag.Arg(0)
	case flag.NArg() == 2 && c.Saving:
		// Two args is a filename and a URL, but only if -o is set
		c.Filename = flag.Arg(0)
		c.URL = flag.Arg(1)
	default:
		return nil, errors.New("wrong number of arguments")
	}

	return c, nil
}
