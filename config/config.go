package config

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
)

type mode int

const (
	get mode = iota
	show
)

type Config struct {
	Mode     mode
	Filename string
	Force    bool
	URL      string
}

func (c *Config) ShouldSave() bool {
	return c.Mode == get
}

func NewConfig(args []string) (*Config, error) {
	arg0 := filepath.Base(args[0])

	c := &Config{}

	showFlag := flag.Bool("show", false, "output to stdout")
	flag.BoolVar(&c.Force, "f", false, "overwrite existing files")
	flag.StringVar(&c.Filename, "o", "", "save output to `filename`")
	flag.Parse()

	if flag.NArg() != 1 {
		return nil, errors.New("path or URL required")
	}

	c.URL = flag.Arg(0)

	switch {
	case arg0 == "get" && *showFlag:
		c.Mode = show
	case arg0 == "get":
		c.Mode = get
	case arg0 == "show" && c.Filename != "":
		c.Mode = get
	case arg0 == "show":
		c.Mode = show
	default:
		return nil, fmt.Errorf("unknown command: %s", arg0)
	}

	return c, nil
}
