package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type mode int

const (
	GET mode = iota
	SHOW
)

type config struct {
	mode mode
	url  string
}

func main() {
	conf, err := newConfig(os.Args)
	if err != nil {
		log.Printf("Error: %v\n%s", err, help())
		os.Exit(1)
	}

	contents, err := fetch(conf.url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if conf.mode == GET {
		get(contents)
	} else {
		show(contents)
	}
}

func help() string {
	return "Usage: get <url>"
}

func newConfig(args []string) (*config, error) {
	c := &config{}

	if len(args) != 2 {
		return nil, errors.New("expected one argument")
	}

	c.url = args[1]

	switch filepath.Base(args[0]) {
	case "get":
		c.mode = GET
	case "show":
		c.mode = SHOW
	default:
		return nil, fmt.Errorf("unknown command: %s", args[0])
	}

	return c, nil
}
