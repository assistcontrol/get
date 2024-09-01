package config

import (
	"errors"
	"fmt"
	"path/filepath"
)

type mode int

const (
	get mode = iota
	show
)

type Config struct {
	Mode mode
	Path string
	URL  string
}

func (c *Config) ShouldSave() bool {
	return c.Mode == get
}

func NewConfig(args []string) (*Config, error) {
	c := &Config{}

	if len(args) != 2 {
		return nil, errors.New("expected one argument")
	}

	c.URL = args[1]

	switch filepath.Base(args[0]) {
	case "get":
		c.Mode = get
	case "show":
		c.Mode = show
	default:
		return nil, fmt.Errorf("unknown command: %s", args[0])
	}

	return c, nil
}
