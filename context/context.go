package context

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Ctx struct {
	Body        []byte
	Destination string
	Filename    string
	Force       bool
	Response    *http.Response
	Path        string
	Save        bool
	URL         string
}

func New() (*Ctx, error) {
	c := &Ctx{}

	flag.BoolVar(&c.Force, "f", false, "overwrite existing files")
	flag.BoolVar(&c.Save, "o", false, "save output to `[filename]`, or leave empty to use a best guess")
	flag.Usage = usage
	flag.Parse()

	switch {
	case flag.NArg() == 1:
		// If there's only one argument, it's a URL
		c.Path = flag.Arg(0)
	case flag.NArg() == 2 && c.Save:
		// Two args is a filename and a URL, but only if -o is set
		c.Filename = flag.Arg(0)
		c.Path = flag.Arg(1)
	default:
		return nil, errors.New("wrong number of arguments")
	}

	return c, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: get [-f] [-o [filename]] <URL or path>\n")
	flag.PrintDefaults()
}