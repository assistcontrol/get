package context

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// Ctx combines configuration and runtime context.
type Ctx struct {
	// Configuration
	Filename string // User-specified filename
	Force    bool   // Overwrite existing files
	Path     string // User-provided URL or local path
	Save     bool   // Whether to save output to a file

	// Runtime context
	Body        []byte         // Contents of the requested resource
	Destination string         // Local filename to save to
	Response    *http.Response // HTTP Response object for later introspection
	URL         string         // Crafted URL to try fetching
}

var force, save bool

func init() {
	flag.BoolVar(&force, "f", false, "overwrite existing files")
	flag.BoolVar(&save, "o", false, "save output to `[filename]`, or leave empty to use a best guess")
	flag.Usage = usage
}

// New creates a new context from the command line arguments.
// It returns an error if the arguments are invalid.
func New(args []string) (*Ctx, error) {
	if err := flag.CommandLine.Parse(args); err != nil {
		return nil, err
	}

	c := &Ctx{
		Force: force,
		Save:  save,
	}

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
