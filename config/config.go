package config

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Filename string
	Force    bool
	Saving   bool
	URL      string
}

var opts struct {
	Force  bool `short:"f" description:"overwrite existing files"`
	Output bool `short:"o" description:"save output to 'filename', or leave empty to use a best guess"`
}

func NewConfig() (*Config, error) {
	c := &Config{}

	p := flags.NewParser(&opts, flags.Default)
	p.Usage = `[-f] [-o [filename]] <URL or path>`

	remaining, err := p.Parse()
	if err != nil {
		return nil, err
	}

	switch {
	case len(remaining) == 1:
		c.URL = remaining[0]
		if opts.Output {
			c.Saving = true
		}
	case len(remaining) == 2 && opts.Output:
		c.Filename = remaining[0]
		c.URL = remaining[1]
		c.Saving = true
	case len(remaining) == 0:
		fmt.Fprintf(os.Stderr, "ERROR: path or URL required\n\n")
		fallthrough
	default:
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	return c, nil
}
